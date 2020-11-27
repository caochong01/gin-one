package routers

import (
	"github.com/caochong01/gin-one/api/impl"
	"github.com/gin-gonic/gin"
)

func MainRouters(router *gin.Engine) {
	// 日记模块路由
	diaryRoute(router, &impl.DiaryImpl{})

	// TODO ...

}
