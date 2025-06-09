package config

import config_model "MVC_DI/config/model"

var Application = &config_model.IApplication{}

// # init
//
// init is called after all the variable declarations in the
// package have evaluated their initializers, and those are
// evaluated only after all the imported packages have been
// loaded. It is used here to read the application configuration
// from a YAML file and to resolve any placeholders in that
// configuration.
func init() {
	Parse("application", "", Application)
	Parse("application", Application.Env, Application)
	Resolve(Application)
}
