package server

import "github.com/gin-gonic/gin"

func RouteNotFound(ctx *gin.Context) {
	MakeErr(404, 1, "not found").Send(ctx)
}

func RouteHome(ctx *gin.Context) {
	MakeOK(200, "welcome").Send(ctx)
}
