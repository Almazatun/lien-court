package inputs

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"password" validate:"required"`
}

type Register struct {
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"password validate:"required,min=4"`
	Name  string `json:"username" validate:"required,`
}

type LinkList struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type JWTData struct {
	ID    string
	Email string
}
