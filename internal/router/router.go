package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/ping/:name", Pong)
	}
	r.GET("/ping", Pong)
	return r
}

func Pong(c *gin.Context) {
	// name := c.Param("name")
	name := c.DefaultQuery("name", "Goldz")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
	})
}
