package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "dockee"

	app.Commands = []cli.Command{
		initCommand,
		reInitCommand,
		runCommand,
		commitCommand,
		listCommand,
		logCommand,
		execCommand,
		startCommand,
		stopCommand,
		removeCommand,
		networkCommand,
	}

	// debug flag for debug mode
	debugFlag := &cli.BoolFlag{
		Name:  "debug",
		Usage: "enable debug mode",
	}
	app.Flags = []cli.Flag{
		debugFlag,
	}

	app.Before = func(context *cli.Context) error {
		// log.SetFormatter(&log.JSONFormatter{})

		// todo 日志改到文件，标准输出有点奇怪
		log.SetOutput(os.Stdout)

		// 设置日志模式
		if context.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
