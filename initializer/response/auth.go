package response

// JWT claims struct
type SignIn struct {
	ID uint
	AppId uint
	T int64
}

type TokenPayload struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}
