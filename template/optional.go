package template

import "github.com/camullen/go_optional/generic"

// template type Optional(A)
type A int

type Optional struct {
	value *A
}

func NewOptional(v A) Optional {
	return Optional{value: &v}
}

func NoneOptional() Optional {
	return Optional{value: nil}
}

func (o Optional) IsNone() bool {
	return o.value == nil
}

func (o Optional) IsSome() bool {
	return !o.IsNone()
}

func (o Optional) Expect(msg string) A {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o Optional) Unwrap() A {
	return o.Expect("Optional unwrapped with nil value")
}

func (o Optional) UnwrapOr(def *A) A {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFunc func(value *A) *A

func (o Optional) Map(mapperFunc MapFunc) Optional {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return Optional{value: newVal}
}

func (o Optional) MapInto(mapIntoFunc MapIntoFunc) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFunc func(value *A) *interface{}

func (o Optional) Set(value A) Optional {
	o.value = &value
	return o
}

func (o *Optional) SetMutate(value A) {
	o.value = &value
}
