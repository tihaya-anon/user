package schema

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
