package pkg

import (
	"github.com/upstars-global/go-service-skeleton/pkg/argumentsresolver"
	"github.com/upstars-global/go-service-skeleton/pkg/config"
	"github.com/upstars-global/go-service-skeleton/pkg/logger"
)

func GetProviders() []interface{} {
	return []interface{}{
		func(a *config.Config) config.GeneralConfigProvider { return a },
		func(a *config.Config) config.LoggerConfigProvider { return a },
		func(a *config.Config) config.DBConfigProvider { return a },
		func(a *config.Config) config.APIPConfigProvider { return a },

		argumentsresolver.New,
		config.New,
		logger.New,
	}
}
