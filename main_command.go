package main

import (
	"bocker/container"
	"fmt"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"

)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			mydocker run -ti [command]`,
			Flags:[]cli.Flag{
				cli.BoolFlag{
					Name: "ti",
					Usage: "enable tty",
				},
			},
			Action: func(context* cli.Context) error{
				if len(context.Args()) < 1 {
					return fmt.Errorf("Missing container command")
				}
				cmd := context.Args().Get(0)
				tty := context.Bool("ti")
				Run(tty,cmd)
				return nil
			},
}

var initCommand = cli.Command{
	Name: "init",
	Usage: "Init container process run user's process in container.Do not call it outside",
	Action: func(context *cli.Context) error{
		log.Info("init come on")
		cmd := context.Args().Get(0)
		log.Info("command %s",cmd)
		err := container.RunContainerInitProcess(cmd,nil)
		return err
	},
}