package command

import (
	"github.com/urfave/cli"
	"go-cli/pkg/database"
	"io/ioutil"
	"os"
	"path"
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
			//columns := db.GetColumns(tb.TableName)
			ioutil.WriteFile(path.Join(c.String("path"), tb.TableName+".go"), []byte(""), os.ModePerm)
		}
		return nil
	},
}

func mkDir(dir string) {

	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, os.ModePerm)
		return
	}

}
