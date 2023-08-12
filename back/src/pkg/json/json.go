package json

import (
	"encoding/json"
	"valtec/pkg/log"
)

func ToJson(value any) string {
	jsonStr, err := json.Marshal(value)
	if err != nil {
		log.Warn("{_r}ToJson{;} : {rx}%s{;}", err.Error())
		panic(err)
	}
	return string(jsonStr)
}
