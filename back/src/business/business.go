package main

import (
	_ "valtec/business/init"
	"valtec/business/router"
	"valtec/pkg/config"
)

func main() {
	router.Router().Run(config.GetStringSevere("gin", "port"))
}
