package send

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/internetarchive/isodos-go/config"
	"github.com/internetarchive/isodos-go/pkg/isodos"
	"github.com/internetarchive/isodos-go/pkg/utils"
	"github.com/urfave/cli"
)

func NewSendListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Send a list of URLs to Isodos",
		Action: CmdSendList,
	}
}

func CmdSendList(c *cli.Context) {
	// Initialize Isodos config
	client := isodos.Init(config.App.AppConfig.S3AccessKey, config.App.AppConfig.S3SecretKey, config.App.AppConfig.Project)

	// Validate input
	if utils.FileExists(c.Args()[0]) == false {
		log.Fatalln(c.Args()[0] + " isn't a valid file.")
	}

	// Load seeds
	seeds, err := utils.LoadSeedsFromFile(c.Args()[0])
	if err != nil {
		log.Fatalln(c.Args()[0]+" isn't a valid file,", err.Error())
	}

	// Send seeds
	response, err := client.Send(seeds)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
