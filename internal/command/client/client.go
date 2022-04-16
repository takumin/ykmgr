package client

import (
	"log"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/takumin/ykmgr/internal/config"
	"github.com/takumin/ykmgr/internal/protobuf/yubikey/v1"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	flags := []cli.Flag{}
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
		conn, err := grpc.DialContext(
			ctx.Context,
			c.Connection.Endpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			return err
		}

		client := yubikey.NewYubikeyServiceClient(conn)
		result, err := client.GetFirmwareVersion(
			ctx.Context,
			&yubikey.GetFirmwareVersionRequest{},
		)
		if err != nil {
			return err
		}

		for _, v := range result.FirmwareVersions {
			log.Printf("Firmware Version: %d.%d.%d\n", v.Major, v.Minor, v.Patch)
		}

		return nil
	}
}
