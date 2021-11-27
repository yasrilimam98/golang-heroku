package config

import "github.com/tkanos/gonfig"

type Configurationdb1 struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfigdb1() Configurationdb1 {
	confdb1 := Configurationdb1{}
	gonfig.GetConf("config/configdb1.json", &confdb1)
	return confdb1
}
