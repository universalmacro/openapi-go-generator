package gen

import "github.com/dave/jennifer/jen"

type Interface struct {
	Id      string
	Methods []Method
}

func (i *Interface) AddMethods(m ...Method) *Interface {
	i.Methods = append(i.Methods, m...)
	return i
}

func (i Interface) Statement() *jen.Statement {
	methods := make([]jen.Code, len(i.Methods))
	for _, m := range i.Methods {
		method := jen.Id(m.Id).Params(jen.Id("ctx").Op("*").Qual("github.com/gin-gonic/gin", "Context"))
		m.Outputs.Apply(method)
		methods = append(methods, method)
	}
	s := jen.Type().Id(i.Id).Interface(methods...)
	return s
}
