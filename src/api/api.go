package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Working!"})
		})
	}

	r.Run(":5005") // اینجا r.Run درسته، نه c.Run
}
