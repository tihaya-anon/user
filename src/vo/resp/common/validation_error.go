package common

type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}
