package gofakeit

import (
	"encoding/xml"
)

type XMLOptions struct {
	Type          string  `json:"type" xml:"type" fake:"{randomstring:[array,single]}"`
	RootElement   string  `json:"root_element" xml:"root_element"`
	RecordElement string  `json:"record_element" xml:"record_element"`
	RowCount      int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Indent        bool    `json:"indent" xml:"indent"`
	Fields        []Field `json:"fields" xml:"fields" fake:"{fields}"`
}

type xmlArray struct {
	XMLName xml.Name
	Array   []xmlMap
}

type xmlMap struct {
	XMLName  xml.Name
	KeyOrder []string
	Map      map[string]any `xml:",chardata"`
}

type xmlEntry struct {
	XMLName xml.Name
	Value   any `xml:",chardata"`
}

func (m xmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func xmlMapLoop(e *xml.Encoder, m *xmlMap) error { _ = "STUB: not implemented"; return nil }

func XML(xo *XMLOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Faker) XML(xo *XMLOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func xmlFunc(f *Faker, xo *XMLOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func addFileXMLLookup() { _ = "STUB: not implemented"; return }
