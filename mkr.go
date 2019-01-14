package main

import (
	"fmt"
	"os"

	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mkr/logger"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mkr"
	app.Version = fmt.Sprintf("%s (rev:%s)", version, gitcommit)
	app.Usage = "A CLI tool for mackerel.io"
	app.Author = "Hatena Co., Ltd."
	app.Commands = Commands
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf",
			Value: config.DefaultConfig.Conffile,
			Usage: "Config file path",
		},
		cli.StringFlag{
			Name: "apibase",
			// this default value is set in config.LoadApibaseFromConfigWithFallback
			Usage: fmt.Sprintf("API Base (default: \"%s\")", config.DefaultConfig.Apibase),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Log("error", err.Error())
		os.Exit(1)
	}
}
