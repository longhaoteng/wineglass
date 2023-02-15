package server

import (
	"context"

	"github.com/longhaoteng/wineglass/config"
)

type Options struct {
	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type Option func(*Options)

func Name(n string) Option {
	// config依赖，直接赋值
	config.Service.Name = n
	return func(o *Options) {}
}

func Version(v string) Option {
	return func(o *Options) {
		config.Service.Version = v
	}
}

func DBDriver(d string) Option {
	return func(o *Options) {
		config.DB.Driver = d
	}
}

func HttpAddr(h string) Option {
	return func(o *Options) {
		config.Service.HttpAddr = h
	}
}

// EnablePprof server enable pprof
func EnablePprof() Option {
	return func(o *Options) {
		config.Service.EnablePprof = true
	}
}

// DisableDB server not use db
func DisableDB() Option {
	return func(o *Options) {
		config.Service.DisableDB = true
	}
}

// DisableAuth server not use casbin auth, use auth must use db
func DisableAuth() Option {
	return func(o *Options) {
		config.Service.DisableAuth = true
	}
}

// DisableRedis server not use redis
func DisableRedis() Option {
	return func(o *Options) {
		config.Service.DisableRedis = true
	}
}

// BeforeStart run funcs before service starts
func BeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// BeforeStop run funcs before service stops
func BeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// AfterStart run funcs after service starts
func AfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// AfterStop run funcs after service stops
func AfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
