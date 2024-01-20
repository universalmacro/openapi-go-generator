package openapi

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/universalmacro/openapigogenerator/gen"
)

var (
	object = "object"
)

type Openapi struct {
	Components *Components     `yaml:"components"`
	Tags       []Tag           `yaml:"tags"`
	Paths      map[string]Path `yaml:"paths"`
}

type Path map[string]HttpMethod

type HttpMethod struct {
	Tags        []string `yaml:"tags"`
	Description *string  `yaml:"description"`
	OperationId *string  `yaml:"operationId"`
}

type RequestBody struct {
	Content map[string]Content `yaml:"content"`
}

type Content struct {
	ApplicationJson *ApplicationJson `yaml:"application/json"`
	Description     *string          `yaml:"description"`
}

type ApplicationJson struct {
	Schema *Schema `yaml:"schema"`
}

func (o Openapi) File() *jen.File {
	f := jen.NewFile("models")
	// var apis map[string]Api = make(map[string]Api)
	// for _, tag := range o.Tags {
	// 	apis[tag.Name] = Api{}
	// }
	for uri, path := range o.Paths {
		for method, httpMethod := range path {
			fmt.Println(method, uri, *httpMethod.OperationId)
		}
	}
	if o.Components != nil {
		if o.Components.Schemas != nil {
			for name, schema := range o.Components.Schemas {
				if schema.Statement() != nil {
					f.Add(jen.Type().Id(name).Add(schema.Statement()))
				}
				// Enum
				if schema.Type != nil && *schema.Type == "string" && schema.Enum != nil {
					en := gen.StringEnum{Id: name, Values: *schema.Enum}.Statement()
					for _, e := range en {
						f.Add(e)
					}
				}
			}
		}
	}
	return f
}

func (o Openapi) String() (s string) {
	s = `Openapi{}`
	return s
}

type Tag struct {
	Name string `yaml:"name"`
}

type Components struct {
	Schemas    map[string]Schema    `yaml:"schemas"`
	Parameters map[string]Parameter `yaml:"parameters"`
}
type Schema struct {
	Type       *string            `yaml:"type"`
	Ref        *string            `yaml:"$ref"`
	Properties *map[string]Schema `yaml:"properties"`
	Format     *string            `yaml:"format"`
	Default    *string            `yaml:"default"`
	Required   *[]string          `yaml:"required"`
	Enum       *[]string          `yaml:"enum"`
	Items      *Schema            `yaml:"items"`
}

func (s Schema) Statement() *jen.Statement {
	if s.Type != nil && *s.Type == object {
		if s.Properties == nil {
			panic("object without properties")
		}
		var fields []jen.Code
		for name, schema := range *s.Properties {
			fields = append(fields, generateField(schema, name, s.Required))
		}
		return jen.Struct(fields...)
	}
	return nil
}

type Parameter struct {
	Name        *string `yaml:"name"`
	In          *string `yaml:"in"`
	Description *string `yaml:"description"`
	Schema      *Schema `yaml:"schema"`
}

func capitalize(str string) string {
	bs := []byte(str)
	if len(bs) == 0 {
		return ""
	}
	bs[0] = byte(bs[0] - 32)
	return string(bs)
}

func generateField(schema Schema, name string, required *[]string) jen.Code {
	var field jen.Code
	if schema.Ref != nil {
		var p = generateFieldName(name, required)
		field = p.Id(refToId(*schema.Ref))
	} else if schema.Type != nil {
		fieldName := generateFieldName(name, required)
		switch *schema.Type {
		case "string":
			return fieldName.String()
		case "integer":
			return generateIntgerField(fieldName, schema.Format)
		case "number":
			return generateFloatField(fieldName, schema.Format)
		case "boolean":
			return fieldName.Bool()
		case "array":
			return generateArrayField(fieldName, schema, name, required)
		}
	}
	return field
}

func generateArrayField(fieldName *jen.Statement, schema Schema, name string, required *[]string) *jen.Statement {
	fieldName.Index()
	if schema.Items == nil {
		panic("array without items")
	} else {
		if schema.Items.Ref != nil {
			fieldName.Id(refToId(*schema.Items.Ref))
		}
	}
	return fieldName
}

func generateIntgerField(fieldName *jen.Statement, format *string) *jen.Statement {
	if format != nil {
		if *format == "int32" {
			fieldName.Int32()
		}
		fieldName.Int64()
	} else {
		fieldName.Int()
	}
	return fieldName
}

func generateFloatField(fieldName *jen.Statement, format *string) *jen.Statement {
	if format != nil {
		if *format == "float" {
			fieldName.Float32()
		}
		fieldName.Float64()
	} else {
		fieldName.Float64()
	}
	return fieldName
}

func generateFieldName(name string, required *[]string) *jen.Statement {
	var fieldName *jen.Statement = jen.Id(capitalize(name))
	if !isRequired(name, required) {
		fieldName.Op("*")
	}
	return fieldName
}

func refToId(ref string) string {
	s := strings.Split(ref, "/")
	if len(s) == 0 {
		panic("ref is empty")
	}
	return s[len(s)-1]
}

func isRequired(name string, required *[]string) bool {
	if required == nil {
		return false
	}
	for _, r := range *required {
		if r == name {
			return true
		}
	}
	return false
}
