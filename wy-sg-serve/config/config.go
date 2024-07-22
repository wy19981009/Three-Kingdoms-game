package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

const configFile = "/conf/conf.ini"

var File *goconfig.ConfigFile

// 加载此文件的时候 会先走此方法
func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	confPath := dir + configFile
	// 文件系统的读取
	File, err := goconfig.LoadConfigFile(confPath)
	if err != nil {
		log.Fatal("读取配置文件出错: ", err)
	}
	fmt.Println(File)
}

func A() {
	fmt.Println("A")
}
