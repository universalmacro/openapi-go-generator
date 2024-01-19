package openapi

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

var (
	object = "object"
)

type Openapi struct {
	Components *Components `yaml:"components"`
}

func (o Openapi) File() *jen.File {
	f := jen.NewFile("models")
	if o.Components != nil {
		if o.Components.Schemas != nil {
			for name, schema := range o.Components.Schemas {
				if schema.Statement() != nil {
					f.Add(jen.Type().Id(name).Add(schema.Statement()))
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
}

func (s Schema) Statement() *jen.Statement {
	if s.Type != nil && *s.Type == object {
		if s.Properties == nil {
			panic("object without properties")
		}
		var fields []jen.Code
		for name, schema := range *s.Properties {
			fields = append(fields, generateField(schema, name))
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

func generateField(schema Schema, name string) jen.Code {
	var field jen.Code
	fmt.Println(name)
	fmt.Println(schema.Required)
	if schema.Ref != nil {
		var p = generateFieldName(name, schema.Required)
		field = p.Id(refToId(*schema.Ref))
	} else if schema.Type != nil {
		fieldName := generateFieldName(name, schema.Required)
		switch *schema.Type {
		case "string":
			return fieldName.String()
		}
	}
	return field
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
		fmt.Println("is nul")
		return false
	}
	for _, r := range *required {

		if r == name {
			return true
		}
	}
	return false
}
