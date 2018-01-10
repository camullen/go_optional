// Code generated by gotemplate. DO NOT EDIT.

package go_optional

import "github.com/camullen/go_optional/generic"

// template type Optional(A)

type OptionalInt struct {
	value *int
}

func NewOptionalInt(v int) OptionalInt {
	return OptionalInt{value: &v}
}

func NoneOptionalInt() OptionalInt {
	return OptionalInt{value: nil}
}

func (o OptionalInt) IsNone() bool {
	return o.value == nil
}

func (o OptionalInt) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalInt) Expect(msg string) int {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalInt) Unwrap() int {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalInt) UnwrapOr(def *int) int {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalInt func(value *int) *int

func (o OptionalInt) Map(mapperFunc MapFuncOptionalInt) OptionalInt {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return OptionalInt{value: newVal}
}

func (o OptionalInt) MapInto(mapIntoFunc MapIntoFuncOptionalInt) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFuncOptionalInt func(value *int) *interface{}

func (o OptionalInt) Set(value int) OptionalInt {
	o.value = &value
	return o
}

func (o *OptionalInt) SetMutate(value int) {
	o.value = &value
}
