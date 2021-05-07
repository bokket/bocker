package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `bocker is a simple container runtime docker implementation.
			  The purpose of this project is to learn how docker works and how to 
write a docker by golang language`

func main()  {
	app := cli.NewApp()
	app.Name="mydocker--bocker"
	app.Usage=usage

	app.Commands = [] cli.Command {
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args) ; err != nil {
		log.Fatal(err)
	}

}