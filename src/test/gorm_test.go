package test

import (
	"MVC_DI/config"
	"MVC_DI/util/gen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_GenQuery(t *testing.T) {
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		t.Errorf("connect to database failed: %v", err)
	}
	gen.GenerateQuery([]string{"user"}, gormDB)
}
