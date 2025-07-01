package dto

import "MVC_DI/gen/proto"

type PublishLoginAuditDto struct {
	UserId     int64
	IpAddress  string
	DeviceInfo string
	Result     proto.LoginResult
}
