package gofakeit

import (
	"sync"
)

var FuncLookups map[string]Info
var lockFuncLookups sync.RWMutex

type MapParams map[string]MapParamsValue

type MapParamsValue []string

type Info struct {
	Display     string                                                `json:"display"`
	Category    string                                                `json:"category"`
	Description string                                                `json:"description"`
	Example     string                                                `json:"example"`
	Output      string                                                `json:"output"`
	Aliases     []string                                              `json:"aliases"`
	Keywords    []string                                              `json:"keywords"`
	ContentType string                                                `json:"content_type"`
	Params      []Param                                               `json:"params"`
	Any         any                                                   `json:"any"`
	Generate    func(f *Faker, m *MapParams, info *Info) (any, error) `json:"-"`
}

type Param struct {
	Field       string   `json:"field"`
	Display     string   `json:"display"`
	Type        string   `json:"type"`
	Optional    bool     `json:"optional"`
	Default     string   `json:"default"`
	Options     []string `json:"options"`
	Description string   `json:"description"`
}

type Field struct {
	Name     string    `json:"name"`
	Function string    `json:"function"`
	Params   MapParams `json:"params"`
}

func init() { initLookup() }

func initLookup() { _ = "STUB: not implemented"; return }

var internalFuncLookups map[string]Info = map[string]Info{
	"fields": {
		Description: "Example fields for generating csv, json, xml, etc",
		Output:      "gofakeit.Field",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			function, _ := GetRandomSimpleFunc(f)
			return Field{
				Name:     function,
				Function: function,
			}, nil
		},
	},
}

func NewMapParams() *MapParams { _ = "STUB: not implemented"; return nil }

func (m *MapParams) Add(field string, value string) { _ = "STUB: not implemented"; return }

func (m *MapParams) Get(field string) []string { _ = "STUB: not implemented"; return nil }

func (m *MapParams) Size() int { _ = "STUB: not implemented"; return 0 }

func (m *MapParamsValue) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func GetRandomSimpleFunc(f *Faker) (string, Info) { _ = "STUB: not implemented"; return "", *new(Info) }

func AddFuncLookup(functionName string, info Info) { _ = "STUB: not implemented"; return }

func GetFuncLookup(functionName string) *Info { _ = "STUB: not implemented"; return nil }

func RemoveFuncLookup(functionName string) { _ = "STUB: not implemented"; return }

func (i *Info) GetAny(m *MapParams, field string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (i *Info) GetMap(m *MapParams, field string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Info) GetField(m *MapParams, field string) (*Param, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (i *Info) GetBool(m *MapParams, field string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (i *Info) GetInt(m *MapParams, field string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Info) GetUint(m *MapParams, field string) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Info) GetFloat32(m *MapParams, field string) (float32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Info) GetFloat64(m *MapParams, field string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (i *Info) GetString(m *MapParams, field string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (i *Info) GetStringArray(m *MapParams, field string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Info) GetIntArray(m *MapParams, field string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Info) GetUintArray(m *MapParams, field string) ([]uint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *Info) GetFloat32Array(m *MapParams, field string) ([]float32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
