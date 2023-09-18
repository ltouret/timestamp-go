package main

import (
	"fmt"
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

func SetupRoutes(router *gin.RouterGroup) {
	router.GET("/:timestamp", func(c *gin.Context) {
		timestamp := c.Param("timestamp")
		timestampUnix, err := strconv.ParseInt(timestamp, 10, 64)
		// var returnDate time.Time
		if err != nil {
			fmt.Println("Not Unix")
		}
		timestampUnix /= 1000
		dtUnix := time.Unix(timestampUnix, 0)
		dt, err := time.Parse("2020-12-25", timestamp)
		if err != nil {
			fmt.Println("Not date")
		}
		// date :=
		fmt.Println(timestampUnix, dt, dtUnix)
		currentTime := time.Now().UTC().Format(time.RFC822)
		fmt.Println("The time is", currentTime, timestamp, c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"unix": currentTime,
			"utc":  currentTime,
		})
	})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
