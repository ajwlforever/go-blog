package config

//
//import (
//	"github.com/BurntSushi/toml"
//	"time"
//)
//
//type Conf struct {
//	Title    string
//	App      app
//	DB       mysql `toml:"mysql"`
//	Redis    map[string]redis
//	Releases releases
//	Company  Company
//}
//
//type app struct {
//	Author  string
//	Org     string `toml:"organization"`
//	Mark    string
//	Release time.Time
//}
//
//type mysql struct {
//	Server  string
//	Ports   []int
//	ConnMax int `toml:"connection_max"`
//	Enabled bool
//}
//
//type redis struct {
//	Host string
//	Port int
//}
//
//type releases struct {
//	Release []string
//	Tags    [][]interface{}
//}
//
//type Company struct {
//	Name   string
//	Detail detail
//}
//
//type detail struct {
//	Type string
//	Addr string
//	ICP  string
//}
//
//var Connn = &Conf{}
//
//// 读取配置
//func init() {
//
//	path := "config/conf.toml"
//	_, err := toml.DecodeFile(path, Connn)
//	if err != nil {
//		panic(err)
//	}
//
//}
