package aliyun

import (
	"fmt"
	"mime/multipart"
	"valtec/pkg/config"
	"valtec/pkg/log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var endPoint, accessKeyID, accessKeySecret, bucketName string
var ossClient *oss.Client
var bucket *oss.Bucket

func init() {
	endPoint = config.GetStringSevere("aliyun", "endPoint")
	accessKeyID = config.GetStringSevere("aliyun", "accessKeyID")
	accessKeySecret = config.GetStringSevere("aliyun", "accessKeySecret")
	bucketName = config.GetStringSevere("aliyun", "bucketName")
	ossClient, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Erro("{_r}init aliyun error{;} : {rx}%s{;}", err.Error())
		panic(err)
	}
	bucket, err = ossClient.Bucket(bucketName)
	if err != nil {
		log.Erro("{_r}init aliyun error{;} : ", err.Error())
		panic(err)
	}
}

func Upload(file multipart.File, filename string) (string, error) {
	path := fmt.Sprintf("https://%s.%s/%s", bucketName, endPoint, filename)
	log.Info("path to upload {g}%s{;}", path)
	return path, bucket.PutObject(filename, file)
}

func Delete(filename string) error {
	path := fmt.Sprintf("https://%s.%s/%s", bucketName, endPoint, filename)
	log.Info("path to delete {g}%s{;}", path)
	return bucket.DeleteObject(filename)
}
