package test

import (
	"MVC_DI/util/gen"
	"testing"
)

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
