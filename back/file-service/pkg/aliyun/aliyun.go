package aliyun

import (
	"fmt"
	"github.com/Liyihwa/logger"
	"mime/multipart"
	"valtec-fileservice/pkg/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var endPoint, accessKeyID, accessKeySecret, bucketName string
var ossClient *oss.Client
var bucket *oss.Bucket

func init() {
	endPoint = config.GetConfigSevere("aliyun", "endPoint")
	accessKeyID = config.GetConfigSevere("aliyun", "accessKeyID")
	accessKeySecret = config.GetConfigSevere("aliyun", "accessKeySecret")
	bucketName = config.GetConfigSevere("aliyun", "bucketName")
	ossClient, err := oss.New(endPoint, accessKeyID, accessKeySecret)
	if err != nil {
		logger.Erro("{_r}init aliyun error{;} : ", err.Error())
		panic(err)
	}
	bucket, err = ossClient.Bucket(bucketName)
	if err != nil {
		logger.Erro("{_r}init aliyun error{;} : ", err.Error())
		panic(err)
	}
}

func Upload(file multipart.File, filename string) (string, error) {
	path := fmt.Sprintf("https://%s.%s/%s", bucketName, endPoint, filename)
	logger.Info("path to upload {g}%s{;}", path)
	return path, bucket.PutObject(filename, file)
}

func Delete(filename string) error {
	path := fmt.Sprintf("https://%s.%s/%s", bucketName, endPoint, filename)
	logger.Info("path to delete {g}%s{;}", path)
	return bucket.DeleteObject(filename)
}
