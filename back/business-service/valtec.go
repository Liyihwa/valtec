package main

import (
	"valtec/pkg/config"
	"valtec/router"
)

func main() {

	router.Router().Run(config.GetConfigSevere("gin", "port"))
}
