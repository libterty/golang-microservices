package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	health = "Server is healthy"
)

func HealthCheck(ctx *gin.Context) {
	ctx.String(http.StatusOK, health)
}
