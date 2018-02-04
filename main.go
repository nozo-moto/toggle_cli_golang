package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func load_env_file() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	return nil
}

func run() error {
	app := cli.NewApp()

	app.Name = "Toggle CLI"
	app.Usage = "This app is CLI Tool of Toggle"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start toggle",
			Action: func(c *cli.Context) error {
				err := start()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop toggle",
			Action: func(c *cli.Context) error {
				err := stop()
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
	return nil
}

var API_TOKEN = ""

func main() {
	err := load_env_file()
	if err != nil {
		log.Fatal(err)
		panic("paniced")
	} else {
		API_TOKEN = os.Getenv("APITOKEN")
	}

	err = run()
	if err != nil {
		log.Fatal(err)
		panic("paniced")
	}
}
