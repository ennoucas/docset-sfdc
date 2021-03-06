package main

var SupportedTypes = []SupportedType{
	// ID Based overrides should come first
	SupportedType{
		ID:       "ref_tag_set_attr_intf",
		TypeName: "Guide",
	},
	SupportedType{
		ID:       "namespaces_intro",
		TypeName: "Guide",
	},
	SupportedType{
		ID:       "namespaces_using_organization",
		TypeName: "Guide",
	},
	SupportedType{
		TypeName:    "Guide",
		ID:          "apex_intro_get_started",
		CascadeType: true,
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "pages_flows_customize_runtime_ui",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "pages_quick_start_controller_shell",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "pages_email_custom_controller",
	},
	SupportedType{
		TypeName:    "Guide",
		IDPrefix:    "apex_qs_",
		CascadeType: true,
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "apex_process_plugin_using",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "apex_platform_cache_builder",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "apex_classes_restful_http_testing_httpcalloutmock",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "apex_classes_namespaces_and_invoking_methods",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "apex_classes_schema_namespace_using",
	},
	// Apex types
	SupportedType{
		TypeName:      "Method",
		TitleSuffix:   "Methods",
		AppendParents: true,
		IsContainer:   true,
		ShowNamespace: true,
	},
	SupportedType{
		TypeName:      "Constructor",
		TitleSuffix:   "Constructors",
		AppendParents: true,
		IsContainer:   true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Class",
		TitleSuffix:   "Class",
		PushName:      true,
		AppendParents: true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Namespace",
		TitleSuffix:   "Namespace",
		PushName:      true,
		AppendParents: true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Interface",
		TitleSuffix:   "Global Interface",
		PushName:      true,
		AppendParents: true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Interface",
		TitleSuffix:   "Interface",
		PushName:      true,
		AppendParents: true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Statement",
		ID:            "langCon_apex_SOQL_query_all_rows",
		TitleOverride: "ALL ROWS",
	},
	SupportedType{
		TypeName:      "Statement",
		TitleSuffix:   "Statement",
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Enum",
		TitleSuffix:   "Enum",
		AppendParents: true,
		ShowNamespace: true,
	},
	SupportedType{
		TypeName:      "Property",
		TitleSuffix:   "Properties",
		AppendParents: true,
		IsContainer:   true,
		ShowNamespace: true,
	},
	SupportedType{
		TypeName:      "Guide",
		TitleSuffix:   "Example Implementation",
		NoTrim:        true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Statement",
		TitleSuffix:   "Statements",
		NoTrim:        true,
		AppendParents: false,
		IsContainer:   true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Field",
		TitleSuffix:   "Fields",
		AppendParents: true,
		PushName:      true,
		IsContainer:   true,
		ShowNamespace: true,
	},
	SupportedType{
		TypeName:      "Exception",
		TitleSuffix:   "Exceptions",
		NoTrim:        true,
		AppendParents: true,
		ShowNamespace: false,
		ParseContent:  true,
	},
	SupportedType{
		TypeName:      "Constant",
		TitleSuffix:   "Constants",
		NoTrim:        true,
		AppendParents: true,
		ShowNamespace: false,
		ParseContent:  true,
	},
	SupportedType{
		TypeName:      "Class",
		TitleSuffix:   "Class (Base Email Methods)",
		PushName:      true,
		AppendParents: true,
		ShowNamespace: false,
	},
	SupportedType{
		TypeName:      "Guide",
		TitlePrefix:   "Best Practices",
		TitleSuffix:   "Best Practices",
		NoTrim:        true,
		PushName:      false,
		AppendParents: false,
		ShowNamespace: false,
	},
	// VF Types
	SupportedType{
		IDPrefix: "pages_compref_",
		TypeName: "Tag",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_maps",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_dynamic_vf",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_comp_cust",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_resources",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_controller",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_styling",
		NoTrim:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_security",
		NoTrim:   true,
	},
	SupportedType{
		TypeName:      "Variables",
		TitleSuffix:   "Global Variables",
		NoTrim:        true,
		AppendParents: true,
		ShowNamespace: false,
		ParseContent:  true,
		IsContainer:   true,
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_variables_functions",
	},
	SupportedType{
		TypeName: "Guide",
		IDPrefix: "pages_variables_operators",
	},
	// Aurora components
	SupportedType{
		TypeName:    "Tag",
		ID:          "aura_compref",
		IsContainer: true,
		CascadeType: true,
	},
	SupportedType{
		TypeName:    "Tag",
		ID:          "ref_messaging",
		IsContainer: true,
		CascadeType: true,
	},
	SupportedType{
		TypeName:    "Interface",
		ID:          "ref_interfaces",
		IsContainer: true,
		CascadeType: true,
	},
	SupportedType{
		TypeName:    "Event",
		ID:          "ref_events",
		IsContainer: true,
		CascadeType: true,
	},
	SupportedType{
		TypeName:    "Event",
		ID:          "ref_events_aura",
		IsContainer: true,
		CascadeType: true,
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "debug_intro",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "components_using",
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "components_overview",
	},
	SupportedType{
		TypeName:    "Guide",
		IDPrefix:    "qs_intro",
		IsContainer: true,
		CascadeType: true,
	},
	SupportedType{
		TypeName: "Guide",
		ID:       "events_intro",
	},
	SupportedType{
		TypeName:    "Guide",
		ID:          "apps_intro",
		IsHidden:    true,
		CascadeType: true,
	},
}
