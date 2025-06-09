package gen

import (
	"MVC_DI/global"
	"MVC_DI/global/module"
	"MVC_DI/util"
	"path"
)

// # GenerateService
//
// generates Service and ServiceImpl
func GenerateService(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateService(pkg, basePath, entity, table)
	}
}
// # _generateService
//
// _generateService generates Service and ServiceImpl
//
// This function generates Service and ServiceImpl based on the provided
// package, base path, entity, and table.
func _generateService(pkg, basePath, entity, table string) {
	serviceTemplatePath := global.PATH.RESOURCE.TEMPLATE.SERVICE.DIR
	serviceDir := append([]string{module.GetSrc(), basePath, entity}, global.PATH.SERVICE.DIR...)
	util.CreateDir(path.Join(serviceDir...))

	// generate service interface
	interfaceTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.INTERFACE...)
	interfacePath := append(serviceDir, global.PATH.SERVICE.INTERFACE...)

	GenerateTemplate(pkg, path.Join(interfaceTemplatePath...), path.Join(interfacePath...), "_service", entity, table)

	// generate service builder
	builderTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.BUILDER...)
	builderPath := append(serviceDir, global.PATH.SERVICE.BUILDER...)

	GenerateTemplate(pkg, path.Join(builderTemplatePath...), path.Join(builderPath...), "_service_builder", entity, table)

	// generate service impl
	implTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.IMPL...)
	implPath := append(serviceDir, global.PATH.SERVICE.IMPL...)

	GenerateTemplate(pkg, path.Join(implTemplatePath...), path.Join(implPath...), "_service_impl", entity, table)

	// generate service test
	testTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.TEST...)
	testPath := append(serviceDir, global.PATH.SERVICE.TEST...)

	GenerateTemplate(pkg, path.Join(testTemplatePath...), path.Join(testPath...), "_service_test", entity, table)
}
