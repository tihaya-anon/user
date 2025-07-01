package auth_enum

import "MVC_DI/global/enum"

type UNKNOWN_CREDENTIAL struct{}
type PASSWORD_WRONG struct{}
type EMAIL_CODE_WRONG struct{}
type GOOGLE_2FA_WRONG struct{}
type OAUTH_WRONG struct{}
type UNKNOWN_SESSION struct{}

const (
	CODE_UNKNOWN_CREDENTIAL enum.Code = "CRE00"
	MSG_UNKNOWN_CREDENTIAL  enum.Msg  = "Unknown Credential"
	CODE_PASSWORD_WRONG     enum.Code = "CRE01"
	MSG_PASSWORD_WRONG      enum.Msg  = "Password Wrong"
	CODE_EMAIL_CODE_WRONG   enum.Code = "CRE02"
	MSG_EMAIL_CODE_WRONG    enum.Msg  = "Email Code Wrong"
	CODE_GOOGLE_2FA_WRONG   enum.Code = "CRE03"
	MSG_GOOGLE_2FA_WRONG    enum.Msg  = "Google 2FA Wrong"
	CODE_OAUTH_WRONG        enum.Code = "CRE04"
	MSG_OAUTH_WRONG         enum.Msg  = "OAuth Wrong"
	CODE_UNKNOWN_SESSION    enum.Code = "SES00"
	MSG_UNKNOWN_SESSION     enum.Msg  = "Unknown Session"
)

var AUTH_STATUS_MAP enum.StatusMap = enum.StatusMap{
	UNKNOWN_CREDENTIAL{}: enum.NewStatus(CODE_UNKNOWN_CREDENTIAL, MSG_UNKNOWN_CREDENTIAL),
	PASSWORD_WRONG{}:     enum.NewStatus(CODE_PASSWORD_WRONG, MSG_PASSWORD_WRONG),
	EMAIL_CODE_WRONG{}:   enum.NewStatus(CODE_EMAIL_CODE_WRONG, MSG_EMAIL_CODE_WRONG),
	GOOGLE_2FA_WRONG{}:   enum.NewStatus(CODE_GOOGLE_2FA_WRONG, MSG_GOOGLE_2FA_WRONG),
	OAUTH_WRONG{}:        enum.NewStatus(CODE_OAUTH_WRONG, MSG_OAUTH_WRONG),
	UNKNOWN_SESSION{}:    enum.NewStatus(CODE_UNKNOWN_SESSION, MSG_UNKNOWN_SESSION),
}
