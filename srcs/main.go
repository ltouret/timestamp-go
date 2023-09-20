package main

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	retcode := 0
	defer func() { os.Exit(retcode) }()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		retcode = 1
		return
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbIp := os.Getenv("DB_IP")
	dbPort := os.Getenv("DB_PORT")
	db, err := sql.Open("mysql", dbUser+dbPass+"@tcp("+dbIp+":"+dbPort+")/")
	if err != nil {
		fmt.Println("Connection Db", err)
		retcode = 1
		return
	}
	defer db.Close()
	defer fmt.Println("Hello")
	if err := SetupDb(db); err != nil {
		retcode = 1
		return
	}
	analyticsDb, err := NewAnalyticsService(db)
	if err != nil {
		fmt.Println("analyticsDb", err)
		retcode = 1
		return
	}
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	router := gin.Default()
	router.Use(gin.Recovery()) // useful for 500 error do i keep this?
	apiGroup := router.Group("/api")
	v1 := apiGroup.Group("/v1", middlewareTimestampAnalytics(analyticsDb))
	SetupRoutes(v1)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
