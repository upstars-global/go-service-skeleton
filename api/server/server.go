package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EndpointsRegistrar func(e *gin.Engine)

type ServiceRegistrar interface {
	Register(e *gin.Engine)
	Middlewares() []gin.HandlerFunc
	Address() string
	DebugEnabled() bool
	LogHTTPRequests() bool
}

type Server struct {
	*http.Server
	gin *Gin
}

func New(registrar ServiceRegistrar) (err error) {
	g := NewGin(registrar.DebugEnabled(), registrar.LogHTTPRequests())
	s := &Server{
		gin: g,
		Server: &http.Server{
			Addr:           registrar.Address(),
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:        g,
		},
	}

	for _, v := range registrar.Middlewares() {
		g.Use(v)
	}

	registrar.Register(g.Engine)

	return s.ListenAndServe()
}
