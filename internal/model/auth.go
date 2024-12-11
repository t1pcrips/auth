package model

type User struct {
	Id       int64
	Email    string
	Password string
	Username string
	Role     string
}

type Tokens struct {
	RefreshToken string
	AccessToken  string
}

type RefreshToken struct {
	Token string
}
