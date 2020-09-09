package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	config "github.com/internetarchive/isodos-go/config"
)

var GlobalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "config",
		Required:    false,
		Value:       "~/.isodos.json",
		Usage:       "The config file to use",
		Destination: &config.App.AppConfigFile,
	},
	&cli.StringFlag{
		Name:        "s3-access-key",
		Required:    false,
		Usage:       "archive.org S3-Like Access Key",
		Destination: &config.App.Flags.S3AccessKey,
	},
	&cli.StringFlag{
		Name:        "s3-secret-key",
		Required:    false,
		Usage:       "archive.org S3-Like Secret Key",
		Destination: &config.App.Flags.S3SecretKey,
	},
	&cli.StringFlag{
		Name:        "project",
		Required:    false,
		Usage:       "Isodos project to use",
		Destination: &config.App.Flags.Project,
	},
	&cli.BoolFlag{
		Name:        "debug",
		Destination: &config.App.Flags.Debug,
	},
}

var Commands []cli.Command

func RegisterCommand(command cli.Command) {
	Commands = append(Commands, command)
}

func CommandNotFound(c *cli.Context, command string) {
	logrus.Errorf("%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
