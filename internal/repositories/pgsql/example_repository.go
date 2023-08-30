package pgsql

import (
	"context"
	"time"

	"github.com/upstars-global/go-service-skeleton/internal/entity"
	"github.com/upstars-global/go-service-skeleton/internal/repositories/pgsql/requester"
	"github.com/upstars-global/go-service-skeleton/internal/usecases/example"
)

type exampleRepository struct {
	db requester.Querier
}

func NewExampleRepository(db requester.Querier) example.RepositoryInterface {
	return &exampleRepository{
		db: db,
	}
}
func (b exampleRepository) ExampleWriterMethod(ctx context.Context) (err error) {
	_, err = b.db.GetCurrentTime(ctx)
	return err
}

func (b exampleRepository) ExampleReaderMethod(ctx context.Context) (example *entity.ExampleEntity, err error) {
	example = &entity.ExampleEntity{}
	dbDate, err := b.db.GetCurrentTime(ctx)
	if err != nil {
		return
	}
	example.Date = dbDate.(time.Time)
	return
}
