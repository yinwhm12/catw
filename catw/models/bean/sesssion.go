package bean

type CreateSession struct {
	Email string `json:"email"`
	Pwd string `json:"pwd"`
}

type OutPutSession struct {
	Uid int `json:"uid"`
	Token string `json:"token"`
	Email string `json:"email"`

}
