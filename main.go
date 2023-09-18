package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//? Can add just this and it will work to router.GET() -> use like this and call it endpoint
// func test(c *gin.Context) {
// 	currentTime := time.Now().UTC().Format(time.RFC822)
// 	timestamp := c.Param("timestamp")
// 	fmt.Println("The time is", currentTime, timestamp, c.ClientIP())
// 	c.JSON(http.StatusOK, gin.H{
// 		"unix": currentTime,
// 		"utc":  currentTime,
// 	})
// }

// ? move this to its own endpoint on top of this
func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/:timestamp", func(c *gin.Context) {
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
		c.JSON(http.StatusOK, gin.H{
			"error": "invalid Date",
		})
	})
}

func main() {
	router := gin.Default()
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1")
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
