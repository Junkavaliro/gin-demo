package errcode

import (
	"fmt"
	"net/http"
)

var (
	Success 				= NewError(0, "成功")
	ServerError 			= NewError(10000000, "服务内部错误")
	InvalidParams 			= NewError(10000001, "入参错误")
	NotFound 				= NewError(10000002, "找不到")
	UnauthorizedAuthNotExist = NewError(100000003, "鉴权失败，秘钥不存在")
	TooManyRequests 		= NewError(100000007, "访问过多")
)

type Error struct {
	code int `json:"code"`
	msg string `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	e.details = []string{}
	for _, d := range details {
		e.details = append(e.details, d)
	}
	return e
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}