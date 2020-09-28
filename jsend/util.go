package jsend

import (
	"github.com/CSystem/gcuu/merrors"
	"net/http"
	"github.com/gin-gonic/gin"
)

func JSONStatusOK(c *gin.Context, data interface{}) {
	resp := R{CTX: c, StatusCode: http.StatusOK, Data: data}
	resp.JSON()
}

func JSONStatusCreated(c *gin.Context, data interface{}) {
	resp := R{CTX: c, StatusCode: http.StatusCreated, Data: data}
	resp.JSON()
}

func JSONStatusBadRequest(c *gin.Context, data interface{}) {
	resp := R{CTX: c, StatusCode: http.StatusBadRequest, Data: data}
	resp.JSON()
}

func JSONStatusServerError(c *gin.Context, data interface{}, message string, code int) {
	resp := R{CTX: c, StatusCode: http.StatusInternalServerError, Data: data, Message: message, Code: code}
	resp.JSON()
}

func JSONStatusServerMError(c *gin.Context, err error) {
	mError := err.(*merrors.MError)
	resp := R{CTX: c, StatusCode: http.StatusInternalServerError, Data: err, Message: mError.Msg}
	resp.JSON()
}

func JSONStatusNotFound(c *gin.Context, data interface{}) {
	resp := R{CTX: c, StatusCode: http.StatusNotFound, Data: data}
	resp.JSON()
}
