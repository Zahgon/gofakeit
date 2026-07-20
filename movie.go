package gofakeit

func MovieName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) MovieName() string { _ = "STUB: not implemented"; return "" }

func movieName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func MovieGenre() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) MovieGenre() string { _ = "STUB: not implemented"; return "" }

func movieGenre(f *Faker) string { _ = "STUB: not implemented"; return "" }

type MovieInfo struct {
	Name  string `json:"name" xml:"name"`
	Genre string `json:"genre" xml:"genre"`
}

func Movie() *MovieInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Movie() *MovieInfo { _ = "STUB: not implemented"; return nil }

func movie(f *Faker) *MovieInfo { _ = "STUB: not implemented"; return nil }

func addMovieLookup() { _ = "STUB: not implemented"; return }
