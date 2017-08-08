// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package crossmodelrelations

import (
	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/utils/clock"
	"gopkg.in/macaroon.v1"

	"github.com/juju/juju/api/base"
	apiwatcher "github.com/juju/juju/api/watcher"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/watcher"
)

var logger = loggo.GetLogger("juju.api.crossmodelrelations")

// Client provides access to the crossmodelrelations api facade.
type Client struct {
	base.ClientFacade
	facade base.FacadeCaller

	cache *MacaroonCache
}

// NewClient creates a new client-side CrossModelRelations facade.
func NewClient(caller base.APICallCloser) *Client {
	return NewClientWithCache(caller, NewMacaroonCache(clock.WallClock))
}

// NewClientWithCache creates a new client-side CrossModelRelations facade
// with the specified cache.
func NewClientWithCache(caller base.APICallCloser, cache *MacaroonCache) *Client {
	frontend, backend := base.NewClientFacade(caller, "CrossModelRelations")
	return &Client{
		ClientFacade: frontend,
		facade:       backend,
		cache:        cache,
	}
}

func (c *Client) Close() error {
	return c.ClientFacade.Close()
}

// handleError is used to process an error obtained when making a facade call.
// If the error indicates that a macaroon discharge is required, this is done
// and the resulting discharge macaroons passed back so the api call can be retried.
func (c *Client) handleError(err error) (macaroon.Slice, error) {
	if params.ErrCode(err) != params.CodeDischargeRequired {
		return nil, err
	}
	errResp := errors.Cause(err).(*params.Error)
	if errResp.Info == nil {
		return nil, errors.Annotatef(err, "no error info found in discharge-required response error")
	}
	logger.Debugf("attempting to discharge macaroon due to error: %v", err)
	ms, err := c.facade.RawAPICaller().BakeryClient().DischargeAll(errResp.Info.Macaroon)
	if err == nil && logger.IsTraceEnabled() {
		logger.Tracef("discharge macaroon ids:")
		for _, m := range ms {
			logger.Tracef("  - %v", m.Id())
		}
	}
	return ms, err
}

func (c *Client) getCachedMacaroon(opName, token string) (macaroon.Slice, bool) {
	ms, ok := c.cache.Get(token)
	if ok {
		logger.Debugf("%s using cached macaroons for %s", opName, token)
		if logger.IsTraceEnabled() {
			for _, m := range ms {
				logger.Tracef("  - %v", m.Id())
			}
		}
	}
	return ms, ok
}

// PublishRelationChange publishes relation changes to the
// model hosting the remote application involved in the relation.
func (c *Client) PublishRelationChange(change params.RemoteRelationChangeEvent) (err error) {
	args := params.RemoteRelationsChanges{
		Changes: []params.RemoteRelationChangeEvent{change},
	}
	// Use any previously cached discharge macaroons.
	if ms, ok := c.getCachedMacaroon("publish relation changed", change.RelationToken); ok {
		args.Changes[0].Macaroons = ms
	}

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.ErrorResults
		if err := c.facade.FacadeCall("PublishRelationChanges", args, &results); err != nil {
			return errors.Trace(err)
		}
		err = results.OneError()
		if params.IsCodeNotFound(err) {
			return errors.NotFoundf("relation for event %v", change)
		}
		if err != nil && i == 0 {
			mac, err2 := c.handleError(err)
			if err2 != nil {
				err = errors.Wrap(err, err2)
				return err
			}
			args.Changes[0].Macaroons = mac
			c.cache.Upsert(args.Changes[0].RelationToken, mac)
		}
	}
	return err
}

func (c *Client) PublishIngressNetworkChange(change params.IngressNetworksChangeEvent) (err error) {
	args := params.IngressNetworksChanges{
		Changes: []params.IngressNetworksChangeEvent{change},
	}
	// Use any previously cached discharge macaroons.
	if ms, ok := c.getCachedMacaroon("publish ingress network change", change.RelationToken); ok {
		args.Changes[0].Macaroons = ms
	}

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.ErrorResults
		if err := c.facade.FacadeCall("PublishIngressNetworkChanges", args, &results); err != nil {
			return errors.Trace(err)
		}
		err = results.OneError()
		if err != nil && i == 0 {
			mac, err2 := c.handleError(err)
			if err2 != nil {
				err = errors.Wrap(err, err2)
				return err
			}
			args.Changes[0].Macaroons = mac
			c.cache.Upsert(args.Changes[0].RelationToken, mac)
		}
	}
	return err
}

// RegisterRemoteRelations sets up the remote model to participate
// in the specified relations.
func (c *Client) RegisterRemoteRelations(relations ...params.RegisterRemoteRelationArg) ([]params.RegisterRemoteRelationResult, error) {
	var (
		args         params.RegisterRemoteRelationArgs
		retryIndices []int
	)

	result := make([]params.RegisterRemoteRelationResult, len(relations))
	args = params.RegisterRemoteRelationArgs{Relations: relations}

	// Use any previously cached discharge macaroons.
	for i, arg := range relations {
		if ms, ok := c.getCachedMacaroon("register remote relation", arg.RelationToken); ok {
			newArg := arg
			newArg.Macaroons = ms
			args.Relations[i] = newArg
		}
	}

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.RegisterRemoteRelationResults
		err := c.facade.FacadeCall("RegisterRemoteRelations", args, &results)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if len(results.Results) != len(args.Relations) {
			return nil, errors.Errorf("expected %d result(s), got %d", len(args.Relations), len(results.Results))
		}

		if i == 0 {
			args = params.RegisterRemoteRelationArgs{}
			// Separate the successful calls from those needing a retry.
			for j, res := range results.Results {
				if res.Error == nil {
					result[j] = res
					continue
				}
				mac, err := c.handleError(res.Error)
				if err != nil {
					res.Error.Message = err.Error()
					result[j] = res
					continue
				}
				retryArg := relations[j]
				retryArg.Macaroons = mac
				args.Relations = append(args.Relations, retryArg)
				retryIndices = append(retryIndices, i)
				c.cache.Upsert(retryArg.RelationToken, mac)
			}
			if len(args.Relations) == 0 {
				break
			}
		}
		// After a retry, insert the results into the original result slice.
		if i == 1 {
			for j, res := range results.Results {
				result[retryIndices[j]] = res
			}
		}
	}
	return result, nil
}

// WatchRelationUnits returns a watcher that notifies of changes to the
// units in the remote model for the relation with the given remote token.
func (c *Client) WatchRelationUnits(remoteRelationArg params.RemoteEntityArg) (watcher.RelationUnitsWatcher, error) {
	args := params.RemoteEntityArgs{Args: []params.RemoteEntityArg{remoteRelationArg}}
	// Use any previously cached discharge macaroons.
	if ms, ok := c.getCachedMacaroon("watch relation units", remoteRelationArg.Token); ok {
		args.Args[0].Macaroons = ms
	}

	var result params.RelationUnitsWatchResult

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.RelationUnitsWatchResults
		err := c.facade.FacadeCall("WatchRelationUnits", args, &results)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if len(results.Results) != 1 {
			return nil, errors.Errorf("expected 1 result, got %d", len(results.Results))
		}
		result = results.Results[0]
		if result.Error != nil && i == 0 {
			mac, err := c.handleError(result.Error)
			if err != nil {
				err = errors.Wrap(result.Error, err)
				return nil, err
			}
			args.Args[0].Macaroons = mac
			c.cache.Upsert(args.Args[0].Token, mac)
		}
	}
	w := apiwatcher.NewRelationUnitsWatcher(c.facade.RawAPICaller(), result)
	return w, nil
}

// RelationUnitSettings returns the relation unit settings for the given relation units in the remote model.
func (c *Client) RelationUnitSettings(relationUnits []params.RemoteRelationUnit) ([]params.SettingsResult, error) {
	var (
		args         params.RemoteRelationUnits
		retryIndices []int
	)

	result := make([]params.SettingsResult, len(relationUnits))
	args = params.RemoteRelationUnits{RelationUnits: relationUnits}

	// Use any previously cached discharge macaroons.
	for i, arg := range args.RelationUnits {
		if ms, ok := c.getCachedMacaroon("relation unit settings", arg.RelationToken); ok {
			newArg := arg
			newArg.Macaroons = ms
			args.RelationUnits[i] = newArg
		}
	}

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.SettingsResults
		err := c.facade.FacadeCall("RelationUnitSettings", args, &results)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if len(results.Results) != len(args.RelationUnits) {
			return nil, errors.Errorf("expected %d result(s), got %d", len(args.RelationUnits), len(results.Results))
		}

		if i == 0 {
			args = params.RemoteRelationUnits{}
			for j, res := range results.Results {
				if res.Error == nil {
					result[j] = res
					continue
				}
				mac, err := c.handleError(res.Error)
				if err != nil {
					res.Error.Message = err.Error()
					result[j] = res
					continue
				}
				retryArg := relationUnits[j]
				retryArg.Macaroons = mac
				args.RelationUnits = append(args.RelationUnits, retryArg)
				retryIndices = append(retryIndices, i)
				c.cache.Upsert(retryArg.RelationToken, mac)
			}
			if len(args.RelationUnits) == 0 {
				break
			}
		}
		if i == 1 {
			for j, res := range results.Results {
				result[retryIndices[j]] = res
			}
		}
	}
	return result, nil
}

// WatchEgressAddressesForRelation returns a watcher that notifies when addresses,
// from which connections will originate to the offering side of the relation, change.
// Each event contains the entire set of addresses which the offering side is required
// to allow for access to the other side of the relation.
func (c *Client) WatchEgressAddressesForRelation(remoteRelationArg params.RemoteEntityArg) (watcher.StringsWatcher, error) {
	args := params.RemoteEntityArgs{Args: []params.RemoteEntityArg{remoteRelationArg}}
	// Use any previously cached discharge macaroons.
	if ms, ok := c.getCachedMacaroon("watch relation egress addresses", remoteRelationArg.Token); ok {
		args.Args[0].Macaroons = ms
	}

	var result params.StringsWatchResult

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.StringsWatchResults
		err := c.facade.FacadeCall("WatchEgressAddressesForRelations", args, &results)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if len(results.Results) != 1 {
			return nil, errors.Errorf("expected 1 result, got %d", len(results.Results))
		}
		result = results.Results[0]
		if result.Error != nil && i == 0 {
			mac, err := c.handleError(result.Error)
			if err != nil {
				err = errors.Wrap(result.Error, err)
				return nil, err
			}
			args.Args[0].Macaroons = mac
			c.cache.Upsert(args.Args[0].Token, mac)
		}
	}
	w := apiwatcher.NewStringsWatcher(c.facade.RawAPICaller(), result)
	return w, nil
}

// WatchRelationStatus starts a RelationStatusWatcher for watching the life and
// status of the specified relation in the remote model.
func (c *Client) WatchRelationStatus(arg params.RemoteEntityArg) (watcher.RelationStatusWatcher, error) {
	args := params.RemoteEntityArgs{Args: []params.RemoteEntityArg{arg}}
	// Use any previously cached discharge macaroons.
	if ms, ok := c.getCachedMacaroon("watch relation status", arg.Token); ok {
		args.Args[0].Macaroons = ms
	}

	var result params.RelationStatusWatchResult

	// We make up to 2 api calls - the second is a retry after a macaroon discharge is obtained.
	for i := 0; i < 2; i++ {
		var results params.RelationStatusWatchResults
		err := c.facade.FacadeCall("WatchRelationsStatus", args, &results)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if len(results.Results) != 1 {
			return nil, errors.Errorf("expected 1 result, got %d", len(results.Results))
		}
		result = results.Results[0]
		if result.Error != nil && i == 0 {
			mac, err := c.handleError(result.Error)
			if err != nil {
				err = errors.Wrap(result.Error, err)
				return nil, err
			}
			args.Args[0].Macaroons = mac
			c.cache.Upsert(arg.Token, mac)
		}
	}
	w := apiwatcher.NewRelationStatusWatcher(c.facade.RawAPICaller(), result)
	return w, nil
}
