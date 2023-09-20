// analytics_service.go

package main

import "database/sql"

type AnalyticsService struct {
	db *sql.DB
}

func NewAnalyticsService(db *sql.DB) (*AnalyticsService, error) {
	return &AnalyticsService{db}, nil
}
