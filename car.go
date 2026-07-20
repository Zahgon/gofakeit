package gofakeit

type CarInfo struct {
	Type         string `json:"type" xml:"type"`
	Fuel         string `json:"fuel" xml:"fuel"`
	Transmission string `json:"transmission" xml:"transmission"`
	Brand        string `json:"brand" xml:"brand"`
	Model        string `json:"model" xml:"model"`
	Year         int    `json:"year" xml:"year"`
}

func Car() *CarInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Car() *CarInfo { _ = "STUB: not implemented"; return nil }

func car(f *Faker) *CarInfo { _ = "STUB: not implemented"; return nil }

func CarType() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CarType() string { _ = "STUB: not implemented"; return "" }

func carType(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CarFuelType() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CarFuelType() string { _ = "STUB: not implemented"; return "" }

func carFuelType(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CarTransmissionType() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CarTransmissionType() string { _ = "STUB: not implemented"; return "" }

func carTransmissionType(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CarMaker() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CarMaker() string { _ = "STUB: not implemented"; return "" }

func carMaker(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CarModel() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CarModel() string { _ = "STUB: not implemented"; return "" }

func carModel(f *Faker) string { _ = "STUB: not implemented"; return "" }

func addCarLookup() { _ = "STUB: not implemented"; return }
