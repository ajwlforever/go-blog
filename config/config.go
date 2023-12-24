package config

import "github.com/BurntSushi/toml"

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
}

var Conf = &Config{}

// 读取配置
func init() {

	path := "config/config.toml"
	_, err := toml.DecodeFile(path, Conf)
	if err != nil {
		panic(err)
	}

}
