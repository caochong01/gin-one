package routers

import (
	"github.com/caochong01/gin-one/api"
	"github.com/gin-gonic/gin"
)

func diaryRoute(router *gin.Engine, api api.DiaryApi) {
	diaryGroup := router.Group("/diary")
	{
		diaryGroup.GET("/addDiary", api.AddDiary())
		diaryGroup.GET("/updateDiary", api.UpdateDiary())
		diaryGroup.GET("/delDiary", api.DelDiary())
	}
}
