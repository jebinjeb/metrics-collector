package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func metrics(c *gin.Context){
	c.String(http.StatusOK, format:"ok")
}