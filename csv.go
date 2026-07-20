package gofakeit

type CSVOptions struct {
	Delimiter string  `json:"delimiter" xml:"delimiter" fake:"{randomstring:[,,tab]}"`
	RowCount  int     `json:"row_count" xml:"row_count" fake:"{number:1,10}"`
	Fields    []Field `json:"fields" xml:"fields" fake:"{fields}"`
}

func CSV(co *CSVOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Faker) CSV(co *CSVOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func csvFunc(f *Faker, co *CSVOptions) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func addFileCSVLookup() { _ = "STUB: not implemented"; return }
