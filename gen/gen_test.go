package gen

import (
	"fmt"
	"testing"

	"github.com/dave/jennifer/jen"
)

func TestFunc(t *testing.T) {
	fun := GFunc{Id: "Test"}
	fun.AddParameters(Variable{Name: "a", Type: Type{Name: "string"}})
	fun.AddOutputs(Variable{Name: "b", Type: Type{Name: "int", IsPointer: true}})
	f := jen.NewFile("models")
	f.Add(fun.Statement())
	fmt.Printf("%#v", f)
}
