package resolver_test

import (
	"testing"

	"github.com/takumin/ykmgr/internal/resolver"
)

func TestResolver(t *testing.T) {
	tests := []struct {
		uri                string
		err                bool
		network            string
		address            string
		grpcEndpoint       string
		isUnixDomainSocket bool
	}{
		//
		// Pass
		//
		{
			uri:                "tcp://127.0.0.1:8080",
			err:                false,
			network:            "tcp",
			address:            "127.0.0.1:8080",
			grpcEndpoint:       "dns://127.0.0.1:8080",
			isUnixDomainSocket: false,
		},
		{
			uri:                "tcp://::1:8080",
			err:                false,
			network:            "tcp",
			address:            "::1:8080",
			grpcEndpoint:       "dns://::1:8080",
			isUnixDomainSocket: false,
		},
		{
			uri:                "unix:///tmp/sock",
			err:                false,
			network:            "unix",
			address:            "/tmp/sock",
			grpcEndpoint:       "unix:///tmp/sock",
			isUnixDomainSocket: true,
		},
		//
		// Failed
		//
		{
			uri: "127.0.0.1",
			err: true,
		},
		{
			uri: "127.0.0.1:8080",
			err: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.uri, func(t *testing.T) {
			r, err := resolver.Parse(tt.uri)
			if tt.err {
				if err == nil {
					t.Fatal("want an error")
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if r.Network() != tt.network {
					t.Errorf("%v, want %v", r.Network(), tt.network)
				}
				if r.Address() != tt.address {
					t.Errorf("%v, want %v", r.Address(), tt.address)
				}
				if r.GrpcEndpoint() != tt.grpcEndpoint {
					t.Errorf("%v, want %v", r.GrpcEndpoint(), tt.grpcEndpoint)
				}
				if r.IsUnixDomainSocket() != tt.isUnixDomainSocket {
					t.Errorf("%v, want %v", r.IsUnixDomainSocket(), tt.isUnixDomainSocket)
				}
			}
		})
	}
}
