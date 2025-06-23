package auth_dto

type UserLoginDto struct {
	Secret     string
	Identifier string
	Type       string
}
