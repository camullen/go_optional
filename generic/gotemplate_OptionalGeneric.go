// Code generated by gotemplate. DO NOT EDIT.

package generic

// template type Optional(A)

type OptionalGeneric struct {
	value *interface{}
}

func NewOptionalGeneric(v interface{}) OptionalGeneric {
	return OptionalGeneric{value: &v}
}

func NoneOptionalGeneric() OptionalGeneric {
	return OptionalGeneric{value: nil}
}

func (o OptionalGeneric) IsNone() bool {
	return o.value == nil
}

func (o OptionalGeneric) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalGeneric) Expect(msg string) interface{} {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalGeneric) Unwrap() interface{} {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalGeneric) UnwrapOr(def *interface{}) interface{} {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalGeneric func(value *interface{}) *interface{}

func (o OptionalGeneric) Map(mapperFunc MapFuncOptionalGeneric) OptionalGeneric {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return NewOptionalGeneric(newVal)
}

func (o OptionalGeneric) MapInto(mapIntoFunc MapIntoFuncOptionalGeneric) OptionalGeneric {
	if o.IsNone() {
		return NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return NewOptionalGeneric(result)
}

type MapIntoFuncOptionalGeneric func(value *interface{}) *interface{}

func (o OptionalGeneric) Set(value interface{}) OptionalGeneric {
	o.value = &value
	return o
}

func (o *OptionalGeneric) SetMutate(value interface{}) {
	o.value = &value
}