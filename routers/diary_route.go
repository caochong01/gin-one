package routers

import (
	"github.com/caochong01/gin-one/api"
	"github.com/gin-gonic/gin"
)

func diaryRoute(router *gin.Engine, api api.DiaryApi) {
	diaryGroup := router.Group("/diary")
	{
		diaryGroup.GET("/getDiary", api.GetDiary())
		diaryGroup.POST("/addDiary", api.AddDiary())
		diaryGroup.PUT("/updateDiary", api.UpdateDiary())
		diaryGroup.DELETE("/delDiary", api.DelDiary())
	}
}
