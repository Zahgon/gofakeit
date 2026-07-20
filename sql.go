package gofakeit

type SQLOptions struct {
	Table  string  `json:"table" xml:"table"`
	Count  int     `json:"count" xml:"count"`
	Fields []Field `json:"fields" xml:"fields"`
}

func SQL(so *SQLOptions) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *Faker) SQL(so *SQLOptions) (string, error) { _ = "STUB: not implemented"; return "", nil }

func sqlFunc(f *Faker, so *SQLOptions) (string, error) { _ = "STUB: not implemented"; return "", nil }

func sqlConvertType(t string, val any) string { _ = "STUB: not implemented"; return "" }

func addDatabaseSQLLookup() { _ = "STUB: not implemented"; return }
