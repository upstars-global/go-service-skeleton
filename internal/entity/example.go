package entity

import (
	"time"

	validate "github.com/go-playground/validator/v10"
)

type ExampleEntity struct {
	Date time.Time
}

func (u *ExampleEntity) Validate() error {
	return validate.New().Struct(u)
}
