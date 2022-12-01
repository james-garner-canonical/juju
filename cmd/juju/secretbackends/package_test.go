// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretbackends

import (
	stdtesting "testing"

	gc "gopkg.in/check.v1"

	"github.com/juju/juju/jujuclient"
)

//go:generate go run github.com/golang/mock/mockgen -package mocks -destination mocks/secretbackendsapi.go github.com/juju/juju/cmd/juju/secretbackends ListSecretBackendsAPI,AddSecretBackendsAPI

func TestPackage(t *stdtesting.T) {
	gc.TestingT(t)
}

// NewListCommandForTest returns a secret backends command for testing.
func NewListCommandForTest(store jujuclient.ClientStore, listSecretsAPI ListSecretBackendsAPI) *listSecretBackendsCommand {
	c := &listSecretBackendsCommand{
		listSecretBackendsAPIFunc: func() (ListSecretBackendsAPI, error) { return listSecretsAPI, nil },
	}
	c.SetClientStore(store)
	return c
}

// NewAddCommandForTest returns an add secret backends command for testing.
func NewAddCommandForTest(store jujuclient.ClientStore, addSecretBackendsAPI AddSecretBackendsAPI) *addSecretBackendCommand {
	c := &addSecretBackendCommand{
		AddSecretBackendsAPIFunc: func() (AddSecretBackendsAPI, error) { return addSecretBackendsAPI, nil },
	}
	c.SetClientStore(store)
	return c
}
