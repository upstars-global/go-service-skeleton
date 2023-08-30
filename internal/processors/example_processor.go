package processors

import (
	"context"

	"github.com/upstars-global/go-service-skeleton/internal/presenters"
	"github.com/upstars-global/go-service-skeleton/internal/usecases/example"
)

type ExampleProcessorInterface interface {
	ExampleProcessorMethod(ctx context.Context) (presenter presenters.Example, err error)
}

type exampleProcessor struct {
	useCase example.UseCaseInterface
}

func NewExampleProcessor(useCase example.UseCaseInterface) ExampleProcessorInterface {
	return &exampleProcessor{
		useCase: useCase,
	}
}

func (b exampleProcessor) ExampleProcessorMethod(ctx context.Context) (presenter presenters.Example, err error) {
	exampleEnty, err := b.useCase.ExampleUseCaseMethod(ctx)
	if err != nil {
		return
	}
	presenter.MapPresenter(exampleEnty)
	return
}
