package gofakeit

type PersonInfo struct {
	FirstName  string          `json:"first_name" xml:"first_name"`
	LastName   string          `json:"last_name" xml:"last_name"`
	Gender     string          `json:"gender" xml:"gender"`
	Age        int             `json:"age" xml:"age"`
	SSN        string          `json:"ssn" xml:"ssn"`
	Hobby      string          `json:"hobby" xml:"hobby"`
	Job        *JobInfo        `json:"job" xml:"job"`
	Address    *AddressInfo    `json:"address" xml:"address"`
	Contact    *ContactInfo    `json:"contact" xml:"contact"`
	CreditCard *CreditCardInfo `json:"credit_card" xml:"credit_card"`
}

func Person() *PersonInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Person() *PersonInfo { _ = "STUB: not implemented"; return nil }

func person(f *Faker) *PersonInfo { _ = "STUB: not implemented"; return nil }

func Name() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Name() string { _ = "STUB: not implemented"; return "" }

func name(f *Faker) string { _ = "STUB: not implemented"; return "" }

func FirstName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) FirstName() string { _ = "STUB: not implemented"; return "" }

func firstName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func MiddleName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) MiddleName() string { _ = "STUB: not implemented"; return "" }

func middleName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func LastName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) LastName() string { _ = "STUB: not implemented"; return "" }

func lastName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func NamePrefix() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) NamePrefix() string { _ = "STUB: not implemented"; return "" }

func namePrefix(f *Faker) string { _ = "STUB: not implemented"; return "" }

func NameSuffix() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) NameSuffix() string { _ = "STUB: not implemented"; return "" }

func nameSuffix(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Age() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Age() int { _ = "STUB: not implemented"; return 0 }

func age(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func Ethnicity() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Ethnicity() string { _ = "STUB: not implemented"; return "" }

func ethnicity(f *Faker) string { _ = "STUB: not implemented"; return "" }

func SSN() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) SSN() string { _ = "STUB: not implemented"; return "" }

func ssn(f *Faker) string { _ = "STUB: not implemented"; return "" }

func EIN() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) EIN() string { _ = "STUB: not implemented"; return "" }

func ein(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Gender() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Gender() string { _ = "STUB: not implemented"; return "" }

func gender(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Hobby() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Hobby() string { _ = "STUB: not implemented"; return "" }

func hobby(f *Faker) string { _ = "STUB: not implemented"; return "" }

func SocialMedia() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) SocialMedia() string { _ = "STUB: not implemented"; return "" }

func socialMedia(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Bio() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Bio() string { _ = "STUB: not implemented"; return "" }

func bio(f *Faker) string { _ = "STUB: not implemented"; return "" }

type ContactInfo struct {
	Phone string `json:"phone" xml:"phone"`
	Email string `json:"email" xml:"email"`
}

func Contact() *ContactInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Contact() *ContactInfo { _ = "STUB: not implemented"; return nil }

func contact(f *Faker) *ContactInfo { _ = "STUB: not implemented"; return nil }

func Phone() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Phone() string { _ = "STUB: not implemented"; return "" }

func phone(f *Faker) string { _ = "STUB: not implemented"; return "" }

func PhoneFormatted() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) PhoneFormatted() string { _ = "STUB: not implemented"; return "" }

func phoneFormatted(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Email() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) Email() string { _ = "STUB: not implemented"; return "" }

func email(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Teams(peopleArray []string, teamsArray []string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (f *Faker) Teams(peopleArray []string, teamsArray []string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func teams(f *Faker, people []string, teams []string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func addPersonLookup() { _ = "STUB: not implemented"; return }
