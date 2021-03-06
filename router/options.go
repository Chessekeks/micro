package router

import (
	"context"

	"github.com/google/uuid"
	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/register"
)

// Options are router options
type Options struct {
	Logger   logger.Logger
	Context  context.Context
	Register register.Register
	Name     string
	Gateway  string
	Network  string
	Id       string
	Address  string
	Precache bool
}

// Id sets Router Id
func Id(id string) Option {
	return func(o *Options) {
		o.Id = id
	}
}

// Address sets router service address
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

// Gateway sets network gateway
func Gateway(g string) Option {
	return func(o *Options) {
		o.Gateway = g
	}
}

// Network sets router network
func Network(n string) Option {
	return func(o *Options) {
		o.Network = n
	}
}

// Logger sets the logger
func Logger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

// Register sets the local register
func Register(r register.Register) Option {
	return func(o *Options) {
		o.Register = r
	}
}

// Precache the routes
func Precache() Option {
	return func(o *Options) {
		o.Precache = true
	}
}

// Name of the router
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// NewOptions returns router default options
func NewOptions(opts ...Option) Options {
	options := Options{
		Id:       uuid.New().String(),
		Network:  DefaultNetwork,
		Register: register.DefaultRegister,
		Logger:   logger.DefaultLogger,
		Context:  context.Background(),
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}
