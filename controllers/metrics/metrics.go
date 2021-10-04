package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Metrics(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
