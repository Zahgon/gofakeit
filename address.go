package gofakeit

type AddressInfo struct {
	Address   string  `json:"address" xml:"address"`
	Street    string  `json:"street" xml:"street"`
	Unit      string  `json:"unit" xml:"unit"`
	City      string  `json:"city" xml:"city"`
	State     string  `json:"state" xml:"state"`
	Zip       string  `json:"zip" xml:"zip"`
	Country   string  `json:"country" xml:"country"`
	Latitude  float64 `json:"latitude" xml:"latitude"`
	Longitude float64 `json:"longitude" xml:"longitude"`
}

func Address() *AddressInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Address() *AddressInfo { _ = "STUB: not implemented"; return nil }

func address(f *Faker) *AddressInfo { _ = "STUB: not implemented"; return nil }

func Street() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Street() string { _ = "STUB: not implemented"; return "" }

func street(f *Faker) string { _ = "STUB: not implemented"; return "" }

func StreetNumber() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) StreetNumber() string { _ = "STUB: not implemented"; return "" }

func streetNumber(f *Faker) string { _ = "STUB: not implemented"; return "" }

func StreetPrefix() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) StreetPrefix() string { _ = "STUB: not implemented"; return "" }

func streetPrefix(f *Faker) string { _ = "STUB: not implemented"; return "" }

func StreetName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) StreetName() string { _ = "STUB: not implemented"; return "" }

func streetName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func StreetSuffix() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) StreetSuffix() string { _ = "STUB: not implemented"; return "" }

func streetSuffix(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Unit() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Unit() string { _ = "STUB: not implemented"; return "" }

func unit(f *Faker) string { _ = "STUB: not implemented"; return "" }

func City() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) City() string { _ = "STUB: not implemented"; return "" }

func city(f *Faker) string { _ = "STUB: not implemented"; return "" }

func State() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) State() string { _ = "STUB: not implemented"; return "" }

func state(f *Faker) string { _ = "STUB: not implemented"; return "" }

func StateAbr() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) StateAbr() string { _ = "STUB: not implemented"; return "" }

func stateAbr(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Zip() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Zip() string { _ = "STUB: not implemented"; return "" }

func zip(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Country() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Country() string { _ = "STUB: not implemented"; return "" }

func country(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CountryAbr() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CountryAbr() string { _ = "STUB: not implemented"; return "" }

func countryAbr(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Latitude() float64 { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Latitude() float64 { _ = "STUB: not implemented"; return 0 }

func latitude(f *Faker) float64 { _ = "STUB: not implemented"; return 0 }

func LatitudeInRange(min, max float64) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *Faker) LatitudeInRange(min, max float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func latitudeInRange(f *Faker, min, max float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func Longitude() float64 { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Longitude() float64 { _ = "STUB: not implemented"; return 0 }

func longitude(f *Faker) float64 { _ = "STUB: not implemented"; return 0 }

func LongitudeInRange(min, max float64) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *Faker) LongitudeInRange(min, max float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func longitudeInRange(f *Faker, min, max float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func addAddressLookup() { _ = "STUB: not implemented"; return }
