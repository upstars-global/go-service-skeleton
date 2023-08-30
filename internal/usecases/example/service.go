package example

import (
	"context"

	"github.com/upstars-global/go-service-skeleton/internal/entity"
)

type service struct {
	repo RepositoryInterface
}

func NewService(r RepositoryInterface) UseCaseInterface {
	return &service{
		repo: r,
	}
}

func (s service) ExampleUseCaseMethod(ctx context.Context) (item *entity.ExampleEntity, err error) {
	item = &entity.ExampleEntity{}
	dbDateEntity, err := s.repo.ExampleReaderMethod(ctx)
	item.Date = dbDateEntity.Date.AddDate(1, 1, 1)
	return
}
