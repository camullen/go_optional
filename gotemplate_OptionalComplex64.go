// Code generated by gotemplate. DO NOT EDIT.

package go_optional

import "github.com/camullen/go_optional/generic"

// template type Optional(A)

type OptionalComplex64 struct {
	value *complex64
}

func NewOptionalComplex64(v complex64) OptionalComplex64 {
	return OptionalComplex64{value: &v}
}

func NoneOptionalComplex64() OptionalComplex64 {
	return OptionalComplex64{value: nil}
}

func (o OptionalComplex64) IsNone() bool {
	return o.value == nil
}

func (o OptionalComplex64) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalComplex64) Expect(msg string) complex64 {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalComplex64) Unwrap() complex64 {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalComplex64) UnwrapOr(def *complex64) complex64 {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalComplex64 func(value *complex64) *complex64

func (o OptionalComplex64) Map(mapperFunc MapFuncOptionalComplex64) OptionalComplex64 {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return OptionalComplex64{value: newVal}
}

func (o OptionalComplex64) MapInto(mapIntoFunc MapIntoFuncOptionalComplex64) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFuncOptionalComplex64 func(value *complex64) *interface{}

func (o OptionalComplex64) Set(value complex64) OptionalComplex64 {
	o.value = &value
	return o
}

func (o *OptionalComplex64) SetMutate(value complex64) {
	o.value = &value
}