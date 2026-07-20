package gofakeit

import (
	"math/rand/v2"
	"sync"
)

var GlobalFaker *Faker = New(0)

type Faker struct {
	Rand rand.Source

	Locked bool
	mu     sync.Mutex
}

func New(seed uint64) *Faker { _ = "STUB: not implemented"; return nil }

func NewFaker(src rand.Source, lock bool) *Faker { _ = "STUB: not implemented"; return nil }

func (f *Faker) Seed(args ...any) error { _ = "STUB: not implemented"; return nil }

func Seed(args ...any) error { _ = "STUB: not implemented"; return nil }
