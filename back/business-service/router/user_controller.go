package router

import (
	"github.com/Liyihwa/logger"
	"github.com/gin-gonic/gin"
	"valtec/database"
	"valtec/model"
	"valtec/pkg/jwt"
	"valtec/pkg/res"
	"valtec/pkg/utils"
)

func register(router *gin.Engine) {
	router.POST("/register", func(c *gin.Context) {
		user := model.User{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			logger.Warn("{_r}register{;} 失败,err : {r}%s{;}", err.Error())
			panic("register err: " + err.Error())
		}
		result := database.SelectOne(&model.User{Email: user.Email}, &model.User{}, "id")
		if result.RowsAffected != 0 {
			logger.Debug("{_r}register{;} 失败,邮箱已存在: {u}%s{;}", user.Email)
			c.JSON(200, res.Fail().Msg("邮箱已存在!"))
			return
		}

		user.Password = utils.SHA256Encrypt(user.Password)
		database.AddOne(&user)
		tok := jwt.New().Set("uid", user.ID).Token()
		c.JSON(200, res.Ok().Data(tok))
	})
}

func login(router *gin.Engine) {
	router.POST("/login", func(c *gin.Context) {
		user := model.User{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			logger.Warn("{_r}login{;} 失败,err : {r}%s{;}", err.Error())
			panic("login err: " + err.Error())
		}
		user.Password = utils.SHA256Encrypt(user.Password)
		resUser := model.User{}
		result := database.SelectOne(&user, &resUser, "id")
		if result.RowsAffected == 0 {
			c.JSON(200, res.Fail().Msg("邮箱或密码错误,登录失败!"))
		} else {
			tok := jwt.New().Set("uid", resUser.ID).Token()
			c.JSON(200, res.Ok().Data(tok))
		}
	})
}
