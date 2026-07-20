package gofakeit

import (
	"reflect"
)

type JSONOptions struct {
	Type     string  `json:"type" xml:"type" fake:"{randomstring:[array,object]}"`
	RowCount int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Indent   bool    `json:"indent" xml:"indent"`
	Fields   []Field `json:"fields" xml:"fields" fake:"{fields}"`
}

type jsonKeyVal struct {
	Key   string
	Value any
}

type jsonOrderedKeyVal []*jsonKeyVal

func (okv jsonOrderedKeyVal) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func JSON(jo *JSONOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Faker) JSON(jo *JSONOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func jsonFunc(f *Faker, jo *JSONOptions) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addFileJSONLookup() { _ = "STUB: not implemented"; return }

func rJsonRawMessage(f *Faker, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

func rJsonNumber(f *Faker, v reflect.Value, tag string) error {
	_ = "STUB: not implemented"
	return nil
}
