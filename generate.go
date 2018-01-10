package go_optional

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalInt(int)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalInt8(int8)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalInt16(int16)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalInt32(int32)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalInt64(int64)


//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalUint(uint)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalUint8(uint8)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalUint16(uint16)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalUint32(uint32)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalUint64(uint64)

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalBool(bool)

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalByte(byte)

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalError(error)

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalString(string)

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalRune(rune)

//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalUintptr(uintptr)


//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalFloat32(float32)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalFloat64(float64)


//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalComplex64(complex64)
//go:generate gotemplate "github.com/camullen/go_optional/template" OptionalComplex128(complex128)

