package gen

import (
	"MVC_DI/global"
	"MVC_DI/global/module"
	"MVC_DI/util"
	"path"
)

// # GenerateGinController
//
// generates the core and builder files for a Gin controller
//
// based on the provided package, base path, entity, and tables.
func GenerateGinController(pkg, basePath, entity string, tables []string) {
	// Iterate over each table in the provided list of tables
	// and generate the core and builder files for the controller
	for _, table := range tables {
		_generateGinController(pkg, basePath, entity, table)
	}
}

// # _generateGinController
//
// generates the core and builder files for a Gin controller
//
// based on the provided package, base path, entity, and table.
func _generateGinController(pkg, basePath, entity, table string) {
	// Set the template path for the controller
	controllerTemplatePath := global.PATH.RESOURCE.TEMPLATE.CONTROLLER.DIR
	// Create the directory path for the controller
	controllerDir := append([]string{module.GetSrc(), basePath, entity}, global.PATH.CONTROLLER.DIR...)
	util.CreateDir(path.Join(controllerDir...))

	// Generate the core controller file
	coreTemplatePath := append(controllerTemplatePath, global.PATH.RESOURCE.TEMPLATE.CONTROLLER.CORE...)
	corePath := append(controllerDir, global.PATH.CONTROLLER.CORE...)
	GenerateTemplate(pkg, path.Join(coreTemplatePath...), path.Join(corePath...), "_controller", entity, table)

	// Generate the builder controller file
	builderTemplatePath := append(controllerTemplatePath, global.PATH.RESOURCE.TEMPLATE.CONTROLLER.BUILDER...)
	builderPath := append(controllerDir, global.PATH.CONTROLLER.BUILDER...)
	GenerateTemplate(pkg, path.Join(builderTemplatePath...), path.Join(builderPath...), "_controller_builder", entity, table)

	// Generate the test controller file
	testTemplatePath := append(controllerTemplatePath, global.PATH.RESOURCE.TEMPLATE.CONTROLLER.TEST...)
	testPath := append(controllerDir, global.PATH.CONTROLLER.TEST...)
	GenerateTemplate(pkg, path.Join(testTemplatePath...), path.Join(testPath...), "_controller_test", entity, table)

	// Generate the router controller file
	routerTemplatePath := append(controllerTemplatePath, global.PATH.RESOURCE.TEMPLATE.CONTROLLER.ROUTER...)
	GenerateTemplate(pkg, path.Join(routerTemplatePath...), path.Join(module.GetSrc(), "router", entity), "_router", entity, table)
}
