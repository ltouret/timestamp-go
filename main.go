package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	currentTime := time.Now()
	router.GET("/:timestamp", func(c *gin.Context) {
		timestamp := c.Param("timestamp")
		fmt.Printf("%T\n", timestamp)
		fmt.Println("The time is", timestamp)
		c.JSON(http.StatusOK, gin.H{
			"unix": currentTime,
			"utc":  currentTime,
		})
	})
}

func main() {
	router := gin.Default()
	SetupRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
