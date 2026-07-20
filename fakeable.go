package gofakeit

import (
	"reflect"
)

type Fakeable interface {
	Fake(faker *Faker) (any, error)
}

func isFakeable(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

func callFake(faker *Faker, v reflect.Value, possibleKinds ...reflect.Kind) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func containsKind(possibleKinds []reflect.Kind, kind reflect.Kind) bool {
	_ = "STUB: not implemented"
	return false
}
