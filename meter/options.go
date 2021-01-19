package meter

import (
	"context"

	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/metadata"
)

var (
	// The Prometheus metrics will be made available on this port:
	defaultPrometheusListenAddress = ":9000"
	// This is the endpoint where the Prometheus metrics will be made available ("/metrics" is the default with Prometheus):
	defaultPath = "/metrics"
	// timingObjectives is the default spread of stats we maintain for timings / histograms:
	defaultTimingObjectives = map[float64]float64{0.0: 0, 0.5: 0.05, 0.75: 0.04, 0.90: 0.03, 0.95: 0.02, 0.98: 0.001, 1: 0}
)

// Option powers the configuration for metrics implementations:
type Option func(*Options)

// Options for metrics implementations:
type Options struct {
	Address          string
	Path             string
	DefaultTags      metadata.Metadata
	TimingObjectives map[float64]float64
	Logger           logger.Logger
	Context          context.Context
}

// NewOptions prepares a set of options:
func NewOptions(opt ...Option) Options {
	opts := Options{
		Address:          defaultPrometheusListenAddress,
		DefaultTags:      metadata.New(2),
		Path:             defaultPath,
		TimingObjectives: defaultTimingObjectives,
		Context:          context.Background(),
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

// Cntext sets the metrics context
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// Path used to serve metrics over HTTP:
func Path(value string) Option {
	return func(o *Options) {
		o.Path = value
	}
}

// Address is the listen address to serve metrics on:
func Address(value string) Option {
	return func(o *Options) {
		o.Address = value
	}
}

// DefaultTags will be added to every metric:
func DefaultTags(md metadata.Metadata) Option {
	return func(o *Options) {
		o.DefaultTags = metadata.Copy(md)
	}
}

// TimingObjectives defines the desired spread of statistics for histogram / timing metrics:
func TimingObjectives(value map[float64]float64) Option {
	return func(o *Options) {
		o.TimingObjectives = value
	}
}

// Logger sets the logger
func Logger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}
