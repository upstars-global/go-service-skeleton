package service

import (
	"net/http"

	"github.com/upstars-global/go-service-skeleton/api/server"
)

var (
	ErrInvalidIdentity    = server.MakeErr(http.StatusBadRequest, 1000, "invalid identity")
	ErrPasswordHashFailed = server.MakeErr(http.StatusInternalServerError, 1001, "failed to hash password")
	ErrCreateUserFailed   = server.MakeErr(http.StatusInternalServerError, 1002, "create user failed")
	ErrUserAlreadyExists  = server.MakeErr(http.StatusConflict, 1003, "user already exists")
	ErrUserNotFound       = server.MakeErr(http.StatusForbidden, 1004, "user not found")
	ErrInvalidPassword    = server.MakeErr(http.StatusForbidden, 1006, "password is invalid")
	ErrUserIsInactive     = server.MakeErr(http.StatusForbidden, 1011, "user is inactive")
)
