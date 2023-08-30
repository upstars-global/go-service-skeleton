package internal

import (
	"github.com/upstars-global/go-service-skeleton/internal/processors"
	"github.com/upstars-global/go-service-skeleton/internal/repositories"
	"github.com/upstars-global/go-service-skeleton/internal/usecases/example"
)

func GetProviders() []interface{} {
	provides := []interface{}{
		example.NewService,
	}
	provides = append(provides, processors.Provides...)
	provides = append(provides, repositories.ProvidesRepositories...)
	return provides
}
