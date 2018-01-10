package example

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalMyType(MyType)
type MyType struct {
	name        string
	PublicEmail string
}
