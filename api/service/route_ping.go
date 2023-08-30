package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/upstars-global/go-service-skeleton/pkg/logger"
)

type routePing struct {
	log logger.Interface
}

func newRoutePing(log logger.Interface) *routePing {
	return &routePing{log: log}
}

func (r *routePing) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
