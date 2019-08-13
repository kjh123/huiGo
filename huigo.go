package main

import (
	"fmt"
	"github.com/kjh123/huiGo/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	app *cli.App
	configFile string
)

func init() {
	app = cli.NewApp()
	app.Name = "app"
	app.Author = "hui"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "config",
			Value: "config.yml",
			Destination: &configFile,
		},
	}
	// Commands
	app.Commands = []cli.Command {
		{
			Name: "migrate",
			Usage: "run migrations",
			Action: func(c *cli.Context) error {
				return cmd.Migrate(configFile)
			},
		},
		{
			Name: "run",
			Usage: "run serve",
			Action: func(c *cli.Context) error {
				return cmd.Serve(configFile)
			},
		},
		{
			Name: "db",
			Usage: "database operations",
			Category: "database",
			Subcommands: []cli.Command {
				{
					Name: "insert",
					Usage: "insert data",
					Action: func(c *cli.Context) error {
						fmt.Println("insert subcommand")
						return nil
					},
				},
				{
					Name: "delete",
					Usage: "delete data",
					Action: func(c *cli.Context) error {
						fmt.Println("delete subcommand")
						return nil
					},
				},
			},
		},
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
