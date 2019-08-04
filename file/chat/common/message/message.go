package message

const (
	LoginMesType            = "LoginMes"
	RegisterMesType         = " RegisterMes"
	ResMesType              = "ResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
)

const (
	UserOnline = iota
	UserOffLine
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

type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type ResMes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  string `json:"data"`
}
