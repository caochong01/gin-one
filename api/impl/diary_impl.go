package impl

import (
	"fmt"
	"github.com/caochong01/gin-one/config"
	"github.com/caochong01/gin-one/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DiaryImpl struct{}

func (n *DiaryImpl) GetDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		var diary entity.Diary
		if err := c.ShouldBindJSON(&diary); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("get")

		c.JSON(http.StatusOK, gin.H{"message": "get success"})
	}
}

func (n *DiaryImpl) AddDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		config.GetDB().Create(entity.Diary{
			DiaryTitle:  "",
			DiaryDetail: "",
		})
		fmt.Println("add")
	}
}

func (n *DiaryImpl) UpdateDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("update")
	}
}

func (n *DiaryImpl) DelDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("delete")
	}
}
