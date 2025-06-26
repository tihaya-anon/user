package schema

import (
	"reflect"

	"google.golang.org/protobuf/proto"
)

type ISchemaMapping struct {
	Schemas []Schema
}

type Schema struct {
	Proto    string
	Message  string
	AvscPath string
	Subject  string
}

func (sm *ISchemaMapping) GetSchemaByMessage(message string) *Schema {
	for _, schema := range sm.Schemas {
		if schema.Message == message {
			return &schema
		}
	}
	return nil
}

func (sm *ISchemaMapping) GetSchemaByObject(object proto.Message) *Schema {
	return sm.GetSchemaByMessage(getName(object))
}

func getName[T any](v T) string {
	t := reflect.TypeOf(v)

	// 递归去掉指针层
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	// 普通具名类型
	if name := t.Name(); name != "" {
		return name
	}

	// 匿名或内建类型：返回详细描述
	return t.String()
}
