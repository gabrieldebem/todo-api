package exceptions

import "github.com/gin-gonic/gin"

type Exception interface {
	GetCode() int
	GetMsg() string
}

type HttpException struct {
	Context *gin.Context
	Error   error
	Message string
	Code    int
}

func (e HttpException) GetCode() int {
	if e.Code != 0 {
		return e.Code
	}
	return 500
}

func (e HttpException) GetMsg() string {
	if e.Message != "" {
		return e.Message
	}

	return e.Error.Error()
}

func (e HttpException) Throw() {
	e.Context.JSON(e.GetCode(), gin.H{
		"message": e.GetMsg(),
	})
}
