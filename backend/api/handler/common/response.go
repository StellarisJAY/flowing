package common

import "errors"

var (
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Total   int64  `json:"total"`
}

func Ok() BaseResp {
	return BaseResp{
		Code:    200,
		Message: "success",
	}
}

func PageResp(data any, total int64) BaseResp {
	return BaseResp{
		Code:    200,
		Message: "success",
		Data:    data,
		Total:   total,
	}
}

func OkWithMessage(message string) BaseResp {
	return BaseResp{
		Code:    200,
		Message: message,
	}
}

func OkWithData(data any) BaseResp {
	return BaseResp{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func Error(err error) BaseResp {
	return BaseResp{
		Code:    500,
		Message: err.Error(),
	}
}
