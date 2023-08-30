package main

import (
	"context"
	"fmt"

	"github.com/upstars-global/go-service-skeleton/internal"
	"github.com/upstars-global/go-service-skeleton/internal/processors"
	"github.com/upstars-global/go-service-skeleton/pkg"
	"github.com/upstars-global/go-service-skeleton/pkg/logger"
	"go.uber.org/fx"
)

func main() {
	provides := []interface{}{
		context.Background,
	}
	provides = append(provides, pkg.GetProviders()...)
	provides = append(provides, internal.GetProviders()...)

	fx.New(
		fx.Provide(provides...),
		fx.Invoke(func(
			ctx context.Context,
			log logger.Interface,
			exProcessor processors.ExampleProcessorInterface,
		) {
			date, err := exProcessor.ExampleProcessorMethod(ctx)
			if err != nil {
				log.With("err", err).Error("can't execute processor request")
			}
			fmt.Println(date.Date)
		}),
	)
}
