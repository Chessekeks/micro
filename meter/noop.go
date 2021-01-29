package meter

import (
	"io"
	"time"
)

// NoopMeter is an noop implementation of Meter
type noopMeter struct {
	opts   Options
	labels Labels
}

// NewMeter returns a configured noop reporter:
func NewMeter(opts ...Option) Meter {
	return &noopMeter{opts: NewOptions(opts...)}
}

func (r *noopMeter) Name() string {
	return r.opts.Name
}

// Init initialize options
func (r *noopMeter) Init(opts ...Option) error {
	for _, o := range opts {
		o(&r.opts)
	}
	return nil
}

// Counter implements the Meter interface
func (r *noopMeter) Counter(name string, opts ...Option) Counter {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	return &noopCounter{labels: options.Labels}
}

// FloatCounter implements the Meter interface
func (r *noopMeter) FloatCounter(name string, opts ...Option) FloatCounter {
	return &noopFloatCounter{}
}

// Gauge implements the Meter interface
func (r *noopMeter) Gauge(name string, f func() float64, opts ...Option) Gauge {
	return &noopGauge{}
}

// Summary implements the Meter interface
func (r *noopMeter) Summary(name string, opts ...Option) Summary {
	return &noopSummary{}
}

// SummaryExt implements the Meter interface
func (r *noopMeter) SummaryExt(name string, window time.Duration, quantiles []float64, opts ...Option) Summary {
	return &noopSummary{}
}

// Histogram implements the Meter interface
func (r *noopMeter) Histogram(name string, opts ...Option) Histogram {
	return &noopHistogram{}
}

// Set implements the Meter interface
func (r *noopMeter) Set(opts ...Option) Meter {
	m := &noopMeter{opts: r.opts}

	for _, o := range opts {
		o(&m.opts)
	}

	return m
}

func (r *noopMeter) Write(w io.Writer, withProcessMetrics bool) error {
	return nil
}

// Options implements the Meter interface
func (r *noopMeter) Options() Options {
	return r.opts
}

// String implements the Meter interface
func (r *noopMeter) String() string {
	return "noop"
}

type noopCounter struct {
	labels Labels
}

func (r *noopCounter) Add(int) {

}

func (r *noopCounter) Dec() {

}

func (r *noopCounter) Get() uint64 {
	return 0
}

func (r *noopCounter) Inc() {

}

func (r *noopCounter) Set(uint64) {

}

type noopFloatCounter struct{}

func (r *noopFloatCounter) Add(float64) {

}

func (r *noopFloatCounter) Get() float64 {
	return 0
}

func (r *noopFloatCounter) Set(float64) {

}

func (r *noopFloatCounter) Sub(float64) {

}

type noopGauge struct{}

func (r *noopGauge) Get() float64 {
	return 0
}

type noopSummary struct{}

func (r *noopSummary) Update(float64) {

}

func (r *noopSummary) UpdateDuration(time.Time) {

}

type noopHistogram struct{}

func (r *noopHistogram) Reset() {

}

func (r *noopHistogram) Update(float64) {

}

func (r *noopHistogram) UpdateDuration(time.Time) {

}

//func (r *noopHistogram) VisitNonZeroBuckets(f func(vmrange string, count uint64)) {}
