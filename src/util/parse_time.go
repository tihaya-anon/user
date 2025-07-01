package util

import (
	"MVC_DI/config/model"
	"time"
)

func GetTime(configTime model.Time) time.Duration {
	return time.Hour*time.Duration(configTime.Hour) + time.Minute*time.Duration(configTime.Minute) + time.Second*time.Duration(configTime.Second)
}
