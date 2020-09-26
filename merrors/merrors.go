package merrors

import (
	"encoding/json"
)

type MError struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	SMsg error `json:"s_msg"`
}

func (e *MError) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int, msg string, sMsg error) *MError {
	return &MError{
		Code: code,
		Msg:   msg,
		SMsg: sMsg,
	}
}
