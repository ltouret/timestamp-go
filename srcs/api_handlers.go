// api_handlers.go

package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func timestampEndpoint(c *gin.Context) {
	timestamp := c.Param("date")
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
		dtUnix := time.UnixMilli(timestampUnix)
		c.JSON(http.StatusOK, gin.H{
			"unix": timestampUnix,
			"utc":  dtUnix.UTC().Format(http.TimeFormat),
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "invalid Date",
	})
}

func noTimestampEndpoint(c *gin.Context) {
	timestamp := time.Now().UTC()
	c.JSON(http.StatusOK, gin.H{
		"unix": timestamp.UnixMilli(),
		"utc":  timestamp.Format(http.TimeFormat),
	})
}

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/", noTimestampEndpoint)
	router.GET("/:date", timestampEndpoint)
}
