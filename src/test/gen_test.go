package test

import (
	"MVC_DI/config"
	"MVC_DI/util/gen"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Gen(t *testing.T) {
	gen.Generate("MVC_DI", []string{"user"})
}
func Test_GenQuery(t *testing.T) {
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		t.Errorf("connect to database failed: %v", err)
	}
	gen.GenerateQuery([]string{"user"}, gormDB)
}

func Test_MVC(t *testing.T) {
	gen.GenerateMVC("MVC_DI", "section", "auth", []string{"auth"})
	gen.GenerateMVC("MVC_DI", "section", "user", []string{"user"})
	gen.GenerateMVC("MVC_DI", "section", "credential", []string{"credential"})
	gen.GenerateMVC("MVC_DI", "section", "permission", []string{"permission"})
	gen.GenerateMVC("MVC_DI", "section", "security", []string{"security"})
}

func Test_GenController(t *testing.T) {
	gen.GenerateGinController("MVC_DI", "section", "test", []string{"test_a"})
}

func Test_GenService(t *testing.T) {
	gen.GenerateService("MVC_DI", "section", "test", []string{"test_a"})
}

func Test_GenMapper(t *testing.T) {
	gen.GenerateMapper("MVC_DI", "section", "test", []string{"test_a"})
}
