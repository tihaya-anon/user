package gen

import (
	"MVC_DI/global/module"
	"MVC_DI/util"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

// # GenerateTemplate
//
// generates the code file based on the given template
// and table information.
//
// It will replace the following placeholders in the template with the
// given table information:
//
//   - entity_name
//   - TableName
//   - tableName
//   - table_name
//   - table_name_hyphen
//   - pkg
//
// The generated file will be saved in the given targetPath with the name
// {table_name}{targetFilePostfix}.go
func GenerateTemplate(pkg, templatePath, targetPath, targetFilePostfix, entity, table string) {
	templatePath = path.Join(filepath.Dir(module.GetSrc()), templatePath)
	templateFile, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("read `%v` failed: %v", templatePath, err)
	}
	coreTemplate := string(templateFile)

	tmpl := template.Must(template.New("template").Parse(coreTemplate))
	util.CreateDir(targetPath)

	targetFile := path.Join(targetPath, table+targetFilePostfix+".go")
	file, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("create `%v` file failed: %v", targetFile, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]interface{}{
		"entity_name":       entity,
		"EntityName":        util.SnakeToPascal(entity),
		"TableName":         util.SnakeToPascal(table),
		"tableName":         util.SnakeToCamel(table),
		"table_name":        table,
		"table_name_hyphen": util.SnakeToHyphen(table),
		"pkg":               pkg,
	}); err != nil {
		log.Fatalf("generate `%v` failed: %v", targetPath, err)
	}
}
