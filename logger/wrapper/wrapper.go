// Package wrapper provides wrapper for Tracer
package wrapper

import (
	"context"

	"github.com/unistack-org/micro/v3/client"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/server"
)

type lWrapper struct {
	client.Client
	serverHandler    server.HandlerFunc
	serverSubscriber server.SubscriberFunc
	clientCallFunc   client.CallFunc
	opts             Options
}

type ClientCallObserver func(context.Context, client.Request, interface{}, []client.CallOption, error) []string
type ClientStreamObserver func(context.Context, client.Request, []client.CallOption, client.Stream, error) []string
type ClientPublishObserver func(context.Context, client.Message, []client.PublishOption, error) []string
type ClientCallFuncObserver func(context.Context, string, client.Request, interface{}, client.CallOptions, error) []string
type ServerHandlerObserver func(context.Context, server.Request, interface{}, error) []string
type ServerSubscriberObserver func(context.Context, server.Message, error) []string

type Options struct {
	Logger                    logger.Logger
	Level                     logger.Level
	Enabled                   bool
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
		Logger:                    logger.DefaultLogger,
		Level:                     logger.TraceLevel,
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

func WithEnabled(b bool) Option {
	return func(o *Options) {
		o.Enabled = b
	}
}

func WithLevel(l logger.Level) Option {
	return func(o *Options) {
		o.Level = l
	}
}

func WithLogger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
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

func DefaultClientCallObserver(ctx context.Context, req client.Request, rsp interface{}, opts []client.CallOption, err error) []string {
	labels := []string{"service", req.Service(), "endpoint", req.Endpoint()}
	if err != nil {
		labels = append(labels, "error", err.Error())
	}
	return labels
}

func DefaultClientStreamObserver(ctx context.Context, req client.Request, opts []client.CallOption, stream client.Stream, err error) []string {
	labels := []string{"service", req.Service(), "endpoint", req.Endpoint()}
	if err != nil {
		labels = append(labels, "error", err.Error())
	}
	return labels
}

func DefaultClientPublishObserver(ctx context.Context, msg client.Message, opts []client.PublishOption, err error) []string {
	labels := []string{"endpoint", msg.Topic()}
	if err != nil {
		labels = append(labels, "error", err.Error())
	}
	return labels
}

func DefaultServerHandlerObserver(ctx context.Context, req server.Request, rsp interface{}, err error) []string {
	labels := []string{"service", req.Service(), "endpoint", req.Endpoint()}
	if err != nil {
		labels = append(labels, "error", err.Error())
	}
	return labels
}

func DefaultServerSubscriberObserver(ctx context.Context, msg server.Message, err error) []string {
	labels := []string{"endpoint", msg.Topic()}
	if err != nil {
		labels = append(labels, "error", err.Error())
	}
	return labels
}

func DefaultClientCallFuncObserver(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions, err error) []string {
	labels := []string{"service", req.Service(), "endpoint", req.Endpoint()}
	if err != nil {
		labels = append(labels, "error", err.Error())
	}
	return labels
}

func (l *lWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	err := l.Client.Call(ctx, req, rsp, opts...)

	if !l.opts.Enabled {
		return err
	}

	var labels []string
	for _, o := range l.opts.ClientCallObservers {
		labels = append(labels, o(ctx, req, rsp, opts, err)...)
	}
	fields := make(map[string]interface{}, int(len(labels)/2))
	for i := 0; i < len(labels); i += 2 {
		fields[labels[i]] = labels[i+1]
	}
	l.opts.Logger.Fields(fields).Log(ctx, l.opts.Level)

	return err
}

func (l *lWrapper) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	stream, err := l.Client.Stream(ctx, req, opts...)

	if !l.opts.Enabled {
		return stream, err
	}

	var labels []string
	for _, o := range l.opts.ClientStreamObservers {
		labels = append(labels, o(ctx, req, opts, stream, err)...)
	}
	fields := make(map[string]interface{}, int(len(labels)/2))
	for i := 0; i < len(labels); i += 2 {
		fields[labels[i]] = labels[i+1]
	}
	l.opts.Logger.Fields(fields).Log(ctx, l.opts.Level)

	return stream, err
}

func (l *lWrapper) Publish(ctx context.Context, msg client.Message, opts ...client.PublishOption) error {
	err := l.Client.Publish(ctx, msg, opts...)

	if !l.opts.Enabled {
		return err
	}

	var labels []string
	for _, o := range l.opts.ClientPublishObservers {
		labels = append(labels, o(ctx, msg, opts, err)...)
	}
	fields := make(map[string]interface{}, int(len(labels)/2))
	for i := 0; i < len(labels); i += 2 {
		fields[labels[i]] = labels[i+1]
	}
	l.opts.Logger.Fields(fields).Log(ctx, l.opts.Level)

	return err
}

func (l *lWrapper) ServerHandler(ctx context.Context, req server.Request, rsp interface{}) error {
	err := l.serverHandler(ctx, req, rsp)

	if !l.opts.Enabled {
		return err
	}

	var labels []string
	for _, o := range l.opts.ServerHandlerObservers {
		labels = append(labels, o(ctx, req, rsp, err)...)
	}
	fields := make(map[string]interface{}, int(len(labels)/2))
	for i := 0; i < len(labels); i += 2 {
		fields[labels[i]] = labels[i+1]
	}
	l.opts.Logger.Fields(fields).Log(ctx, l.opts.Level)

	return err
}

func (l *lWrapper) ServerSubscriber(ctx context.Context, msg server.Message) error {
	err := l.serverSubscriber(ctx, msg)

	if !l.opts.Enabled {
		return err
	}

	var labels []string
	for _, o := range l.opts.ServerSubscriberObservers {
		labels = append(labels, o(ctx, msg, err)...)
	}
	fields := make(map[string]interface{}, int(len(labels)/2))
	for i := 0; i < len(labels); i += 2 {
		fields[labels[i]] = labels[i+1]
	}
	l.opts.Logger.Fields(fields).Log(ctx, l.opts.Level)

	return err
}

// NewClientWrapper accepts an open tracing Trace and returns a Client Wrapper
func NewClientWrapper(opts ...Option) client.Wrapper {
	return func(c client.Client) client.Client {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}
		return &lWrapper{opts: options, Client: c}
	}
}

// NewClientCallWrapper accepts an opentracing Tracer and returns a Call Wrapper
func NewClientCallWrapper(opts ...Option) client.CallWrapper {
	return func(h client.CallFunc) client.CallFunc {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}

		l := &lWrapper{opts: options, clientCallFunc: h}
		return l.ClientCallFunc
	}
}

func (l *lWrapper) ClientCallFunc(ctx context.Context, addr string, req client.Request, rsp interface{}, opts client.CallOptions) error {
	err := l.clientCallFunc(ctx, addr, req, rsp, opts)

	if !l.opts.Enabled {
		return err
	}

	var labels []string
	for _, o := range l.opts.ClientCallFuncObservers {
		labels = append(labels, o(ctx, addr, req, rsp, opts, err)...)
	}
	fields := make(map[string]interface{}, int(len(labels)/2))
	for i := 0; i < len(labels); i += 2 {
		fields[labels[i]] = labels[i+1]
	}
	l.opts.Logger.Fields(fields).Log(ctx, l.opts.Level)

	return err
}

// NewServerHandlerWrapper accepts an options and returns a Handler Wrapper
func NewServerHandlerWrapper(opts ...Option) server.HandlerWrapper {
	return func(h server.HandlerFunc) server.HandlerFunc {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}

		l := &lWrapper{opts: options, serverHandler: h}
		return l.ServerHandler
	}
}

// NewServerSubscriberWrapper accepts an opentracing Tracer and returns a Subscriber Wrapper
func NewServerSubscriberWrapper(opts ...Option) server.SubscriberWrapper {
	return func(h server.SubscriberFunc) server.SubscriberFunc {
		options := NewOptions()
		for _, o := range opts {
			o(&options)
		}

		l := &lWrapper{opts: options, serverSubscriber: h}
		return l.ServerSubscriber
	}
}
