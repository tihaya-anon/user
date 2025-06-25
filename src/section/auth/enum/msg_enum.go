package auth_enum

var MSG = struct {
	PASSWORD_WRONG       string
	EMAIL_CODE_WRONG     string
	GOOGLE_2FA_WRONG     string
	OAUTH_WRONG          string
	UNKNOWN_CREDENTIAL   string
	UNKNOWN_TRIGGER_MODE string
}{
	PASSWORD_WRONG:       "pwd wrong",
	EMAIL_CODE_WRONG:     "email code wrong",
	GOOGLE_2FA_WRONG:     "google 2fa wrong",
	OAUTH_WRONG:          "oauth wrong",
	UNKNOWN_CREDENTIAL:   "unknown credential",
	UNKNOWN_TRIGGER_MODE: "unknown trigger mode",
}
