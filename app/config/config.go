package config

import (
	"fmt"
	"log"

	"gopkg.in/go-ini/ini.v1"
	_ "gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

// main関数の前にConfigを設定
func init() {
	LoadConfig()
}

func LoadConfig() {
	fmt.Println("test")
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DBName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
