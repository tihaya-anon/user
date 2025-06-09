package global

var PATH = Path{
	RESOURCE: Resource{
		TEMPLATE: Template{
			CONTROLLER: Controller{
				DIR:     []string{"resource", "template", "controller"},
				BUILDER: []string{"controller_builder.tmpl"},
				CORE:    []string{"controller_core.tmpl"},
				ROUTER:  []string{"controller_router.tmpl"},
				TEST:    []string{"controller_test.tmpl"},
			},
			MAPPER: Mapper{
				DIR:       []string{"resource", "template", "mapper"},
				INTERFACE: []string{"mapper_interface.tmpl"},
				BUILDER:   []string{"mapper_builder.tmpl"},
				IMPL:      []string{"mapper_impl.tmpl"},
			},
			SERVICE: Service{
				DIR:       []string{"resource", "template", "service"},
				INTERFACE: []string{"service_interface.tmpl"},
				BUILDER:   []string{"service_builder.tmpl"},
				IMPL:      []string{"service_impl.tmpl"},
				TEST:      []string{"service_test.tmpl"},
			},
		},
	},
	CONTROLLER: Controller{
		BUILDER: []string{"controller", "builder"},
		CORE:    []string{"controller"},
		TEST:    []string{"controller", "test"},
	},
	MAPPER: Mapper{
		INTERFACE: []string{"mapper"},
		BUILDER:   []string{"mapper", "builder"},
		IMPL:      []string{"mapper", "impl"},
	},
	SERVICE: Service{
		INTERFACE: []string{"service"},
		BUILDER:   []string{"service", "builder"},
		IMPL:      []string{"service", "impl"},
		TEST:      []string{"service", "test"},
	},
}

type Path = struct {
	RESOURCE   Resource
	CONTROLLER Controller
	MAPPER     Mapper
	SERVICE    Service
}

type Resource = struct {
	TEMPLATE Template
}

type Template = struct {
	CONTROLLER Controller
	MAPPER     Mapper
	SERVICE    Service
}

type Controller = struct {
	DIR     []string
	BUILDER []string
	CORE    []string
	ROUTER  []string
	TEST    []string
}

type Mapper = struct {
	DIR       []string
	INTERFACE []string
	BUILDER   []string
	IMPL      []string
}

type Service = struct {
	DIR       []string
	BUILDER   []string
	INTERFACE []string
	IMPL      []string
	TEST      []string
}
