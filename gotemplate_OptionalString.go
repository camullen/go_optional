// Code generated by gotemplate. DO NOT EDIT.

package go_optional

import "github.com/camullen/go_optional/generic"

// template type Optional(A)

type OptionalString struct {
	value *string
}

func NewOptionalString(v string) OptionalString {
	return OptionalString{value: &v}
}

func NoneOptionalString() OptionalString {
	return OptionalString{value: nil}
}

func (o OptionalString) IsNone() bool {
	return o.value == nil
}

func (o OptionalString) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalString) Expect(msg string) string {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalString) Unwrap() string {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalString) UnwrapOr(def *string) string {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalString func(value *string) *string

func (o OptionalString) Map(mapperFunc MapFuncOptionalString) OptionalString {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return OptionalString{value: newVal}
}

func (o OptionalString) MapInto(mapIntoFunc MapIntoFuncOptionalString) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFuncOptionalString func(value *string) *interface{}

func (o OptionalString) Set(value string) OptionalString {
	o.value = &value
	return o
}

func (o *OptionalString) SetMutate(value string) {
	o.value = &value
}
