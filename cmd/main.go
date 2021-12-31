package main

import (
	"github.com/urfave/cli"
	"go-cli/cmd/command"
	"log"
	"os"
)

func main() {

	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "go-cli"
	//程序的用途描述
	oApp.Usage = "创造一切"
	//程序的版本号
	oApp.Version = "v1.0.0"

	oApp.Commands = append(oApp.Commands, command.Db2StructCommand,command.ApiCommand,command.WebCommand)

	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

