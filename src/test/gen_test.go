package test

import (
	"MVC_DI/util/gen"
	"testing"
)

func Test_Gen(t *testing.T) {
	gen.Generate("MVC_DI", []string{"user"})
}
