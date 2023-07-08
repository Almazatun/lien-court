package inputs

type Login struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
}

type Register struct {
	Email string `json:"email"`
	Pass  string `json:"password"`
	Name  string `json:"username"`
}
