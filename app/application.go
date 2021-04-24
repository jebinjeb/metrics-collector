package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jebinjeb/metrics-collector/controllers/health"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}

func mapUrls() {
	router.GET("/health", health.Health)
	router.GET("/metrics", metrics.Metrics)
}
