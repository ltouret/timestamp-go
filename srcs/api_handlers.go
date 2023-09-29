// api_handlers.go

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// ! if the original url already exists then dont write to the db, return the uuid that exists already
func GETUrlShortenerEndpoint(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "http://www.google.com")
}

// ! use better validator
type ShortUrlJsonBody struct {
	Url string `json:"url"`
}

func POSTUrlShortenerEndpoint(c *gin.Context) {
	// var dummyDb map[string]string = make(map[string]string)
	var body ShortUrlJsonBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
		return
	}
	parsedUrl, err := url.ParseRequestURI(body.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid url"})
		return
	}
	uuidV4 := uuid.New()
	fmt.Println(uuidV4, parsedUrl)
	c.JSON(http.StatusCreated, gin.H{"original_url": body.Url, "short_url": uuidV4})
	// c.Redirect(http.StatusPermanentRedirect, "http://www.google.com")
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
		headerParser := router.Group("/header-parser", middlewareheaderParserAnalytics(analyticsDb))
		headerParser.GET("/whoami", whoAmIEndpoint)
	}
	// URL Shortener Routes
	{
		urlShortener := router.Group("/url-shortener") // add mdw for analytics?
		urlShortener.POST("/", POSTUrlShortenerEndpoint)
		urlShortener.GET("/:uuid", GETUrlShortenerEndpoint)
	}
}
