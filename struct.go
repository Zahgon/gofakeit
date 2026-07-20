package gofakeit

import (
	"reflect"
)

var RecursiveDepth = 10

func Struct(v any) error { _ = "STUB: not implemented"; return nil }

func (f *Faker) Struct(v any) error { _ = "STUB: not implemented"; return nil }

func structFunc(f *Faker, v any) error { _ = "STUB: not implemented"; return nil }

func r(f *Faker, t reflect.Type, v reflect.Value, tag string, size int, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func rCustom(f *Faker, v reflect.Value, tag string) error { _ = "STUB: not implemented"; return nil }

func rStruct(f *Faker, t reflect.Type, v reflect.Value, tag string, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func rPointer(f *Faker, t reflect.Type, v reflect.Value, tag string, size int, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func rSlice(f *Faker, t reflect.Type, v reflect.Value, tag string, size int, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func rMap(f *Faker, t reflect.Type, v reflect.Value, tag string, size int, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func rString(f *Faker, t reflect.Type, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func rInt(f *Faker, t reflect.Type, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func rUint(f *Faker, t reflect.Type, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func rFloat(f *Faker, t reflect.Type, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func rBool(f *Faker, t reflect.Type, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func rTime(f *Faker, t reflect.StructField, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}
