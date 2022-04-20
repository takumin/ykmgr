package server

import (
	"context"
	"fmt"
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
	"github.com/takumin/ykmgr/internal/resolver"
	"github.com/takumin/ykmgr/pkg/yubikey/v1"
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
		ykServer, err := NewServer()
		if err != nil {
			return err
		}
		defer ykServer.Close()

		grpcServer := grpc.NewServer(
			grpc.Creds(insecure.NewCredentials()),
		)
		yubikey.RegisterYubikeyServiceServer(grpcServer, ykServer)

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
			httpGrpcRouter(grpcServer, router),
			&http2.Server{},
		))
	}
}

type server struct {
	yubikey.UnimplementedYubikeyServiceServer

	yk map[uint32]*piv.YubiKey
}

func NewServer() (*server, error) {
	cards, err := piv.Cards()
	if err != nil {
		return nil, err
	}

	yubikeys := make(map[uint32]*piv.YubiKey, len(cards))
	for _, v := range cards {
		handle, err := piv.Open(v)
		if err != nil {
			return nil, err
		}

		serial, err := handle.Serial()
		if err != nil {
			return nil, err
		}

		yubikeys[serial] = handle
	}

	return &server{
		yk: yubikeys,
	}, nil
}

func (s *server) Close() map[uint32]error {
	err := make(map[uint32]error, len(s.yk))
	for k, v := range s.yk {
		err[k] = v.Close()
	}
	return err
}

func (s *server) GetSerials(ctx context.Context, req *yubikey.GetSerialsRequest) (*yubikey.GetSerialsResponse, error) {
	keys := make([]uint32, 0, len(s.yk))
	for k := range s.yk {
		keys = append(keys, k)
	}
	return &yubikey.GetSerialsResponse{
		Serials: keys,
	}, nil
}

func (s *server) GetVersion(ctx context.Context, req *yubikey.GetVersionRequest) (*yubikey.GetVersionResponse, error) {
	var version yubikey.Version
	if _, ok := s.yk[req.GetSerial()]; ok {
		version.Major = uint32(s.yk[req.GetSerial()].Version().Major)
		version.Minor = uint32(s.yk[req.GetSerial()].Version().Minor)
		version.Patch = uint32(s.yk[req.GetSerial()].Version().Patch)
	}
	return &yubikey.GetVersionResponse{
		Version: &version,
	}, nil
}

func (s *server) GetRetries(ctx context.Context, req *yubikey.GetRetriesRequest) (*yubikey.GetRetriesResponse, error) {
	if _, ok := s.yk[req.GetSerial()]; !ok {
		return nil, fmt.Errorf("not found serial")
	}
	retries, err := s.yk[req.GetSerial()].Retries()
	if err != nil {
		return nil, err
	}
	return &yubikey.GetRetriesResponse{
		Retries: uint32(retries),
	}, nil
}
