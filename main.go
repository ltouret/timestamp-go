package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func truncateText(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:max]
}

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

// ? add in .env
// DB_USER = root
// DB_PASS =
func SetupDb() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/")
	if err != nil {
		fmt.Println("Connection Db", err)
	}
	defer db.Close()

	// Check if the database exists
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS analytics")
	if err != nil {
		fmt.Println("Create Db", err)
	}

	// Switch to the newly created database
	_, err = db.Exec("USE analytics")
	if err != nil {
		fmt.Println("Use Db", err)
	}

	// Create the table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS test (
            id INT AUTO_INCREMENT PRIMARY KEY,
            userAgent VARCHAR(255) DEFAULT NULL,
            clientIP VARCHAR(255) DEFAULT NULL,
            timestamp VARCHAR(255) DEFAULT NULL,
            responseTime VARCHAR(255) DEFAULT NULL,
            queryParameters VARCHAR(255) DEFAULT NULL,
			userUuid CHAR(36) DEFAULT NULL
        )
    `)
	if err != nil {
		fmt.Println("Create Table", err)
	}

	fmt.Println("Database and table initialization completed.")
}

func main() {
	SetupDb() //? add here or in another main...
	router := gin.Default()
	router.Use(gin.Recovery()) // useful for 500 error do i keep this?
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1", v1MiddlewareTest())
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
