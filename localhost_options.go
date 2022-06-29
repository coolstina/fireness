package fireness

// LocalhostOption implement Option model.
type LocalhostOption interface {
	apply(option *localhostOption)
}

// LocalhostOptionFunc Options function structure.
type LocalhostOptionFunc func(option *localhostOption)

func (l LocalhostOptionFunc) apply(option *localhostOption) {
	l(option)
}

type localhostOption struct {
	excludes    []string
	excludeIPv6 bool
	excludeIPv4 bool
}

// WithLocalhostExclude Get IP list exclude specify IP.
func WithLocalhostExclude(excludes ...string) LocalhostOption {
	return LocalhostOptionFunc(func(option *localhostOption) {
		option.excludes = excludes
	})
}

// WithLocalhostExcludeIPv6 Get IP list exclude IPv6 address.
func WithLocalhostExcludeIPv6(exclude bool) LocalhostOption {
	return LocalhostOptionFunc(func(option *localhostOption) {
		option.excludeIPv6 = exclude
	})
}

// WithLocalhostExcludeIPv4 Get IP list exclude IPv4 address.
func WithLocalhostExcludeIPv4(exclude bool) LocalhostOption {
	return LocalhostOptionFunc(func(option *localhostOption) {
		option.excludeIPv4 = exclude
	})
}
