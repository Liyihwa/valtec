package main

import (
	_ "valtec/file-service/init"
	"valtec/file-service/router"
	"valtec/pkg/config"
)

func main() {
	router.Router().Run(config.GetStringSevere("gin", "port"))
}
