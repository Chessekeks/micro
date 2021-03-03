// Package wrapper provides wrapper for Tracer
package wrapper

import (
	"context"
	"fmt"

	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/metadata"
	"github.com/unistack-org/micro/v3/server"
	"github.com/unistack-org/micro/v3/tracer"
)

type tWrapper struct {
	opts             Options
	serverHandler    server.HandlerFunc
	serverSubscriber server.SubscriberFunc
	clientCallFunc   client.CallFunc
	client.Client
}

type ClientCallObserver func(context.Context, client.Request, interface{}, []client.CallOption, tracer.Span, error)
type ClientStreamObserver func(context.Context, client.Request, []client.CallOption, client.Stream, tracer.Span, error)
type ClientPublishObserver func(context.Context, client.Message, []client.PublishOption, tracer.Span, error)
type ClientCallFuncObserver func(context.Context, string, client.Request, interface{}, client.CallOptions, tracer.Span, error)
type ServerHandlerObserver func(context.Context, server.Request, interface{}, tracer.Span, error)
type ServerSubscriberObserver func(context.Context, server.Message, tracer.Span, error)

type Options struct {
	Tracer                    tracer.Tracer
	ClientCallObservers       []ClientCallObserver
	ClientStreamObservers     []ClientStreamObserver
	ClientPublishObservers    []ClientPublishObserver
	ClientCallFuncObservers   []ClientCallFuncObserver
	ServerHandlerObservers    []ServerHandlerObserver
	ServerSubscriberObservers []ServerSubscriberObserver
}

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Tracer:                    tracer.DefaultTracer,
		ClientCallObservers:       []ClientCallObserver{DefaultClientCallObserver},
		ClientStreamObservers:     []ClientStreamObserver{DefaultClientStreamObserver},
		ClientPublishObservers:    []ClientPublishObserver{DefaultClientPublishObserver},
		ClientCallFuncObservers:   []ClientCallFuncObserver{DefaultClientCallFuncObserver},
		ServerHandlerObservers:    []ServerHandlerObserver{DefaultServerHandlerObserver},
		ServerSubscriberObservers: []ServerSubscriberObserver{DefaultServerSubscriberObserver},
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithTracer(t tracer.Tracer) Option {
	return func(o *Options) {
		o.Tracer = t
	}
}

func WithClientCallObservers(ob ...ClientCallObserver) Option {
	return func(o *Options) {
		o.ClientCallObservers = ob
	}
}

func WithClientStreamObservers(ob ...ClientStreamObserver) Option {
	return func(o *Options) {
		o.ClientStreamObservers = ob
	}
}

func WithClientPublishObservers(ob ...ClientPublishObserver) Option {
	return func(o *Options) {
		o.ClientPublishObservers = ob
	}
}

func WithClientCallFuncObservers(ob ...ClientCallFuncObserver) Option {
	return func(o *Options) {
		o.ClientCallFuncObservers = ob
	}
}

func WithServerHandlerObservers(ob ...ServerHandlerObserver) Option {
	return func(o *Options) {
		o.ServerHandlerObservers = ob
	}
}

func WithServerSubscriberObservers(ob ...ServerSubscriberObserver) Option {
	return func(o *Options) {
		o.ServerSubscriberObservers = ob
	}
}

func DefaultClientCallObserver(ctx context.Context, req client.Request, rsp interface{}, opts []client.CallOption, sp tracer.Span, err error) {
	sp.SetName(fmt.Sprintf("%s.%s", req.Service(), req.Endpoint()))
	var labels []tracer.Label
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		labels = make([]tracer.Label, 0, len(md))
		for k, v := range md {
			labels = append(labels, tracer.String(k, v))
		}
	}
	if err != nil {
		labels = append(labels, tracer.Bool("error", true))
	}
	sp.SetLabels(labels...)
}

func DefaultClientStreamObserver(ctx context.Context, req client.Request, opts []client.CallOption, stream client.Stream, sp tracer.Span, err error) {
	sp.SetName(fmt.Sprintf("%s.%s", req.Service(), req.Endpoint()))
	var labels []tracer.Label
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		labels = make([]tracer.Label, 0, len(md))
		for k, v := range md {
			labels = append(labels, tracer.String(k, v))
		}
	}
	if err != nil {
		labels = append(labels, tracer.Bool("error", true))
	}
	sp.SetLabels(labels...)
}

func DefaultClientPublishObserver(ctx context.Context, msg client.Message, opts []client.PublishOption, sp tracer.Span, err error) {
	sp.SetName(fmt.Sprintf("Pub to %s", msg.Topic()))
	var labels []tracer.Label
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		labels = make([]tracer.Label, 0, len(md))
		for k, v := range md {
			labels = append(labels, tracer.String(k, v))
		}
	}
	if err != nil {
		labels = append(labels, tracer.Bool("error", true))
	}
	sp.SetLabels(labels...)
}

func DefaultServerHandlerObserver(ctx context.Context, req server.Request, rsp interface{}, sp tracer.Span, err error) {
	sp.SetName(fmt.Sprintf("%s.%s", req.Service(), req.Endpoint()))
	var labels []tracer.Label
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		labels = make([]tracer.Label, 0, len(md))
		for k, v := range md {
			labels = append(labels, tracer.String(k, v))
		}
	}
	if err != nil {
		labels = append(labels, tracer.Bool("error", true))
	}
	sp.SetLabels(labels...)
}

func DefaultServerSubscriberObserver(ctx context.Context, msg server.Message, sp tracer.Span, err error) {
	sp.SetName(fmt.Sprintf("Sub from %s", msg.Topic()))
	var labels []tracer.Label
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		labels = make([]tracer.Label, 0, len(md))
		for k, v := range md {
			labels = append(labels, tracer.String(k, v))
		}
	}
	if err != nil {
		labels = append(labels, tracer.Bool("error", true))
	}
	sp.SetLabels(labels...)
}

func DefaultClientCallFuncObserver(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions, sp tracer.Span, err error) {
	sp.SetName(fmt.Sprintf("%s.%s", req.Service(), req.Endpoint()))
	var labels []tracer.Label
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		labels = make([]tracer.Label, 0, len(md))
		for k, v := range md {
			labels = append(labels, tracer.String(k, v))
		}
	}
	if err != nil {
		labels = append(labels, tracer.Bool("error", true))
	}
	sp.SetLabels(labels...)
}

func (ot *tWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	sp := tracer.SpanFromContext(ctx)
	defer sp.Finish()

	err := ot.Client.Call(ctx, req, rsp, opts...)

	for _, o := range ot.opts.ClientCallObservers {
		o(ctx, req, rsp, opts, sp, err)
	}

	return err
}

func (ot *tWrapper) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	sp := tracer.SpanFromContext(ctx)
	defer sp.Finish()

	stream, err := ot.Client.Stream(ctx, req, opts...)

	for _, o := range ot.opts.ClientStreamObservers {
		o(ctx, req, opts, stream, sp, err)
	}

	return stream, err
}

func (ot *tWrapper) Publish(ctx context.Context, msg client.Message, opts ...client.PublishOption) error {
	sp := tracer.SpanFromContext(ctx)
	defer sp.Finish()

	err := ot.Client.Publish(ctx, msg, opts...)

	for _, o := range ot.opts.ClientPublishObservers {
		o(ctx, msg, opts, sp, err)
	}

	return err
}

func (ot *tWrapper) ServerHandler(ctx context.Context, req server.Request, rsp interface{}) error {
	sp := tracer.SpanFromContext(ctx)
	defer sp.Finish()

	err := ot.serverHandler(ctx, req, rsp)

	for _, o := range ot.opts.ServerHandlerObservers {
		o(ctx, req, rsp, sp, err)
	}

	return err
}

func (ot *tWrapper) ServerSubscriber(ctx context.Context, msg server.Message) error {
	sp := tracer.SpanFromContext(ctx)
	defer sp.Finish()

	err := ot.serverSubscriber(ctx, msg)

	for _, o := range ot.opts.ServerSubscriberObservers {
		o(ctx, msg, sp, err)
	}

	return err
}

// NewClientWrapper accepts an open tracing Trace and returns a Client Wrapper
func NewClientWrapper(opts ...Option) client.Wrapper {
	return func(c client.Client) client.Client {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}
		return &tWrapper{opts: options, Client: c}
	}
}

// NewClientCallWrapper accepts an opentracing Tracer and returns a Call Wrapper
func NewClientCallWrapper(opts ...Option) client.CallWrapper {
	return func(h client.CallFunc) client.CallFunc {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}

		ot := &tWrapper{opts: options, clientCallFunc: h}
		return ot.ClientCallFunc
	}
}

func (ot *tWrapper) ClientCallFunc(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
	sp := tracer.SpanFromContext(ctx)
	defer sp.Finish()

	err := ot.clientCallFunc(ctx, addr, req, rsp, opts)

	for _, o := range ot.opts.ClientCallFuncObservers {
		o(ctx, addr, req, rsp, opts, sp, err)
	}

	return err
}

// NewServerHandlerWrapper accepts an options and returns a Handler Wrapper
func NewServerHandlerWrapper(opts ...Option) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}

		ot := &tWrapper{opts: options, serverHandler: h}
		return ot.ServerHandler
	}
}

// NewServerSubscriberWrapper accepts an opentracing Tracer and returns a Subscriber Wrapper
func NewServerSubscriberWrapper(opts ...Option) server.SubscriberWrapper {
	return func(h server.SubscriberFunc) server.SubscriberFunc {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}

		ot := &tWrapper{opts: options, serverSubscriber: h}
		return ot.ServerSubscriber
	}
}
