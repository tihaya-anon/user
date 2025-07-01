package enum

type Code string
type Msg string
type Status struct {
	code Code
	msg  Msg
}

func NewStatus(code Code, msg Msg) *Status {
	return &Status{code: code, msg: msg}
}
func (s *Status) Code() Code {
	return s.code
}

func (s *Status) Msg() Msg {
	return s.msg
}

type StatusMap map[any]*Status
