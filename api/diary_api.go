package api

import "github.com/gin-gonic/gin"

type DiaryApi interface {
	AddDiary() gin.HandlerFunc

	UpdateDiary() gin.HandlerFunc

	DelDiary() gin.HandlerFunc

	GetDiary() gin.HandlerFunc
}
