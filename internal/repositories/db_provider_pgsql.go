package repositories

import (
	"github.com/upstars-global/go-service-skeleton/internal/repositories/pgsql"
	"github.com/upstars-global/go-service-skeleton/internal/repositories/pgsql/requester"
)

var ProvidesRepositories = []interface{}{
	func(a *pgsql.DB) requester.DBTX { return a },
	func(a *pgsql.DB) pgsql.DBTXer { return a },
	func(a *requester.Queries) requester.Querier { return a },

	requester.New,
	pgsql.New,

	pgsql.NewExampleRepository,
}
