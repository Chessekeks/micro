package token

import (
	"time"

	"github.com/unistack-org/micro/v3/store"
)

// Options holds the options for token
type Options struct {
	// Store to persist the tokens
	Store store.Store
	// PublicKey base64 encoded, used by JWT
	PublicKey string
	// PrivateKey base64 encoded, used by JWT
	PrivateKey string
}

// Option func signature
type Option func(o *Options)

// WithStore sets the token providers store
func WithStore(s store.Store) Option {
	return func(o *Options) {
		o.Store = s
	}
}

// WithPublicKey sets the JWT public key
func WithPublicKey(key string) Option {
	return func(o *Options) {
		o.PublicKey = key
	}
}

// WithPrivateKey sets the JWT private key
func WithPrivateKey(key string) Option {
	return func(o *Options) {
		o.PrivateKey = key
	}
}

// NewOptions returns options struct filled by opts
func NewOptions(opts ...Option) Options {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}
	// set default store
	if options.Store == nil {
		options.Store = store.DefaultStore
	}
	return options
}

// GenerateOptions holds the generate options
type GenerateOptions struct {
	// Expiry for the token
	Expiry time.Duration
}

// GenerateOption func signature
type GenerateOption func(o *GenerateOptions)

// WithExpiry for the generated account's token expires
func WithExpiry(d time.Duration) GenerateOption {
	return func(o *GenerateOptions) {
		o.Expiry = d
	}
}

// NewGenerateOptions from a slice of options
func NewGenerateOptions(opts ...GenerateOption) GenerateOptions {
	var options GenerateOptions
	for _, o := range opts {
		o(&options)
	}
	// set default Expiry of token
	if options.Expiry == 0 {
		options.Expiry = time.Minute * 15
	}
	return options
}
