package gen

import (
	"MVC_DI/config"
	"MVC_DI/global/module"
	"log"
	"path"
	"strings"

	"github.com/google/uuid"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type ICommonMethod struct {
	ID int64
}

func (commonMethod *ICommonMethod) BeforeCreate(tx *gorm.DB) error {
	commonMethod.ID = int64(uuid.New().ID())
	return nil
}

// # Generate
//
// generate code for the given entities
//
// This function will generate query, service, service impl, mapper, mapper impl, and gin controller for the given entities.
func Generate(pkg string, entities []string) {
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to database failed: %v", err)
	}
	log.Printf("connect to: %v\n", config.Application.Database.Uri)
	basePath := "./section/"
	//  generate query
	GenerateQuery(entities, gormDB)
	for _, entity := range entities {

		// get all tables for the current entity
		tables := getEntityTables(gormDB, entity)

		//  generate mapper and mapper impl
		GenerateMapper(pkg, basePath, entity, tables)

		//  generate service and service impl
		GenerateService(pkg, basePath, entity, tables)

		//  generate gin controller
		GenerateGinController(pkg, basePath, entity, tables)
	}
}

// # getEntityTables
//
// returns all tables for the given entity
func getEntityTables(db *gorm.DB, entity string) []string {
	// execute SHOW TABLES LIKE 'entity%' to get all tables
	// that start with the given entity
	rows, err := db.Raw("SHOW TABLES LIKE '" + entity + "%'").Rows()
	if err != nil {
		log.Fatalf("get tables failed: %v", err)
	}
	defer rows.Close()

	// iterate over the result set and append the table names
	// to the tables slice
	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			log.Fatalf("scan table failed: %v", err)
		}
		tables = append(tables, table)
	}
	return tables
}

// # GenerateQuery
//
// generate Query for the given entity
//
// This function generates query code for the given entity. The code is generated
// in a temporary directory and then moved to the final location.
func GenerateQuery(entityList []string, gormDB *gorm.DB) {

	// initialize the generator
	g := gen.NewGenerator(gen.Config{
		OutPath: path.Join(module.GetSrc(), "database"),
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(gormDB)

	// generate Query code
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		for _, entity := range entityList {
			if strings.HasPrefix(tableName, entity) {
				return tableName
			}
		}
		return ""
	})

	// gen.WithMethod(CommonMethod{})
	g.ApplyBasic(g.GenerateAllTable(
		gen.FieldType("id", "int64"),
		gen.FieldJSONTag("id", "id"),
		gen.WithMethod(ICommonMethod{}))...)
	g.Execute()

}
