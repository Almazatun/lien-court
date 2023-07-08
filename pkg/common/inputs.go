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

type LinkList struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type JWTData struct {
	ID    string
	Email string
}
