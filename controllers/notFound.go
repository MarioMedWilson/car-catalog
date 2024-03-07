package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	fmt.Println("Endpoint Not Found")
	c.JSON(404, gin.H{
		"message": "Endpoint Not Found",
	})
}

func Test(c *gin.Context) {
	fmt.Println("Test")
	c.JSON(200, gin.H{
		"message": "Test",
	})
}

