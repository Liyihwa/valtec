package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"valtec/business/model"
	"valtec/pkg/database"
	"valtec/pkg/json"
	"valtec/pkg/log"
	"valtec/pkg/redis"
	"valtec/pkg/res"
)

func selectPositions(router *gin.Engine) {
	router.GET("/positions/:mid/:hid/:skill", func(c *gin.Context) {
		where := model.Position{}
		err := c.ShouldBindUri(&where)
		if err != nil {
			log.Warn("{_r}selectPositions bindUri{;} : {rx}%s{;}", err.Error())
			panic(err)
		}
		redisKey := fmt.Sprintf("positions:%d:%d:%s", where.MapID, where.HeroID, where.Skill)
		if redis.HandleJsonCache(c, redisKey) {
			return
		}

		var positions []model.Position
		database.Select(where, &positions, "id", "stand_x", "stand_y", "put_x", "put_y")
		positionsJson := json.ToJson(positions)

		redis.AddJsonCache(redisKey, positionsJson)
		c.JSON(200, res.Ok().Data(positionsJson))
	})
}

func selectPosition(router *gin.Engine) {
	router.GET("/position/:id", func(c *gin.Context) {
		where := model.Position{}
		err := c.ShouldBindUri(&where)
		if err != nil {
			log.Warn("{_r}selectPosition bindUri{;} : {rx}%s{;}", err.Error())
			panic(err)
		}
		redisKey := fmt.Sprintf("position:%d", where.ID)
		if redis.HandleJsonCache(c, redisKey) {
			return
		}

		position := model.Position{}
		dbRes := database.SelectOne(where, &position, "description", "like", "dislike")
		if dbRes.RowsAffected == 0 {
			log.Warn("{_r}selectPosition{;} {rx}Unknow id %d{;}", where.ID)
			c.JSON(200, res.Fail().Msg("No such id"))
		} else {
			positionJson := json.ToJson(position)
			redis.AddJsonCache(redisKey, positionJson)
			c.JSON(200, res.Ok().Data(positionJson))
		}
	})
}
