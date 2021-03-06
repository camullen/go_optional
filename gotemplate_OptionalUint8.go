// Code generated by gotemplate. DO NOT EDIT.

package go_optional

import "github.com/camullen/go_optional/generic"

// template type Optional(A)

type OptionalUint8 struct {
	value *uint8
}

func NewOptionalUint8(v uint8) OptionalUint8 {
	return OptionalUint8{value: &v}
}

func NoneOptionalUint8() OptionalUint8 {
	return OptionalUint8{value: nil}
}

func (o OptionalUint8) IsNone() bool {
	return o.value == nil
}

func (o OptionalUint8) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalUint8) Expect(msg string) uint8 {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalUint8) Unwrap() uint8 {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalUint8) UnwrapOr(def *uint8) uint8 {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalUint8 func(value *uint8) *uint8

func (o OptionalUint8) Map(mapperFunc MapFuncOptionalUint8) OptionalUint8 {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return OptionalUint8{value: newVal}
}

func (o OptionalUint8) MapInto(mapIntoFunc MapIntoFuncOptionalUint8) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFuncOptionalUint8 func(value *uint8) *interface{}

func (o OptionalUint8) Set(value uint8) OptionalUint8 {
	o.value = &value
	return o
}

func (o *OptionalUint8) SetMutate(value uint8) {
	o.value = &value
}
