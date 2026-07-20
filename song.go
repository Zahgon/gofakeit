package gofakeit

func SongName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) SongName() string { _ = "STUB: not implemented"; return "" }

func songName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func SongArtist() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) SongArtist() string { _ = "STUB: not implemented"; return "" }

func songArtist(f *Faker) string { _ = "STUB: not implemented"; return "" }

func SongGenre() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) SongGenre() string { _ = "STUB: not implemented"; return "" }

func songGenre(f *Faker) string { _ = "STUB: not implemented"; return "" }

type SongInfo struct {
	Name   string `json:"name" xml:"name"`
	Artist string `json:"artist" xml:"artist"`
	Genre  string `json:"genre" xml:"genre"`
}

func Song() *SongInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Song() *SongInfo { _ = "STUB: not implemented"; return nil }

func song(f *Faker) *SongInfo { _ = "STUB: not implemented"; return nil }

func addSongLookup() { _ = "STUB: not implemented"; return }
