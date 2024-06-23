package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/stleox/dockee/cgroups/subsystems"
	"github.com/stleox/dockee/container"
	"github.com/stleox/dockee/network"
	"github.com/urfave/cli"
)

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container (Internal command, do not use directly)",
	Action: func(context *cli.Context) error {
		log.Infof("initing the container")
		cmd := context.Args().Get(0)
		log.Debugf("command %s", cmd)
		err := container.RunContainerInitProcess()
		return err
	},
}
var reInitCommand = cli.Command{
	Name:  "reinit",
	Usage: "Reinit container (Internal command, do not use directly)",
	Action: func(context *cli.Context) error {
		log.Infof("reinit come on")
		cmd := context.Args().Get(0)
		log.Debugf("command %s", cmd)
		err := container.RunContainerInitProcess()
		return err
	},
}

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpu",
			Usage: "cpu limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
		cli.StringFlag{
			Name:  "image",
			Usage: "the image name used to build the container",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
		cli.StringSliceFlag{
			Name:  "e",
			Usage: "set environment",
		},
		cli.StringFlag{
			Name:  "net",
			Usage: "container network",
		},
		cli.StringSliceFlag{
			Name:  "p",
			Usage: "port mapping",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		// enable tty
		tty := context.Bool("it")
		detach := context.Bool("d")
		if tty && detach {
			return fmt.Errorf("it and d paramter can not both provided")
		}
		// cgroup configure
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			Cpu:         context.String("cpu"),
			Cpuset:      context.String("cpuset"),
		}
		// volume
		volume := context.String("v")
		// image name
		imageName := context.String("image")
		// container name
		containerName := context.String("name")
		// env list
		envSlice := context.StringSlice("e")
		// netowrk
		network := context.String("net")
		// port
		portmapping := context.StringSlice("p")
		err := Run(tty, cmdArray, resConf, volume, imageName, containerName,
			envSlice, network, portmapping)
		if err != nil {
			return fmt.Errorf("run error %v", err)
		}
		return nil
	},
}

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "commit a container into images",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container name")
		}
		containerName := context.Args().Get(0)
		imageName := context.Args().Get(1)
		commitContainer(containerName, imageName)
		return nil
	},
}

var listCommand = cli.Command{
	Name:  "ps",
	Usage: "list all the containers",
	Action: func(context *cli.Context) error {
		ListContainers()
		return nil
	},
}

var logCommand = cli.Command{
	Name:  "logs",
	Usage: "print logs of a container",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Please input your container name")
		}
		containerName := context.Args().Get(0)
		logContainer(containerName)
		return nil
	},
}

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "exec a command into container",
	Action: func(context *cli.Context) error {
		//This is for callback
		if os.Getenv(ENV_EXEC_PID) != "" {
			log.Debugf("pid callback pid %d", os.Getgid())
			return nil
		}

		if len(context.Args()) < 2 {
			return fmt.Errorf("Missing container name or command")
		}
		containerName := context.Args().Get(0)
		var commandArray []string
		for _, arg := range context.Args().Tail() {
			commandArray = append(commandArray, arg)
		}
		ExecContainer(containerName, commandArray)
		return nil
	},
}
var startCommand = cli.Command{
	Name:  "start",
	Usage: "start a container",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container name")
		}
		containerName := context.Args().Get(0)
		startContainer(containerName)
		return nil
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "stop a container",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container name")
		}
		containerName := context.Args().Get(0)
		stopContainer(containerName)
		return nil
	},
}

var removeCommand = cli.Command{
	Name:  "rm",
	Usage: "remove unused containers",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container name")
		}
		containerName := context.Args().Get(0)
		removeContainer(containerName)
		return nil
	},
}

var networkCommand = cli.Command{
	Name:  "network",
	Usage: "container network commands",
	Subcommands: []cli.Command{
		{
			Name:  "create",
			Usage: "create a container network",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "driver",
					Usage: "network driver",
				},
				cli.StringFlag{
					Name:  "subnet",
					Usage: "subnet cidr",
				},
			},
			Action: func(context *cli.Context) error {
				if len(context.Args()) < 1 {
					return fmt.Errorf("Missing network name")
				}
				err := network.Init()
				if err != nil {
					return fmt.Errorf("init network error: %+v", err)
				}
				err = network.CreateNetwork(context.String("driver"),
					context.String("subnet"), context.Args()[0])
				if err != nil {
					return fmt.Errorf("create network error: %+v", err)
				}
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list container network",
			Action: func(context *cli.Context) error {
				err := network.Init()
				if err != nil {
					return fmt.Errorf("init network error: %+v", err)
				}
				network.ListNetwork()
				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove container network",
			Action: func(context *cli.Context) error {
				if len(context.Args()) < 1 {
					return fmt.Errorf("Missing network name")
				}
				err := network.Init()
				if err != nil {
					return fmt.Errorf("init network error: %+v", err)
				}
				err = network.DeleteNetwork(context.Args()[0])
				if err != nil {
					return fmt.Errorf("remove network error: %+v", err)
				}
				return nil
			},
		},
	},
}
