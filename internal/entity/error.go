package entity

import "errors"

var ErrNotFound = errors.New("not found")

var ErrInvalidEntity = errors.New("invalid entity")

var ErrCannotBeDeleted = errors.New("cannot be deleted")
