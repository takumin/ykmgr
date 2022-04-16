package fish

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/takumin/ykmgr/internal/config"
)

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "fish",
		Usage:    "fish completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			fish, err := ctx.App.ToFishCompletion()
			if err != nil {
				return err
			}
			fmt.Println(fish)
			return nil
		},
	}
}
