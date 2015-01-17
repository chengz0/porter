package main

import (
	"github.com/chengz0/porter/global"
	"github.com/chengz0/porter/goetcd"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	global.InitEnv("config.ini")
	//
	client := goetcd.NewEtcdClient("http://54.223.148.9:4001")

	//
	app := cli.NewApp()
	app.Name = "porter"
	app.Author = "caizi"
	app.Version = "1.0.1-dev"
	app.Email = "chengzhang@deepglint.com"
	app.Usage = "dev ops"
	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "list modules of product",
			Action: func(context *cli.Context) {
				goetcd.ListModules(client)
			},
		},
		{
			Name:  "init",
			Usage: "init file system on etcd",
			Action: func(context *cli.Context) {
				log.Println(context.Args())
				goetcd.InitDir4Module(client, context.Args().First())
			},
		},
		{
			Name:  "push",
			Usage: "porter push <module> <path>",
			Action: func(context *cli.Context) {
				goetcd.InitDirTree4Module(client, context.Args().First(), context.Args().Get(1))
			},
		},
		{
			Name:        "iter",
			Usage:       "porter iter <module> <version> <role>",
			Description: "iter version of module to etcd",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "module",
					Usage: "name of module to be iter",
				},
				cli.StringFlag{
					Name:  "version",
					Usage: "version of module to be iter",
				},
				cli.StringFlag{
					Name:  "role",
					Value: "dev",
					Usage: "role of developer",
				},
			},
			Action: func(context *cli.Context) {
				goetcd.AddIterVersion(client, context.String("module"), context.String("version"), context.String("role"))
			},
		},
		{
			Name:        "pass",
			Usage:       "porter pass <module> <from> <to>",
			Description: "pass version of module to next phase",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "module",
					Usage: "module to phase",
				},
				cli.StringFlag{
					Name:  "from",
					Usage: "from phase",
				},
				cli.StringFlag{
					Name:  "to",
					Usage: "to phase",
				},
			},
			Action: func(context *cli.Context) {
				goetcd.Pass2NextPhase(client, context.String("module"), context.String("from"), context.String("to"))
			},
		},
	}

	app.Run(os.Args)
}
