package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)


var ApiCommand =cli.Command{
	Name:  "api",
	Usage: "构建一个api项目",
	Flags: []cli.Flag{
		cli.StringFlag{
		   Name: "name",
		   Usage: "名称",

		},
	},
	Action: func(c *cli.Context) error {


		fName:=c.String("name")

		dir, _:= filepath.Abs(filepath.Dir(os.Args[0]))
		dir =path.Join(dir, fName)
		os.Mkdir(dir,os.ModePerm)

		//service
		servicePath:=path.Join(dir,"service")
		os.Mkdir(servicePath,os.ModePerm)
		os.Mkdir(path.Join(servicePath,"default"),os.ModePerm)
		ioutil.WriteFile(path.Join(servicePath,"default","service.go"),[]byte(
			`package default

`),os.ModePerm)
		//controller
		controllerPath:=path.Join(dir,"controller")
		os.Mkdir(controllerPath,os.ModePerm)
		os.Mkdir(path.Join(controllerPath,"default"),os.ModePerm)
		ioutil.WriteFile(path.Join(controllerPath,"default","controller.go"),[]byte(
			`package default
type DefaultController struct{

}
func NewDefaultController()*DefaultController{
     return &DefaultController{}
}

func Test(c *gin.Context){

    c.JSON(http.StatusOK,gin.H{"data":"ok"})

}

`),os.ModePerm)

		//models
		modelsPath:=path.Join(dir,"models")
		os.Mkdir(modelsPath,os.ModePerm)
		os.Mkdir(path.Join(modelsPath,"default"),os.ModePerm)

		ioutil.WriteFile(path.Join(modelsPath,"default","model.go"),[]byte(
			`package  default
               
type DefaultModel struct {
     Id int `+"`gorm:\"columns:id\"`"+`
}
func(tb *DefaultModel)TableName()string{
    return "default"
}
        `),os.ModePerm)
		ioutil.WriteFile(path.Join(modelsPath,"default","dto.go"),[]byte(
			`package default

type DefaultDTO struct{
    Id int `+"`json:\"id\"`"+`
}
`),os.ModePerm)

		ioutil.WriteFile(path.Join(modelsPath,"default","enum.go"),[]byte(
			`package default

type DefaultType int

const(
   _DefaultType=iota
   DefaultA
)

func(e DefaultType)GetDescription()string{
         
     switch e{
       
        case DefaultA:
         return "A"
        default:
         return ""
      }

}

`),os.ModePerm)

		//dao
		daoPath:=path.Join(dir,"dao")
		os.Mkdir(daoPath,os.ModePerm)
		os.Mkdir(path.Join(daoPath,"default"),os.ModePerm)
		ioutil.WriteFile(path.Join(daoPath,"default","dao.go"),[]byte(
			`package default

`),os.ModePerm)
		//controller
		os.Mkdir(path.Join(dir,"controller"),os.ModePerm)



		//main

		ioutil.WriteFile(path.Join(dir,"main.go"),[]byte(
			`package main

func main(){


}

`),os.ModePerm)


		fmt.Println("[show]:")
		fmt.Println("")
		showPro(dir,1)
		fmt.Println("")
	   return nil

	},
}

func showPro(pathname string,lev int) error {
	rd, err := ioutil.ReadDir(pathname)

	for _, fi := range rd {
		if fi.IsDir() {
			fmt.Println(strings.Repeat(" ",lev)+fi.Name())
			showPro(path.Join(pathname,fi.Name()),lev+1)

		} else {

			//fmt.Println(fi.Name())
			var fName=""
			if lev==1{
				fName=fi.Name()
			}else{
				fName=strings.Repeat(" ",lev+1)+"-"+fi.Name()
			}
			fmt.Println(fName)

		}
	}

	return err
}

