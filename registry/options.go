package registry

import "time"

// Options ...
type Options struct {
	Addrs        []string
	Timeout      time.Duration
	HeartBeat    int64
	RegistryPath string
}

// Option ...
type Option func(opts *Options)

// WithAddrs ...
func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

// WithTimeout ...
func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

// WithHeartBeat ...
func WithHeartBeat(heartbeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartbeat
	}
}

// WithRegistryPath ...
func WithRegistryPath(registrypath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registrypath
	}
}
