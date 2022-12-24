// Package tracer provides an interface for distributed tracing
package tracer // import "go.unistack.org/micro/v3/tracer"

import (
	"context"
)

// DefaultTracer is the global default tracer
var DefaultTracer = NewTracer()

// Tracer is an interface for distributed tracing
type Tracer interface {
	// Name return tracer name
	Name() string
	// Init tracer with options
	Init(...Option) error
	// Start a trace
	Start(ctx context.Context, name string, opts ...SpanOption) (context.Context, Span)
}

type Span interface {
	// Tracer return underlining tracer
	Tracer() Tracer
	// Finish complete and send span
	Finish(opts ...SpanOption)
	// AddEvent add event to span
	AddEvent(name string, opts ...EventOption)
	// Context return context with span
	Context() context.Context
	// SetName set the span name
	SetName(name string)
	// SetLabels set the span labels
	SetLabels(labels ...interface{})
	// AddLabels append the span labels
	AddLabels(labels ...interface{})
}
