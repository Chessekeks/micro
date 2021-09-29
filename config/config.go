// Package config is an interface for dynamic configuration.
package config

import (
	"context"
	"errors"
	"time"
)

// DefaultConfig default config
var DefaultConfig Config = NewConfig()

// DefaultWatcherMinInterval default min interval for poll changes
var DefaultWatcherMinInterval = 5 * time.Second

// DefaultWatcherMaxInterval default max interval for poll changes
var DefaultWatcherMaxInterval = 9 * time.Second

var (
	// ErrCodecMissing is returned when codec needed and not specified
	ErrCodecMissing = errors.New("codec missing")
	// ErrInvalidStruct is returned when the target struct is invalid
	ErrInvalidStruct = errors.New("invalid struct specified")
	// ErrWatcherStopped is returned when source watcher has been stopped
	ErrWatcherStopped = errors.New("watcher stopped")
)

// Config is an interface abstraction for dynamic configuration
type Config interface {
	// Name returns name of config
	Name() string
	// Init the config
	Init(opts ...Option) error
	// Options in the config
	Options() Options
	// Load config from sources
	Load(context.Context, ...LoadOption) error
	// Save config to sources
	Save(context.Context, ...SaveOption) error
	// Watch a config for changes
	Watch(context.Context, ...WatchOption) (Watcher, error)
	// String returns config type name
	String() string
}

// Watcher is the config watcher
type Watcher interface {
	// Next blocks until update happens or error returned
	Next() (map[string]interface{}, error)
	// Stop stops watcher
	Stop() error
}

// Load loads config from config sources
func Load(ctx context.Context, cs []Config, opts ...LoadOption) error {
	var err error
	for _, c := range cs {
		if err = c.Init(); err != nil {
			return err
		}
		if err = c.Load(ctx, opts...); err != nil {
			return err
		}
	}
	return nil
}

var (
	DefaultAfterLoad = func(ctx context.Context, c Config) error {
		for _, fn := range c.Options().AfterLoad {
			if err := fn(ctx, c); err != nil {
				c.Options().Logger.Errorf(ctx, "%s AfterLoad err: %v", c.String(), err)
				if !c.Options().AllowFail {
					return err
				}
			}
		}
		return nil
	}

	DefaultAfterSave = func(ctx context.Context, c Config) error {
		for _, fn := range c.Options().AfterSave {
			if err := fn(ctx, c); err != nil {
				c.Options().Logger.Errorf(ctx, "%s AfterSave err: %v", c.String(), err)
				if !c.Options().AllowFail {
					return err
				}
			}
		}
		return nil
	}

	DefaultBeforeLoad = func(ctx context.Context, c Config) error {
		for _, fn := range c.Options().BeforeLoad {
			if err := fn(ctx, c); err != nil {
				c.Options().Logger.Errorf(ctx, "%s BeforeLoad err: %v", c.String(), err)
				if !c.Options().AllowFail {
					return err
				}
			}
		}
		return nil
	}

	DefaultBeforeSave = func(ctx context.Context, c Config) error {
		for _, fn := range c.Options().BeforeSave {
			if err := fn(ctx, c); err != nil {
				c.Options().Logger.Errorf(ctx, "%s BeforeSavec err: %v", c.String(), err)
				if !c.Options().AllowFail {
					return err
				}
			}
		}
		return nil
	}
)
