package parser

type DefaultValue interface {}

type ConfigFieldType int
const (
	SimpleConfigField ConfigFieldType = iota
	SliceConfigField
	StructConfigField
)

type ConfigField struct {
	name string
	fieldType ConfigFieldType
	defaultValue DefaultValue
	fields []ConfigField //
}

