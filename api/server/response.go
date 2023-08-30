package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ErrBadRequest          = MakeErr(http.StatusBadRequest, 1, "bad request")
	ErrNotFound            = MakeErr(http.StatusNotFound, 2, "not found")
	ErrUnauthorized        = MakeErr(http.StatusUnauthorized, 3, "unauthorized")
	ErrForbidden           = MakeErr(http.StatusForbidden, 4, "forbidden")
	ErrConflict            = MakeErr(http.StatusConflict, 5, "conflict")
	ErrInternalServerError = MakeErr(http.StatusInternalServerError, 6, "internal server error")
)

type ResponseSenderInterface interface {
	Send(ctx *gin.Context)
}

type Err struct {
	status int16
	Error  struct {
		Code     int16     `json:"code"`
		Message  string    `json:"message"`
		DateTime time.Time `json:"datetime" format:"date-time"`
	} `json:"error"`
}

func (r *Err) Send(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(int(r.status), r)
}

type OK struct {
	status int16
	Data   interface{} `json:"data"`
}

func (r *OK) Send(ctx *gin.Context) {
	ctx.JSON(int(r.status), r)
}

type OKSet struct {
	status int16
	Data   interface{} `json:"data"`
	Meta   meta        `json:"meta"`
}

func (r *OKSet) Send(ctx *gin.Context) {
	ctx.JSON(int(r.status), r)
}

type meta struct {
	Pagination pagination `json:"pagination"`
}

type pagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Count  int `json:"count"`
}

func MakeErr(status, code int16, message string) *Err {
	return &Err{
		status: status,
		Error: struct {
			Code     int16     `json:"code"`
			Message  string    `json:"message"`
			DateTime time.Time `json:"datetime" format:"date-time"`
		}{
			Code:     code,
			Message:  message,
			DateTime: time.Now(),
		},
	}
}

func MakeOK(status int16, data interface{}) *OK {
	return &OK{
		status: status,
		Data:   data,
	}
}

func MakeOKSet(status int16, data interface{}, offset int, limit int, count int) *OKSet {
	return &OKSet{
		status: status,
		Data:   data,
		Meta: meta{
			Pagination: pagination{
				Offset: offset,
				Limit:  limit,
				Count:  count,
			},
		},
	}
}
