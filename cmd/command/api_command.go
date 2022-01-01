package command

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"syscall"
)

var (
	kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)

	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

type Color struct {
	black        int // 黑色
	blue         int // 蓝色
	green        int // 绿色
	cyan         int // 青色
	red          int // 红色
	purple       int // 紫色
	yellow       int // 黄色
	light_gray   int // 淡灰色（系统默认值）
	gray         int // 灰色
	light_blue   int // 亮蓝色
	light_green  int // 亮绿色
	light_cyan   int // 亮青色
	light_red    int // 亮红色
	light_purple int // 亮紫色
	light_yellow int // 亮黄色
	white        int // 白色
}

// 输出有颜色的字体
func ColorPrint(s string, i int) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	println(s)
	CloseHandle.Call(handle)
}

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
		dir=path.Join(dir,fName)
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

		ColorPrint("",FontColor.white)
		ColorPrint("[show]:",FontColor.white)
		ColorPrint("",FontColor.white)
		showPro(dir,1)
		ColorPrint("",FontColor.white)
	   return nil

	},
}

func showPro(pathname string,lev int) error {
	rd, err := ioutil.ReadDir(pathname)

	for _, fi := range rd {
		if fi.IsDir() {
			ColorPrint(strings.Repeat(" ",lev)+fi.Name(),FontColor.light_yellow)
			showPro(path.Join(pathname,fi.Name()),lev+1)

		} else {

			//fmt.Println(fi.Name())
			var fName=""
			if lev==0{
				fName=fi.Name()
			}else{
				fName=strings.Repeat(" ",lev+1)+"-"+fi.Name()
			}
			ColorPrint(fName,FontColor.light_green)

		}
	}

	return err
}

