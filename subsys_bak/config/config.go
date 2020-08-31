package config

import (
	"time"
	"log"
	"github.com/go-ini/ini"
)

var (
	Version   = "0.0.0"
	Commit    = ""
	BuildTime = "2020-8-13"
)

type ServerConfig struct {
	Common *CommonConfig
	Bcos *BcosConfig
	Db *DbConfig
	Path *PathConfig
}

type CommonConfig struct {
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

//数据库配置
type DbConfig struct {
	Driver string
	Connstr string
}

//Bcos配置
type BcosConfig struct {
	Connstr string
	GroupId uint
	Keydir string
	Contractaddr string
}

//路径配置
type PathConfig struct {
	HomeSavePath string
	UserSavePath string
	PointSavePath string
	LogSavePath string
}


var CommonSetting = &CommonConfig{}
var DbSetting = &DbConfig{}
var BcosSetting = &BcosConfig{}
var PathSetting = &PathConfig{}

var Config = &ServerConfig{} //引用配置文件结构


var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatal("failed to run config setup, and parsed config file")
	}
	mapTo("common", CommonSetting)
	mapTo("db", DbSetting)
	mapTo("bcos", BcosSetting)
	mapTo("path", PathSetting)

	Config = &ServerConfig{
		Common : CommonSetting,
		Db : DbSetting,
		Bcos : BcosSetting,
		Path : PathSetting,
	}

	Config.Common.ReadTimeout = Config.Common.ReadTimeout * time.Second
	Config.Common.WriteTimeout = Config.Common.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
