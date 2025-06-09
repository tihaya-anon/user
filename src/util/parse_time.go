package util

import (
	config_model "MVC_DI/config/model"
	"time"
)

func GetTime(configTime config_model.Time) time.Duration {
	return time.Hour*time.Duration(configTime.Hour) + time.Minute*time.Duration(configTime.Minute) + time.Second*time.Duration(configTime.Second)
}
