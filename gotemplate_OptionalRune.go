// Code generated by gotemplate. DO NOT EDIT.

package go_optional

import "github.com/camullen/go_optional/generic"

// template type Optional(A)

type OptionalRune struct {
	value *rune
}

func NewOptionalRune(v rune) OptionalRune {
	return OptionalRune{value: &v}
}

func NoneOptionalRune() OptionalRune {
	return OptionalRune{value: nil}
}

func (o OptionalRune) IsNone() bool {
	return o.value == nil
}

func (o OptionalRune) IsSome() bool {
	return !o.IsNone()
}

func (o OptionalRune) Expect(msg string) rune {
	if o.IsNone() {
		panic(msg)
	}
	return *o.value
}

func (o OptionalRune) Unwrap() rune {
	return o.Expect("Optional unwrapped with nil value")
}

func (o OptionalRune) UnwrapOr(def *rune) rune {
	if o.IsNone() {
		return *def
	}
	return *o.value
}

type MapFuncOptionalRune func(value *rune) *rune

func (o OptionalRune) Map(mapperFunc MapFuncOptionalRune) OptionalRune {
	if o.IsNone() {
		return o
	}

	newVal := mapperFunc(o.value)
	return OptionalRune{value: newVal}
}

func (o OptionalRune) MapInto(mapIntoFunc MapIntoFuncOptionalRune) generic.OptionalGeneric {
	if o.IsNone() {
		return generic.NoneOptionalGeneric()
	}
	result := mapIntoFunc(o.value)
	return generic.NewOptionalGeneric(result)
}

type MapIntoFuncOptionalRune func(value *rune) *interface{}

func (o OptionalRune) Set(value rune) OptionalRune {
	o.value = &value
	return o
}

func (o *OptionalRune) SetMutate(value rune) {
	o.value = &value
}
