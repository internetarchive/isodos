package setup

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/AlecAivazis/survey.v1"

	config "github.com/internetarchive/isodos-go/config"
)

func Setup() {
	stringResponse := ""

	input := &survey.Input{
		Message: "What is your archive.org S3 access key? (see: https://archive.org/account/s3.php)",
	}

	err := survey.AskOne(input, &stringResponse, nil)
	if err != nil {
		logrus.Panic(err)
	}

	if len(stringResponse) != 16 {
		logrus.Fatalln("Invalid S3 access key (need to be 16 chars long)")
	}

	config.App.AppConfig.S3AccessKey = stringResponse

	input = &survey.Input{
		Message: "What is your archive.org S3 secret key? (see: https://archive.org/account/s3.php)",
	}

	err = survey.AskOne(input, &stringResponse, nil)
	if err != nil || len(stringResponse) != 16 {
		logrus.Panic(err)
	}

	if len(stringResponse) != 16 {
		logrus.Fatalln("Invalid S3 secret key (need to be 16 chars long)")
	}

	config.App.AppConfig.S3SecretKey = stringResponse

	input = &survey.Input{
		Message: "Where should the Isodos config file be saved?",
		Default: "~/",
	}

	err = survey.AskOne(input, &stringResponse, nil)
	if err != nil {
		logrus.Panic(err)
	}

	config.App.AppConfig.ConfigFilePath = stringResponse

	input = &survey.Input{
		Message: "How do you want the Isodos config file named?",
		Default: ".isodos.json",
	}

	err = survey.AskOne(input, &stringResponse, nil)
	if err != nil {
		logrus.Panic(err)
	}

	config.App.AppConfig.ConfigFileName = stringResponse

	input = &survey.Input{
		Message: "What's the name of the Isodos project you want to use?",
	}

	err = survey.AskOne(input, &stringResponse, nil)
	if err != nil {
		logrus.Panic(err)
	}

	config.App.AppConfig.Project = stringResponse

	config.SaveConfig(config.App.AppConfig)
}
