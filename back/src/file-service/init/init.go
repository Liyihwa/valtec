package init

import "valtec/pkg/config"

func init() {
	config.Init("file-service/config.ini")
}
