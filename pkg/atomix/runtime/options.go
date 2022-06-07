// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package runtime

import "github.com/atomix/runtime/pkg/atomix/primitive"

type Options struct {
	ProxyService   ProxyServiceOptions
	ControlService ControlServiceOptions
	ConfigFile     string
	CacheDir       string
}

func (o Options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(&o)
	}
}

type Option func(*Options)

type ServerOptions struct {
	Host string
	Port int
}

type ProxyServiceOptions struct {
	ServerOptions
	Kinds []primitive.Kind
}

type ControlServiceOptions struct {
	ServerOptions
}

func WithOptions(opts Options) Option {
	return func(options *Options) {
		*options = opts
	}
}

func WithProxyHost(host string) Option {
	return func(options *Options) {
		options.ProxyService.Host = host
	}
}

func WithProxyPort(port int) Option {
	return func(options *Options) {
		options.ProxyService.Port = port
	}
}

func WithProxyKinds(kinds ...primitive.Kind) Option {
	return func(options *Options) {
		options.ProxyService.Kinds = append(options.ProxyService.Kinds, kinds...)
	}
}

func WithControlHost(host string) Option {
	return func(options *Options) {
		options.ControlService.Host = host
	}
}

func WithControlPort(port int) Option {
	return func(options *Options) {
		options.ControlService.Port = port
	}
}

func WithConfigFile(file string) Option {
	return func(options *Options) {
		options.ConfigFile = file
	}
}

func WithCacheDir(dir string) Option {
	return func(options *Options) {
		options.CacheDir = dir
	}
}