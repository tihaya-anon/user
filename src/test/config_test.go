package test

import (
	"MVC_DI/config"
	"fmt"
	"testing"
)

func Test_Parse(t *testing.T) {
	fmt.Println(config.Application)
	fmt.Println("hello")
	fmt.Println(config.Application.Database)
}
