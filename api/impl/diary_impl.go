package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Diary struct{}

func (n *Diary) AddDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("add")
	}
}

func (n *Diary) UpdateDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("update")
	}
}

func (n *Diary) DelDiary() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("delete")
	}
}
