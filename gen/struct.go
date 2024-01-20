package gen

import "github.com/dave/jennifer/jen"

type Struct struct {
	Id      string
	Fields  []Field
	Methods []Method
}

type Method struct {
	Func
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
	s.Id(gs.Id).Params()
	var statements []jen.Code
	for _, statement := range gs.Func.Statements {
		statements = append(statements, statement)
	}
	s.Block(statements...)
	return s
}

func (gs *Struct) AddFields(f ...Field) *Struct {
	gs.Fields = append(gs.Fields, f...)
	return gs
}

func (gs *Struct) AddMethods(m ...Method) *Struct {
	gs.Methods = append(gs.Methods, m...)
	return gs
}

func (gs Struct) Statement() []*jen.Statement {
	s := jen.Type().Id(gs.Id)
	fields := make([]jen.Code, len(gs.Fields))
	for _, f := range gs.Fields {
		fields = append(fields, f.Code())
	}
	s.Struct(fields...)
	statements := []*jen.Statement{s}
	for _, m := range gs.Methods {
		statements = append(statements, m.Statement(gs.Id))
	}
	return statements
}

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
