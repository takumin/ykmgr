package client

import (
	"log"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/takumin/ykmgr/internal/config"
	"github.com/takumin/ykmgr/internal/resolver"
	"github.com/takumin/ykmgr/pkg/yubikey/v1"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "endpoint",
			Aliases:     []string{"ep"},
			Usage:       "connect endpoint url",
			EnvVars:     []string{"ENDPOINT"},
			Value:       c.Client.Endpoint,
			Destination: &c.Client.Endpoint,
		},
	}
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c", "cli"},
		Usage:   "yubikey rpc client",
		Flags:   append(flags, f...),
		Action:  action(c),
	}
}

func action(c *config.Config) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
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

		client := yubikey.NewYubikeyServiceClient(conn)

		versions, err := client.GetVersions(
			ctx.Context,
			&yubikey.GetVersionsRequest{},
		)
		if err != nil {
			return err
		}

		for _, v := range versions.Versions {
			log.Printf("Version: %d.%d.%d\n", v.Major, v.Minor, v.Patch)
		}

		serials, err := client.GetSerials(
			ctx.Context,
			&yubikey.GetSerialsRequest{},
		)
		if err != nil {
			return err
		}

		for _, v := range serials.Serials {
			log.Printf("Serial: %d\n", v)
		}

		return nil
	}
}
