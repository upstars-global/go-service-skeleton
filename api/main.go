package main

import (
	"context"
	"time"

	"github.com/upstars-global/go-service-skeleton/api/server"
	"github.com/upstars-global/go-service-skeleton/api/service"
	"github.com/upstars-global/go-service-skeleton/internal"
	"github.com/upstars-global/go-service-skeleton/pkg"
	"github.com/upstars-global/go-service-skeleton/pkg/config"
	"github.com/upstars-global/go-service-skeleton/pkg/logger"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

func main() {
	provides := []interface{}{
		context.Background,
	}
	provides = append(provides, internal.GetProviders()...)
	provides = append(provides, pkg.GetProviders()...)
	provides = append(provides, service.Provides...)

	app := fx.New(
		fx.Provide(provides...),
		fx.Invoke(func(
			ctx context.Context,
			publicSvc *service.Service,
			cfg *config.Config,
			log logger.Interface,
		) {
			var group errgroup.Group
			group.Go(func() error {
				log.Debugf("registered public api server: %s", cfg.GetAPIAddress())
				return server.New(publicSvc)
			})
		}),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		panic(err)
	}
	<-app.Done()
}
