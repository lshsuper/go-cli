package command

import (
	"fmt"
	"github.com/urfave/cli"
	"go-cli/pkg/database"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var Db2StructCommand = cli.Command{
	Name:  "db2Struct",
	Usage: "数据库映射模型",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Usage: "用户库host",
		},
		cli.IntFlag{
			Name:  "port",
			Usage: "数据库端口",
		},
		cli.StringFlag{
			Name:  "uid",
			Usage: "用户名",
		},
		cli.StringFlag{
			Name:  "pwd",
			Usage: "密码",
		},
		cli.StringFlag{
			Name:  "db",
			Usage: "数据库名称",
		},
		cli.StringFlag{
			Name:  "path",
			Usage: "存放路径",
		},
		cli.IntFlag{
			Name:  "dtype",
			Usage: "数据库类型(1->mysql|2-sqlserver)",
		},
	},
	Action: func(c *cli.Context) error {

		db := database.Register(database.DbType(c.Int("dtype")), database.DbConfig{
			Port:      c.Int("port"),
			UserName:  c.String("uid"),
			Password:  c.String("pwd"),
			DefaultDb: c.String("db"),
			Host:      c.String("host"),
		})

		tbs := db.GetTables(c.String("db"))
		for _, tb := range tbs {

			columns := db.GetColumns(tb.TableName)
			fileName:=fmt.Sprintf("%s.go", strings.ToLower(tb.TableName))
			structName:=strings.ReplaceAll(tb.TableName,"_","")
			structName=strings.Title(structName)
			structName=fmt.Sprintf("%sModel",structName)
			structContent:=new(strings.Builder)
			structContent.WriteString(fmt.Sprintf("//%s\r\n",tb.TableComment))
			structContent.WriteString(fmt.Sprintf("type %s struct{\r\n",structName))
			for _,c:=range columns{
				cName:=strings.Title(strings.ReplaceAll(c.Field,"_",""))
				structContent.WriteString(fmt.Sprintf("    %s %s   `gorm:\"column:%s\"` //%s \r\n ",cName,c.GetGoType(),c.Field,c.Comment))
			}
			structContent.WriteString("}")
			ioutil.WriteFile(path.Join(c.String("path"),fileName), []byte(structContent.String()), os.ModePerm)

			fmt.Println(tb.TableName," build success...")

		}
		fmt.Println("all table build success")
		return nil
	},
}

func mkDir(dir string) {

	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModePerm)
		return
	}

}
