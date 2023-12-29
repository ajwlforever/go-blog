package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	App app   `toml:"app"`
	Db  mysql `toml:"mysql"`
}

type app struct {
	Author string `toml:"author"`
}
type mysql struct {
	Server   string `toml:"server"`
	Ports    []int  `toml:"ports"`
	ConnMax  int    `toml:"connection_max"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Dbname   string `toml:"dbname"`
}

var Conf = &Config{}

// 读取配置
func init() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting")
	}
	path += "/config/config.toml"
	fmt.Printf("配置文件地址是:%s\n", path)
	_, err = toml.DecodeFile(path, Conf)
	if err != nil {
		panic(err)
	}

}
