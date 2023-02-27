package modules

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	fmt.Println("I am healthy!")

	c.JSON(200, gin.H{"status": 200, "message": "I am healthy!"})
}
