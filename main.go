package main

import (
	"os"

	"github.com/internetarchive/isodos-go/cmd"
	_ "github.com/internetarchive/isodos-go/cmd/all"
	"github.com/internetarchive/isodos-go/config"
	"github.com/internetarchive/isodos-go/setup"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Version - defined default version if it's not passed through flags during build
var Version string = "master"

func main() {
	app := cli.NewApp()
	app.Name = "isodos-go"
	app.Version = Version
	app.Author = "Corentin Barreau"
	app.Email = "corentin@archive.org"
	app.Usage = ""

	app.Flags = cmd.GlobalFlags
	app.Commands = cmd.Commands
	app.CommandNotFound = cmd.CommandNotFound
	app.Before = func(context *cli.Context) error {
		config.LoadConfig()

		if config.App.AppConfig.S3AccessKey == "" || config.App.AppConfig.S3SecretKey == "" || config.App.AppConfig.Project == "" {
			cont := false
			prompt := &survey.Confirm{
				Message: "It looks like Isodos isn't setup. Start the setup process?",
			}
			err := survey.AskOne(prompt, &cont, nil)
			if err != nil {
				return err
			}

			if !cont {
				logrus.Info("Exiting.")
				os.Exit(1)
			}

			setup.Setup()
		}

		if context.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}

		return nil
	}
	app.After = func(context *cli.Context) error {
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Panic(err)
	}
}
