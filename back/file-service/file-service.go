package main

import (
	"valtec-fileservice/pkg/config"
	"valtec-fileservice/router"
)

func main() {
	router.Router().Run(config.GetConfigSevere("gin", "port"))
}
