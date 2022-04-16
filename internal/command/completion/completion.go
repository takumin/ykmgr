package completion

import (
	"github.com/urfave/cli/v2"

	"github.com/takumin/ykmgr/internal/command/completion/bash"
	"github.com/takumin/ykmgr/internal/command/completion/fish"
	"github.com/takumin/ykmgr/internal/command/completion/powershell"
	"github.com/takumin/ykmgr/internal/command/completion/zsh"
	"github.com/takumin/ykmgr/internal/config"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	cmds := []*cli.Command{
		bash.NewCommands(c, f),
		fish.NewCommands(c, f),
		zsh.NewCommands(c, f),
		powershell.NewCommands(c, f),
	}
	return &cli.Command{
		Name:        "completion",
		Usage:       "command completion",
		Subcommands: cmds,
		HideHelp:    true,
	}
}
