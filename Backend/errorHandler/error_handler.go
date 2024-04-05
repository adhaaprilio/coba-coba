package errorHandler

import (
	"backend/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error400 struct {
	Message string
}

type Error401 struct {
	Message string
}

type Error403 struct {
	Message string
}

type Error404 struct {
	Message string
}

type Error409 struct {
	Message string
}

type Error500 struct {
	Message string
}

func (e *Error400) Error() string {
	return e.Message
}

func (e *Error401) Error() string {
	return e.Message
}

func (e *Error403) Error() string {
	return e.Message
}

func (e *Error404) Error() string {
	return e.Message
}

func (e *Error409) Error() string {
	return e.Message
}

func (e *Error500) Error() string {
	return e.Message
}

func HandleError(c *gin.Context, err error) {
	var statusCode int
	log.Print(err.Error())
	switch err.(type) {
	case *Error400:
		statusCode = http.StatusBadRequest
	case *Error401:
		statusCode = http.StatusUnauthorized
	case *Error403:
		statusCode = http.StatusForbidden
	case *Error404:
		statusCode = http.StatusNotFound
	case *Error409:
		statusCode = http.StatusConflict
	case *Error500:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(helper.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
