// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package rpc

import "context"

type contextKey string

const (
	tracingKey contextKey = "tracing"
)

// WithTracing returns a context with the given traceID and spanID.
func WithTracing(ctx context.Context, traceID, spanID string) context.Context {
	return context.WithValue(ctx, tracingKey, &trace{
		traceID: traceID,
		spanID:  spanID,
	})
}

// TracingFromContext returns the traceID and spanID from the context.
func TracingFromContext(ctx context.Context) (string, string) {
	val := ctx.Value(tracingKey)
	if val == nil {
		return "", ""
	}
	t, ok := val.(*trace)
	if !ok {
		return "", ""
	}
	return t.traceID, t.spanID
}

type trace struct {
	traceID string
	spanID  string
}
