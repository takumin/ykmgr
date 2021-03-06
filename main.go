package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/takumin/ykmgr/internal/command/client"
	"github.com/takumin/ykmgr/internal/command/completion"
	"github.com/takumin/ykmgr/internal/command/server"
	"github.com/takumin/ykmgr/internal/config"
)

var (
	AppName  string = "ykmgr"
	Usage    string = "yubikey manager"
	Version  string = "unknown"
	Revision string = "unknown"
)

func main() {
	cfg := config.NewConfig(
		config.ServerListenURL("unix:///tmp/ykmgr.sock"),
		config.ClientEndpoint("unix:///tmp/ykmgr.sock"),
	)

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "log-level",
			Aliases:     []string{"loglv"},
			Usage:       "logging level",
			EnvVars:     []string{"LOG_LEVEL"},
			Value:       cfg.LogLevel,
			Destination: &cfg.LogLevel,
		},
	}

	cmds := []*cli.Command{
		completion.NewCommands(cfg, flags),
		server.NewCommands(cfg, flags),
		client.NewCommands(cfg, flags),
	}

	app := &cli.App{
		Name:                 AppName,
		Usage:                Usage,
		Version:              fmt.Sprintf("%s (%s)", Version, Revision),
		Flags:                flags,
		Commands:             cmds,
		EnableBashCompletion: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
