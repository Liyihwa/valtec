package utils

import (
	"crypto/sha256"
)

func SHA256Encrypt(datas ...string) string {
	hash := sha256.New()
	for _, data := range datas {
		hash.Write([]byte(data))
	}
	// 计算哈希值3
	return string(hash.Sum(nil))
}
