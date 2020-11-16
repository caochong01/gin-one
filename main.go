package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "add",
		})
	})
	r.GET("/update", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "update",
		})
	})
	r.GET("/delete", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"message": name,
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
