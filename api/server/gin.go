package server

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	*gin.Engine
}

func NewGin(debugEnabled, logHTTPRequests bool) (engine *Gin) {
	if !debugEnabled {
		gin.SetMode(gin.ReleaseMode)
	}

	engine = &Gin{gin.New()}
	engine.Use(gin.Recovery())
	if logHTTPRequests {
		engine.Use(logger.SetLogger())
	}
	engine.NoRoute(RouteNotFound)
	engine.GET("/", RouteHome)

	return
}
