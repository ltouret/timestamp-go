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
func v1MiddlewareTest(analyticsDb *AnalyticsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()

		timestamp := truncateText(t.Format(http.TimeFormat), 255)
		// query := truncateText(c.Param("timestamp"), 255) //! sanitize this before saving it in db
		// userAgent := truncateText(c.Request.UserAgent(), 255)
		// clientIP := c.ClientIP()
		// referer := truncateText(c.Request.Referer(), 255)
		// latency := time.Since(t)

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

type AnalyticsService struct {
	db *sql.DB
}

// NewUserService 'constructs' a UserService that is ready to use.
// It requires an initialized sql.DB instance.
func NewAnalyticsService(db *sql.DB) (*AnalyticsService, error) {
	return &AnalyticsService{db}, nil
}

// ? add in .env
// DB_USER = root
// DB_PASS =
// ? dependency injection this or not?
// ! setup only for timestamp-go, will need more work for others
func SetupDb(db *sql.DB) {
	// Check if the database exists
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS analytics")
	if err != nil {
		fmt.Println("Create Db", err)
	}

	// Switch to the newly created database
	_, err = db.Exec("USE analytics")
	if err != nil {
		fmt.Println("Use Db", err)
	}

	// Create the table if it doesn't exist
	// ? if i use this same porject to do the other api for FCC then i will to create more tables, and each table will have the analytics of a microservice
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
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/")
	if err != nil {
		//! quit code here
		fmt.Println("Connection Db", err)
	}
	defer db.Close()
	SetupDb(db) //? add here or in another main... -> if in another main no dep injection(?)
	analyticsDb, err := NewAnalyticsService(db)
	if err != nil {
		//! quit code here
		fmt.Println("analyticsDb", err)
	}
	router := gin.Default()
	router.Use(gin.Recovery()) // useful for 500 error do i keep this?
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1", v1MiddlewareTest(analyticsDb))
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
