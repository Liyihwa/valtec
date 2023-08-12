package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"valtec/business/model"
	"valtec/pkg/database"
	"valtec/pkg/email"
	"valtec/pkg/jwt"
	"valtec/pkg/log"
	"valtec/pkg/rand"
	"valtec/pkg/redis"
	"valtec/pkg/res"
	"valtec/pkg/utils"
)

func emailVerif(router *gin.Engine) {
	router.GET("/emailVerif/:email", func(c *gin.Context) {
		userEmail := c.Param("email")
		if userEmail == "" {
			log.Warn("[{_r}emailVerif{;}]: {rx}%s{;}", "email is empty") //出现此情况说明用户可能绕过了前端
			c.JSON(200, res.Fail().Msg("邮箱不能为空!"))
			return
		}

		emcode := redis.Get("emcode:" + userEmail)
		if emcode != "" {
			log.Warn("[{_r}emailVerif{;}] email:{ux}%s{;},err:{rx}%s{;}", userEmail, "emcode has setted") //出现此情况说明用户可能绕过了前端
			c.JSON(200, res.Fail().Msg("验证码已发送"))
			return
		}

		result := database.SelectOne(&model.User{Email: userEmail}, &model.User{}, "id")
		if result.RowsAffected != 0 {
			c.JSON(200, res.Fail().Msg("邮箱已存在,请直接登录!"))
			return
		}

		code := rand.RandString(6, rand.Nums)
		err := email.Send(userEmail, "注册valtec.top的邮箱验证码", "<p>您的验证码为:</p><h2>"+code+"</h2><p>有效期60秒</p>")
		if err != nil {
			log.Erro("[{_r}email.Send{;}] :{rx}%s{;}", err.Error()) //服务器内部出错
			c.JSON(200, res.Fail().Msg("验证码获取失败,请稍后重试"))
		}
		redis.Set("emcode:"+userEmail, code, 60*time.Second)

		c.JSON(200, res.Ok())
	})
}

func register(router *gin.Engine) {
	router.POST("/register", func(c *gin.Context) {
		user := model.User{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Warn("[{_r}register{;}] : {rx}%s{;}", err.Error())
			panic(err)
		}
		code := redis.Get("emcode:" + user.Email)
		if user.Captcha != code {
			c.JSON(200, res.Fail().Msg("验证码错误!"))
			return
		}

		result := database.SelectOne(&model.User{Email: user.Email}, &model.User{}, "id")
		if result.RowsAffected != 0 {
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
			log.Warn("[{_r}login{;}] : {rx}%s{;}", err.Error())
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
