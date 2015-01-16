package main

import (
	"github.com/chengz0/porter/goetcd"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
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
			Usage: "push files of module to etcd",
			Action: func(context *cli.Context) {
				goetcd.InitDirTree(client, context.Args().First(), context.Args().Get(1))
			},
		},
	}

	app.Run(os.Args)
}
