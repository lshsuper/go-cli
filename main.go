package main

import (
	"fmt"
	"github.com/lshsuper/go-cli/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {

	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "github.com/lshsuper/go-cli"
	//程序的用途描述
	oApp.Usage = "创造一切"

	//程序的版本号
	oApp.Version =`v1.0.0 
                                .__  .__ 
   ____   ____             ____ |  | |__|
  / ___\ /  _ \   ______ _/ ___\|  | |  |
 / /_/  >  <_> ) /_____/ \  \___|  |_|  |
 \___  / \____/           \___  >____/__|
/_____/                       \/         

`

	oApp.Commands = append(oApp.Commands, cmd.Db2StructCommand, cmd.ApiCommand, cmd.WebCommand)

	//启动
	if err := oApp.Run(os.Args); err != nil {
		fmt.Println("ok")
		log.Fatal(err)
	}

}

