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

func whoAmIEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ipaddress": c.ClientIP(),
		"language":  c.GetHeader("Accept-Language"),
		"software":  c.Request.UserAgent(),
	})
}

func SetupRoutes(router *gin.RouterGroup, analyticsDb *AnalyticsService) {
	// Timestamp Routes
	{
		timestamp := router.Group("/timestamp", middlewareTimestampAnalytics(analyticsDb))
		timestamp.GET("/", noTimestampEndpoint)
		timestamp.GET("/:date", timestampEndpoint)
	}
	// Request Header Routes
	{
		headerParser := router.Group("/header-parser") // add mdw for analytics?
		headerParser.GET("/whoami", whoAmIEndpoint, middlewareheaderParserAnalytics(analyticsDb))
	}
}
