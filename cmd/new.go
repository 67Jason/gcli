// new project
package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"strings"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "创建个很厉害的新项目",
	Long: `创建个很厉害不会倒闭的项目  
gcli new [appname] [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]  [-gopath=false] [-gcli=v1.12.1]
`,
	Run: func(cmd *cobra.Command, args []string) {
		initNew(cmd, args)
	},
}

var apiconf = `appname = {{.Appname}}
httpport = 8080
runmode = dev
autorender = false
copyrequestbody = true
EnableDocs = true
sqlconn = {{.SQLConnStr}}
`
var goVersion = 1.14
var apiMaingo = `package main
import (
	_ "{{.Appname}}/routers"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"{{.Appname}}/config"
	"{{.Appname}}/logger"
	"{{.Appname}}/orm"
)
func main() {
	// 获取运行的必要参数
	envName := flag.String("env", "", "set the program running env")
	flag.Parse()
	if *envName == "" {
		logrus.Panic("env params error")
	}
	// 读取配置文件
	config.Config.Start("config/" + *envName + ".yml")
	// 开启logger
	logger.Start()
	// 开启orm
	orm.Start()
	// gin api
	router := gin.New()
	routes.Start(router)
	logrus.Infof("asd")
	// 启动端口服务
	err := router.Run(config.Config.ServerAddr)
	if err != nil {
		logrus.Panicf("gin run err:%s", err.Error())
	}
}
`

var apiMainconngo = ``

var goMod = `
module %s

go %s

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/EDDYCJY/go-gin-example v0.0.0-20191007083155-a98c25f2172a
	github.com/boombuler/barcode v1.0.1-0.20180315051053-3c06908149f7
	github.com/go-ini/ini v1.51.1
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gomodule/redigo v2.0.1-0.20180401191855-9352ab68be13+incompatible
)
`

var apirouter = ``

var APIModels = ``

func init() {
	rootCmd.AddCommand(newCmd)
}

func initNew(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		newProject(cmd, args[0])
	} else {
		log.Println("no args")
	}
}

func newProject(cmd *cobra.Command, projectName string) error {
	err := pwd()
	if err != nil {
		return err
	}
	cmdStr := "mkdir apiresponse config constants logger orm routes"
	err = mkdir(cmdStr)
	if err != nil {
		return err
	}
	WriteToFile("main.go",strings.Replace(
		apiMaingo, "{{.Appname}}", projectName, -1),)
	WriteToFile("go.mod", fmt.Sprintf(goMod, projectName, goVersion))

	fmt.Println("successful")
	return nil
}

// 输出当前所在目录
func pwd() error {
	cmd := exec.Command("sh", "-c", "pwd")
	opBytes, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(opBytes))
	return nil
}

// 创建所有文件夹
func mkdir(cmdStr string) error {
	if cmdStr == "" {
		return nil
	}

	fmt.Println(cmdStr)
	cmd := exec.Command("sh", "-c", cmdStr)
	var stderr bytes.Buffer
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}
	return nil
}

func WriteToFile(filename, content string) error {
	f, err := os.Create(filename)
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}
