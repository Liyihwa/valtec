package router

import (
	"github.com/gin-gonic/gin"
	"strings"
	"valtec/pkg/aliyun"
	"valtec/pkg/log"
)

func upload(router *gin.Engine) {
	router.POST("/upload", func(c *gin.Context) {
		f, _ := c.FormFile("file")
		if f == nil {
			c.JSON(200, gin.H{"errno": 1, "message": "Not upload any file"})
		}
		log.Info("File uploaded {g}%s{;} size :%d ", f.Filename, f.Size)
		file, err := f.Open()
		if err != nil {
			log.Warn("{_r}File open errpr{;} :%s", err.Error())
			c.JSON(200, gin.H{"errno": 1, "message": "上传失败,请稍后再试.."})
			return
		}
		path, err := aliyun.Upload(file, f.Filename)
		if err != nil {
			log.Warn("{_y}oss upload failed{;} :%s", err.Error())
			c.JSON(200, gin.H{"errno": 1, "message": "上传失败,请稍后再试.."})
			return
		}

		c.JSON(200, gin.H{"errno": 0, "data": gin.H{
			"url": path,
		}})
	})
}

func deleteFile(router *gin.Engine) {
	router.DELETE("/delete", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Warn("{_r}ShouldBindJSON in deleteFile error{;} :%s", err.Error())
			c.JSON(200, gin.H{"errno": 1, "message": "删除失败,请稍后重试!"})
			return
		}
		url := data["url"].(string)
		tokens := strings.Split(url, "/")
		err := aliyun.Delete(tokens[len(tokens)-1])
		if err != nil {
			log.Warn("{_r}oss delete error{;} :%s", err.Error())
			c.JSON(200, gin.H{"errno": 1, "message": "删除失败,请稍后重试!"})
			return
		}
		c.JSON(200, gin.H{"errno": 0, "message": "删除成功!"})
	})
}
