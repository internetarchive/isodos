package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type Flags struct {
	AppConfig
	S3AccessKey string
	S3SecretKey string
	Project     string
	Debug       bool
}

type AppConfig struct {
	S3AccessKey    string
	S3SecretKey    string
	Project        string
	ConfigFilePath string
	ConfigFileName string
}

type Application struct {
	AppConfigFile string
	AppConfig     AppConfig
	Flags         Flags
}

var App *Application

func SaveConfig(config interface{}) {
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		logrus.Panic(err)
		return
	}

	var savePath string
	switch config.(type) {
	case AppConfig:
		savePath = App.AppConfigFile
		break
	default:
		logrus.Panic("unknown config type")
		break
	}

	err = ioutil.WriteFile(savePath, bytes, 0755)
	if err != nil {
		logrus.Panic(err)
		return
	}
}

func LoadConfig() {
	// Little fix for when the config file path contains a "~"
	usr, _ := user.Current()
	dir := usr.HomeDir
	if App.AppConfigFile == "~" {
		// In case of "~", which won't be caught by the "else if"
		App.AppConfigFile = dir
	} else if strings.HasPrefix(App.AppConfigFile, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		App.AppConfigFile = filepath.Join(dir, App.AppConfigFile[2:])
	}

	logrus.Debugf("Loading config from %s", App.AppConfigFile)
	appConfigContent, err := ioutil.ReadFile(App.AppConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			logrus.Infof("config file doesnt exist. Creating empty one")
			CreateDefaultAppConfig()
			return
		}

		logrus.Panic(err)
		return
	}

	var appConfig AppConfig
	err = json.Unmarshal(appConfigContent, &appConfig)
	if err != nil {
		logrus.Panic(err)
		return
	}

	if App.Flags.S3AccessKey != "" {
		appConfig.S3AccessKey = App.Flags.S3AccessKey
	}

	if App.Flags.S3SecretKey != "" {
		appConfig.S3SecretKey = App.Flags.S3SecretKey
	}

	if App.Flags.Project != "" {
		appConfig.Project = App.Flags.Project
	}

	App.AppConfig = appConfig
}

func CreateDefaultAppConfig() {
	SaveConfig(AppConfig{
		S3AccessKey: "",
		S3SecretKey: "",
		Project:     "",
	})
}

func init() {
	App = &Application{}
}
