package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	// Message Type
	Type string `json:"type"`
	// Message content
	Data string `json:"data"`
}

type LoginMes struct {
	// User ID
	UserId int `json:"user_id"`
	// User Password
	UserPwd string `json:"user_pwd"`
	// User Name
	// UserName string `json:"user_name"`
}

type LoginResMes struct {
	// Code 200 indicates success, 500 indicates failure
	Code int `json:"code"`
	// Error message
	Error string `json:"error"`
}
