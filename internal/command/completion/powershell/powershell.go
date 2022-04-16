package powershell

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/takumin/ykmgr/internal/config"
)

const powershellCompletion = `
$fn = $($MyInvocation.MyCommand.Name)
$name = $fn -replace "(.*)\.ps1$", '$1'
Register-ArgumentCompleter -Native -CommandName $name -ScriptBlock {
  param($commandName, $wordToComplete, $cursorPosition)
  $other = "$wordToComplete --generate-bash-completion"
  Invoke-Expression $other | ForEach-Object {
    [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
  }
}
`

func NewCommands(c *config.Config, f []cli.Flag) *cli.Command {
	return &cli.Command{
		Name:     "powershell",
		Usage:    "powershell completion",
		HideHelp: true,
		Action: func(ctx *cli.Context) error {
			fmt.Println(strings.TrimSpace(powershellCompletion))
			return nil
		},
	}
}
