package gen

import (
	"MVC_DI/global"
	"MVC_DI/global/module"
	"MVC_DI/util"
	"path"
)

// # GenerateMapper
//
// generates Mapper and MapperImpl
func GenerateMapper(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateMapper(pkg, basePath, entity, table)
	}
}

func _generateMapper(pkg, basePath, entity, table string) {
	mapperTemplatePath := global.PATH.RESOURCE.TEMPLATE.MAPPER.DIR
	mapperDir := append([]string{module.GetSrc(), basePath, entity}, global.PATH.MAPPER.DIR...)
	util.CreateDir(path.Join(mapperDir...))

	interfaceTemplatePath := append(mapperTemplatePath, global.PATH.RESOURCE.TEMPLATE.MAPPER.INTERFACE...)
	interfacePath := append(mapperDir, global.PATH.MAPPER.INTERFACE...)

	GenerateTemplate(pkg, path.Join(interfaceTemplatePath...), path.Join(interfacePath...), "_mapper", entity, table)

	implTemplatePath := append(mapperTemplatePath, global.PATH.RESOURCE.TEMPLATE.MAPPER.IMPL...)
	implPath := append(mapperDir, global.PATH.MAPPER.IMPL...)

	GenerateTemplate(pkg, path.Join(implTemplatePath...), path.Join(implPath...), "_mapper_impl", entity, table)

	builderTemplatePath := append(mapperTemplatePath, global.PATH.RESOURCE.TEMPLATE.MAPPER.BUILDER...)
	builderPath := append(mapperDir, global.PATH.MAPPER.BUILDER...)

	GenerateTemplate(pkg, path.Join(builderTemplatePath...), path.Join(builderPath...), "_mapper_builder", entity, table)
}
