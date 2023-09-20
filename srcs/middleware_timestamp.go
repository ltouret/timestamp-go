// middleware.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func truncateText(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:max]
}

func middlewareTimestampAnalytics(analyticsDb *AnalyticsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		responseTime := time.Since(t).String()
		timestamp := truncateText(t.Format(http.TimeFormat), 255)
		queryParameters := truncateText(c.Param("date"), 255) //? do i leave this null if no query?
		userAgent := truncateText(c.Request.UserAgent(), 255)
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		insertStatement := "INSERT INTO timestamp (timestamp, statusCode, queryParameters, userAgent, clientIP, responseTime) VALUES (?, ?, ?, ?, ?, ?)"
		_, err := analyticsDb.db.Exec(insertStatement, timestamp, statusCode, queryParameters, userAgent, clientIP, responseTime)
		if err != nil {
			fmt.Println("Error inserting data:", err)
		}
	}
}
