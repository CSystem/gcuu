package jsend

import (
	"github.com/CSystem/gcuu/merrors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JSONStatusOK(c *gin.Context, data interface{}) {
	data.(gin.H)["t"] = time.Now().Unix()
	resp := R{CTX: c, StatusCode: http.StatusOK, Data: data}
	resp.JSON()
}

func JSONStatusCreated(c *gin.Context, data interface{}) {
	data.(gin.H)["t"] = time.Now().Unix()
	resp := R{CTX: c, StatusCode: http.StatusCreated, Data: data}
	resp.JSON()
}

func JSONStatusBadRequest(c *gin.Context, data interface{}) {
	resp := R{CTX: c, StatusCode: http.StatusBadRequest, Data: data}
	resp.JSON()
}

func JSONStatusServerError(c *gin.Context, data interface{}, message string, code int) {
	data.(gin.H)["t"] = time.Now().Unix()
	resp := R{CTX: c, StatusCode: http.StatusInternalServerError, Data: data, Message: message, Code: code}
	resp.JSON()
}

func JSONStatusServerMError(c *gin.Context,data interface{}, err error) {
	mError := err.(*merrors.MError)
	data.(gin.H)["t"] = time.Now().Unix()
	data.(gin.H)["error"] = mError.SMsg
	resp := R{CTX: c, StatusCode: http.StatusInternalServerError, Data: data, Message: mError.Msg, Code: mError.Code}
	resp.JSON()
}

func JSONStatusNotFound(c *gin.Context, data interface{}) {
	data.(gin.H)["t"] = time.Now().Unix()
	resp := R{CTX: c, StatusCode: http.StatusNotFound, Data: data}
	resp.JSON()
}
