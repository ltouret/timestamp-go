package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//? add middleware that will log analytics

// ? clean this
func timestampEndpoint(c *gin.Context) {
	fmt.Println(c.Request.UserAgent(), c.ClientIP())
	timestamp := c.Param("timestamp")
	normalTimestamp, err := time.Parse("2006-01-02", timestamp)
	if err == nil {
		unixTime := normalTimestamp.UTC().Unix() * 1000
		c.JSON(http.StatusOK, gin.H{
			"unix": unixTime,
			"utc":  normalTimestamp.UTC().Format(http.TimeFormat),
		})
		return
	}
	timestampUnix, err := strconv.ParseInt(timestamp, 10, 64)
	if err == nil {
		timestampUnix /= 1000
		dtUnix := time.Unix(timestampUnix, 0)
		c.JSON(http.StatusOK, gin.H{
			"unix": timestampUnix * 1000,
			"utc":  dtUnix.UTC().Format(http.TimeFormat),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "invalid Date",
	})
}

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/:timestamp", timestampEndpoint)
}

func main() {
	router := gin.Default()
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1")
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
