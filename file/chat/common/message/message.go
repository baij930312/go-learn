package message

const (
	LoginMesType    = "LoginMes"
	RegisterMesType = " RegisterMes"
	ResMesType      = "ResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}
type RegisterMes struct {
	UserId   int    `json:"userId"`
	Password string `json:"password"`
	UserName string `json:"userName"`
}

type ResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
