package model

type app struct {
	Host string
	Port int
	Uri  string
}

type database struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
	Uri      string
}

type jwt struct {
	Secret     string
	Expiration Time
}

type schemaRegistry struct {
	Host string
	Port int
}

type Time struct {
	Hour   int
	Minute int
	Second int
}

type IApplication struct {
	Database       database
	App            app
	Jwt            jwt
	Env            string
	SchemaRegistry schemaRegistry
}
