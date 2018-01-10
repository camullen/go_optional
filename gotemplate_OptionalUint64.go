// Code generated by gotemplate. DO NOT EDIT.

package go_optional

import "github.com/camullen/go_optional/generic"

// template type Optional(A)

type OptionalUint64 struct {
	value *uint64
}

func NewOptionalUint64(v uint64) OptionalUint64 {
	return OptionalUint64{value: &v}
}

func NoneOptionalUint64() OptionalUint64 {
	return OptionalUint64{value: nil}
}

func (o OptionalUint64) IsNone() bool {
	return o.value == nil
}

func (o OptionalUint64) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalUint64) Expect(msg string) uint64 {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalUint64) Unwrap() uint64 {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalUint64) UnwrapOr(def *uint64) uint64 {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalUint64 func(value *uint64) *uint64

func (o OptionalUint64) Map(mapperFunc MapFuncOptionalUint64) OptionalUint64 {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return OptionalUint64{value: newVal}
}

func (o OptionalUint64) MapInto(mapIntoFunc MapIntoFuncOptionalUint64) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFuncOptionalUint64 func(value *uint64) *interface{}

func (o OptionalUint64) Set(value uint64) OptionalUint64 {
	o.value = &value
	return o
}

func (o *OptionalUint64) SetMutate(value uint64) {
	o.value = &value
}