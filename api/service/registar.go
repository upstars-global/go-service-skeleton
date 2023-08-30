package service

import (
	"github.com/gin-gonic/gin"
	"github.com/upstars-global/go-service-skeleton/api/middleware"
	"github.com/upstars-global/go-service-skeleton/internal/repositories/pgsql/requester"
	"github.com/upstars-global/go-service-skeleton/pkg/config"
	"github.com/upstars-global/go-service-skeleton/pkg/logger"
)

type Service struct {
	cfg *config.Config
	db  requester.Querier
	log logger.Interface

	routeGetUser *routePing
}

func newService(
	cfg *config.Config,
	db requester.Querier,
	log logger.Interface,

	routeGetUserFn *routePing,
) *Service {
	return &Service{
		cfg: cfg,
		db:  db,
		log: log,

		routeGetUser: routeGetUserFn,
	}
}

func (s Service) Address() string {
	return s.cfg.GetAPIAddress()
}

func (s *Service) DebugEnabled() bool {
	return s.cfg.DebugEnabled()
}

func (s *Service) LogHTTPRequests() bool {
	return s.cfg.GetLogHttpRequests()
}

func (s Service) Register(e *gin.Engine) {
	e.GET("ping", s.routeGetUser.Handle)
}

func (s Service) Middlewares() (mws []gin.HandlerFunc) {
	mws = make([]gin.HandlerFunc, 0)

	mws = append(mws, middleware.CORS())

	return
}
