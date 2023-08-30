package presenters

import (
	"time"

	"github.com/upstars-global/go-service-skeleton/internal/entity"
)

type Example struct {
	Date time.Time `json:"date"`
}

func (u *Example) MapPresenter(example *entity.ExampleEntity) {
	u.Date = example.Date
}
