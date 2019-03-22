package main

import (
	"flag"
	"fmt"
	"runtime"
)

var (
	confFile string //配置文件路径
)

//解析命令行参数
func initArgs() {
	flag.StringVar(&confFile, "config", "./master.json", "")

}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //多和数量和电脑核心数量相同，最大限度发挥多核优势
}

func main() {
	var (
		err error
	)
	initArgs()
	//初始化线程
	initEnv()

	//加载配置
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}

	//启动api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

ERR:
	fmt.Println("err")

}
