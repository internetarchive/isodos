package send

import (
	"github.com/internetarchive/isodos-go/cmd"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.RegisterCommand(
		cli.Command{
			Name:  "send",
			Usage: "All commands that sends something",
			Subcommands: []*cli.Command{
				NewSendListCommand(),
			},
		})
}
