package gen

import "github.com/dave/jennifer/jen"

type GStruct struct {
	Id      string
	Fields  []Field
	Methods []Method
}

type Method struct {
	Method    GFunc
	IsPointer bool
	SelfId    string
}

func (gs Method) Statement(structId string) *jen.Statement {
	s := jen.Func()
	s.Op("(")
	s.Id(gs.SelfId)
	if gs.IsPointer {
		s.Op("*")
	}
	s.Id(structId)
	s.Op(")")
	s.Block()
	return s
}

func (gs *GStruct) AddFields(f ...Field) *GStruct {
	gs.Fields = append(gs.Fields, f...)
	return gs
}

func (gs *GStruct) AddMethods(m ...Method) *GStruct {
	gs.Methods = append(gs.Methods, m...)
	return gs
}

func (gs GStruct) Statement() *jen.Statement {
	s := jen.Type().Id(gs.Id)
	fields := make([]jen.Code, len(gs.Fields))
	for _, f := range gs.Fields {
		fields = append(fields, f.Code())
	}
	s.Struct(fields...)
	return s
}

// func (gs GStruct) MethodsStatement() []jen.Statement {
// 	var methods []jen.Code
// 	for _, m := range gs.Methods {
// 		methods = append(methods, m.Method.Statement(gs.Id).Block())
// 	}
// 	return methods
// }

type Field struct {
	Variable
	Tags map[string]string
}

func (f *Field) AddTag(key, value string) *Field {
	if f.Tags == nil {
		f.Tags = make(map[string]string)
	}
	f.Tags[key] = value
	return f
}

func (f Field) Code() *jen.Statement {
	c := jen.Id(f.Id)
	if f.Type.IsPointer {
		c = c.Op("*")
	}
	c.Id(f.Type.Id)
	c.Tag(f.Tags)
	return c
}
