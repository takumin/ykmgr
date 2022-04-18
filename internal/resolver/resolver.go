package resolver

import (
	"fmt"
	"net/url"
)

type Resolver struct {
	network string
	address string
}

func Parse(uri string) (*Resolver, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	var r Resolver
	switch u.Scheme {
	case "tcp":
		r.network = "tcp"
		r.address = u.Host
	case "unix":
		r.network = "unix"
		r.address = u.Path
	default:
		return nil, fmt.Errorf("unsupported schema: %s", u.Scheme)
	}

	return &r, nil
}

func (r *Resolver) Network() string {
	return r.network
}

func (r *Resolver) Address() string {
	return r.address
}

func (r *Resolver) IsUnixDomainSocket() bool {
	return r.network == "unix"
}

func (r *Resolver) GrpcEndpoint() string {
	if r.network == "unix" {
		return "unix://" + r.address
	} else {
		return "dns://" + r.address
	}
}
