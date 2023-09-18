package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ? add middleware that will log analytics
// use dependecy injection to add db here?
func v1MiddlewareTest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		timestamp := c.Param("timestamp") //! sanitize this before saving it in db

		t := time.Now()

		c.Next()

		// after request
		latency := time.Since(t)
		fmt.Println(c.Request.UserAgent(), c.ClientIP(), c.Request.Referer(), t.Format(http.TimeFormat), latency, timestamp) //! erase later, add this to mdw
	}
}

// ? clean this
func timestampEndpoint(c *gin.Context) {
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
	router.GET("/:timestamp", timestampEndpoint)
}

func main() {
	router := gin.Default()
	router.Use(gin.Recovery()) // useful?
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1", v1MiddlewareTest())
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
