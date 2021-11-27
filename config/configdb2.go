package config

import "github.com/tkanos/gonfig"

type Configurationdb2 struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfigdb2() Configurationdb2 {
	confdb2 := Configurationdb2{}
	gonfig.GetConf("config/configdb2.json", &confdb2)
	return confdb2
}
