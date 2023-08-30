package example

import (
	"context"

	"github.com/upstars-global/go-service-skeleton/internal/entity"
)

type ReaderRepositoryInterface interface {
	ExampleReaderMethod(ctx context.Context) (example *entity.ExampleEntity, err error)
}

type WriterRepositoryInterface interface {
	ExampleWriterMethod(ctx context.Context) (err error)
}

type RepositoryInterface interface {
	ReaderRepositoryInterface
	WriterRepositoryInterface
}

type UseCaseInterface interface {
	ExampleUseCaseMethod(ctx context.Context) (item *entity.ExampleEntity, err error)
}
