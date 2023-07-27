package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"strconv"
)

var conf *ini.File

const filename = "config.ini"

func init() {
	var err error
	conf, err = ini.Load(filename)
	if err != nil {
		panic("config init error :" + err.Error())
	}
}

func GetConfigSevere(section string, key string) string {
	if !conf.HasSection(section) {
		panic(fmt.Errorf("No section named %s ", section))
	}
	sec := conf.Section(section)
	if !sec.HasKey(key) {
		panic(fmt.Errorf("No key named %s in section %s ", key, section))
	}
	return sec.Key(key).String()
}
func GetIntConfigSevere(section string, key string) int {
	str := GetConfigSevere(section, key)
	intVal, err := strconv.Atoi(str)
	if err != nil {
		panic("config get error :" + err.Error())
	}
	return intVal
}
