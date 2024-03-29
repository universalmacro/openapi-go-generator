package openapi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/universalmacro/openapi-go-generator/gen"
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

func (o Openapi) File(model string) *jen.File {
	model = strings.ToLower(model)
	re := regexp.MustCompile("[^a-z0-9]")
	model = re.ReplaceAllString(model, "")
	f := jen.NewFile(model)
	f.ImportName("github.com/gin-gonic/gin", "gin")
	apis := make(map[string]gen.Interface, 0)
	bindingFuncs := map[string]gen.Func{}
	for uri, path := range o.Paths {
		for method, httpMethod := range path {
			if httpMethod.OperationId == nil {
				panic(method + " " + uri + "no operationId")
			}
			if len(httpMethod.Tags[0]) == 0 {
				fmt.Println("warning: " + method + " " + uri + "no tag")
				continue
			}
			tag := httpMethod.Tags[0]
			if _, ok := bindingFuncs[tag]; !ok {
				params := gen.Parameters{gen.Variable{Id: "router", Type: gen.Type{
					Id:        "gin.Engine",
					IsPointer: true,
				}}, gen.Variable{Id: "api", Type: gen.Type{Id: tag + "Api"}}}
				bindingFuncs[tag] = gen.Func{Id: tag + "ApiBinding", Parameters: params}
			}
			bindingFunc := bindingFuncs[tag]
			bindingFunc.AddStatements(
				jen.
					Id("router").
					Dot(strings.ToUpper(method)).
					Params(jen.Lit(pathConvert(uri)), jen.Id("api").Dot(*httpMethod.OperationId)))
			bindingFuncs[tag] = bindingFunc
			if _, ok := apis[tag]; !ok {
				apiInterface := gen.Interface{Id: tag + "Api"}
				apis[tag] = apiInterface
			}
			apiInterface := apis[tag]
			apiInterface.AddMethods(gen.Method{Func: gen.Func{Id: *httpMethod.OperationId}})
			apis[tag] = apiInterface
		}
	}
	for _, bindingFunc := range bindingFuncs {
		f.Add(bindingFunc.Statement())
	}
	for _, api := range apis {
		f.Add(api.Statement())
	}
	if o.Components != nil {
		if o.Components.Schemas != nil {
			for id, schema := range o.Components.Schemas {
				if schema.Statement(id) != nil {
					for _, st := range schema.Statement(id) {
						f.Add(st)
					}
				}
				// Enum
				if schema.Type != nil && *schema.Type == "string" && schema.Enum != nil {
					en := gen.StringEnum{Id: id, Values: *schema.Enum}.Statement()
					for _, e := range en {
						f.Add(e)
					}
				}
			}
		}
	}
	return f
}

type Tag struct {
	Name string `yaml:"name"`
}

type Components struct {
	Schemas    map[string]Schema    `yaml:"schemas"`
	Parameters map[string]Parameter `yaml:"parameters"`
}

type Schema struct {
	Type          *string            `yaml:"type"`
	Ref           *string            `yaml:"$ref"`
	Properties    *map[string]Schema `yaml:"properties"`
	Format        *string            `yaml:"format"`
	Default       *string            `yaml:"default"`
	Required      *[]string          `yaml:"required"`
	Enum          *[]string          `yaml:"enum"`
	Items         *Schema            `yaml:"items"`
	OneOf         *[]Schema          `yaml:"oneOf"`
	AllOf         *[]Schema          `yaml:"allOf"`
	AnyOf         *[]Schema          `yaml:"anyOf"`
	Discriminator *Discriminator     `yaml:"discriminator"`
}

type Discriminator struct {
	PropertyName string `yaml:"propertyName"`
}

func (s Schema) Statement(id string) []*jen.Statement {
	if s.Type != nil && *s.Type == object {
		if s.Properties == nil {
			panic("object without properties")
		}
		st := gen.Struct{Id: id}
		for fieldId, schema := range *s.Properties {
			tags := map[string]string{"json": fieldId, "xml": fieldId}
			field := gen.Field{
				Variable: gen.Variable{
					Id: capitalize(fieldId),
					Type: gen.Type{
						Id:        generateTypeString(schema),
						IsPointer: !isRequired(fieldId, s.Required),
					},
				},
				Tags: tags,
			}
			st.AddFields(field)
		}
		return st.Statement()
	}
	return nil
}

type Parameter struct {
	Name        *string `yaml:"name"`
	In          *string `yaml:"in"`
	Description *string `yaml:"description"`
	Schema      *Schema `yaml:"schema"`
}
