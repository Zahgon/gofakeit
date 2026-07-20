package gofakeit

const (
	idLength       = 20
	idBitsPerChar  = 5
	idAlphabetMask = (1 << idBitsPerChar) - 1

	idAlphabetStr = "23456789abcdefgghjkmnpqrstuvwxyz"
	hexDigits     = "0123456789abcdef"
)

var (
	idAlphabet = []byte(idAlphabetStr)
)

func ID() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ID() string { _ = "STUB: not implemented"; return "" }

func id(f *Faker) string { _ = "STUB: not implemented"; return "" }

func UUID() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) UUID() string { _ = "STUB: not implemented"; return "" }

func uuid(f *Faker) string { _ = "STUB: not implemented"; return "" }

func encodeHexLower(dst, src []byte) { _ = "STUB: not implemented"; return }

func addIDLookup() { _ = "STUB: not implemented"; return }
