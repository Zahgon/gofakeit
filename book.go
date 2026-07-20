package gofakeit

func BookTitle() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BookTitle() string { _ = "STUB: not implemented"; return "" }

func bookTitle(f *Faker) string { _ = "STUB: not implemented"; return "" }

func BookAuthor() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BookAuthor() string { _ = "STUB: not implemented"; return "" }

func bookAuthor(f *Faker) string { _ = "STUB: not implemented"; return "" }

func BookGenre() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BookGenre() string { _ = "STUB: not implemented"; return "" }

func bookGenre(f *Faker) string { _ = "STUB: not implemented"; return "" }

type BookInfo struct {
	Title  string `json:"title" xml:"name"`
	Author string `json:"author" xml:"author"`
	Genre  string `json:"genre" xml:"genre"`
}

func Book() *BookInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Book() *BookInfo { _ = "STUB: not implemented"; return nil }

func book(f *Faker) *BookInfo { _ = "STUB: not implemented"; return nil }

func addBookLookup() { _ = "STUB: not implemented"; return }
