package jsend

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

const (
	StatusSuccess = "success"
	StatusError   = "error"
	StatusFail    = "fail"
)

type Response struct {
	CTX        *gin.Context `json:"-"`
	StatusCode int          `json:"-"`
	Status     string       `json:"status"`
	Data       interface{}  `json:"data"`
	Message    string       `json:"message,omitempty"`
	Code       int          `json:"code,omitempty"`
}

// R is a shortcut for Response
type R = Response

// JSON implements the JSend specification (https://labs.omniti.com/labs/jsend).
func (resp R) JSON() {
	c := resp.CTX

	var respJSON Response
	if resp.StatusCode >= http.StatusBadRequest {
		respJSON = Response{
			Status:  getStatus(resp.StatusCode),
			Message: resp.Message,
			Data:    formatError(resp.Data,resp.Code),
			Code:    resp.StatusCode,
		}
	} else {
		respJSON = Response{
			Status:  getStatus(resp.StatusCode),
			Message: resp.Message,
			Data:    resp.Data,
			Code:    resp.StatusCode,
		}
	}

	c.JSON(resp.StatusCode, respJSON)
}

// getStatus
func getStatus(code int) string {
	switch {
	case code >= 500:
		return StatusError
	case code >= 400 && code < 500:
		return StatusFail
	}

	return StatusSuccess
}

// https://github.com/gin-gonic/gin/issues/1372
func formatError(data interface{},code int) map[string]string {
	errMsg := make(map[string]string)

	switch it := data.(type) {
	case validator.ValidationErrors:
		for _, e := range it {
			jsonKey := strings.ToLower(e.Field)
			message := e.Tag
			errMsg[jsonKey] = message
		}
	case error:
		errMsg["error"] = it.Error()
		errMsg["code"] = string(code)
	}

	return errMsg
}
