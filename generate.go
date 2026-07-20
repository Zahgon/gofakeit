package gofakeit

import (
	"regexp/syntax"
)

func Generate(dataVal string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *Faker) Generate(dataVal string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func generate(f *Faker, dataVal string) (string, error) { _ = "STUB: not implemented"; return "", nil }

type FixedWidthOptions struct {
	RowCount int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Fields   []Field `json:"fields" xml:"fields" fake:"{fields}"`
}

func FixedWidth(co *FixedWidthOptions) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *Faker) FixedWidth(co *FixedWidthOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func fixeWidthFunc(f *Faker, co *FixedWidthOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func Regex(regexStr string) string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Regex(regexStr string) string { _ = "STUB: not implemented"; return "" }

func regex(f *Faker, regexStr string) (gen string) { _ = "STUB: not implemented"; return "" }

func regexGenerate(f *Faker, re *syntax.Regexp, limit int) string {
	_ = "STUB: not implemented"
	return ""
}

func Map() map[string]any { _ = "STUB: not implemented"; return nil }

func (f *Faker) Map() map[string]any { _ = "STUB: not implemented"; return nil }

func mapFunc(f *Faker) map[string]any { _ = "STUB: not implemented"; return nil }

func addGenerateLookup() { _ = "STUB: not implemented"; return }
