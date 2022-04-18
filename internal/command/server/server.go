package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/go-piv/piv-go/piv"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/takumin/ykmgr/internal/config"
	"github.com/takumin/ykmgr/internal/protobuf/yubikey/v1"
	"github.com/takumin/ykmgr/internal/resolver"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "listen-url",
			Aliases:     []string{"listen"},
			Usage:       "listen url",
			EnvVars:     []string{"LISTEN_URL", "LISTEN"},
			Value:       c.Server.ListenURL,
			Destination: &c.Server.ListenURL,
		},
	}
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s", "serv"},
		Usage:   "yubikey rpc server",
		Flags:   append(flags, f...),
		Action:  action(c),
	}
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}

func action(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		srv := grpc.NewServer(
			grpc.Creds(insecure.NewCredentials()),
		)
		yubikey.RegisterYubikeyServiceServer(srv, &server{})

		resolver, err := resolver.Parse(c.Server.ListenURL)
		if err != nil {
			return err
		}

		conn, err := grpc.DialContext(
			ctx.Context,
			resolver.GrpcEndpoint(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return err
		}

		router := runtime.NewServeMux()
		if err := yubikey.RegisterYubikeyServiceHandler(ctx.Context, router, conn); err != nil {
			return err
		}

		if resolver.IsUnixDomainSocket() {
			os.RemoveAll(resolver.Address())
		}

		listener, err := net.Listen(resolver.Network(), resolver.Address())
		if err != nil {
			return err
		}
		defer listener.Close()

		return http.Serve(listener, h2c.NewHandler(
			httpGrpcRouter(srv, router),
			&http2.Server{},
		))
	}
}

type server struct {
	yubikey.UnimplementedYubikeyServiceServer
}

func (s *server) GetVersions(ctx context.Context, req *yubikey.GetVersionsRequest) (*yubikey.GetVersionsResponse, error) {
	cards, err := piv.Cards()
	if err != nil {
		return nil, err
	}

	versions := make([]*yubikey.Version, len(cards))
	for i, v := range cards {
		yk, err := piv.Open(v)
		if err != nil {
			return nil, err
		}
		defer yk.Close()

		v := yk.Version()
		versions[i] = &yubikey.Version{
			Major: uint32(v.Major),
			Minor: uint32(v.Minor),
			Patch: uint32(v.Patch),
		}
	}

	return &yubikey.GetVersionsResponse{
		Versions: versions,
	}, nil
}

func (s *server) GetSerials(ctx context.Context, req *yubikey.GetSerialsRequest) (*yubikey.GetSerialsResponse, error) {
	cards, err := piv.Cards()
	if err != nil {
		return nil, err
	}

	serials := make([]uint32, len(cards))
	for i, v := range cards {
		yk, err := piv.Open(v)
		if err != nil {
			return nil, err
		}
		defer yk.Close()

		v, err := yk.Serial()
		if err != nil {
			return nil, err
		}

		serials[i] = v
	}

	return &yubikey.GetSerialsResponse{
		Serials: serials,
	}, nil
}
