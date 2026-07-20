package gofakeit

type CurrencyInfo struct {
	Short string `json:"short" xml:"short"`
	Long  string `json:"long" xml:"long"`
}

func Currency() *CurrencyInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Currency() *CurrencyInfo { _ = "STUB: not implemented"; return nil }

func currency(f *Faker) *CurrencyInfo { _ = "STUB: not implemented"; return nil }

func CurrencyShort() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CurrencyShort() string { _ = "STUB: not implemented"; return "" }

func currencyShort(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CurrencyLong() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CurrencyLong() string { _ = "STUB: not implemented"; return "" }

func currencyLong(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Price(min, max float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Price(min, max float64) float64 { _ = "STUB: not implemented"; return 0 }

func price(f *Faker, min, max float64) float64 { _ = "STUB: not implemented"; return 0 }

type CreditCardInfo struct {
	Type   string `json:"type" xml:"type"`
	Number string `json:"number" xml:"number"`
	Exp    string `json:"exp" xml:"exp"`
	Cvv    string `json:"cvv" xml:"cvv"`
}

func CreditCard() *CreditCardInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) CreditCard() *CreditCardInfo { _ = "STUB: not implemented"; return nil }

func creditCard(f *Faker) *CreditCardInfo { _ = "STUB: not implemented"; return nil }

func CreditCardType() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CreditCardType() string { _ = "STUB: not implemented"; return "" }

func creditCardType(f *Faker) string { _ = "STUB: not implemented"; return "" }

type CreditCardOptions struct {
	Types []string `json:"types"`
	Bins  []string `json:"bins"`
	Gaps  bool     `json:"gaps"`
}

func CreditCardNumber(cco *CreditCardOptions) string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CreditCardNumber(cco *CreditCardOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func creditCardNumber(f *Faker, cco *CreditCardOptions) string {
	_ = "STUB: not implemented"
	return ""
}

func CreditCardExp() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CreditCardExp() string { _ = "STUB: not implemented"; return "" }

func creditCardExp(f *Faker) string { _ = "STUB: not implemented"; return "" }

func CreditCardCvv() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) CreditCardCvv() string { _ = "STUB: not implemented"; return "" }

func creditCardCvv(f *Faker) string { _ = "STUB: not implemented"; return "" }

func isLuhn(s string) bool { _ = "STUB: not implemented"; return false }

func AchRouting() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) AchRouting() string { _ = "STUB: not implemented"; return "" }

func achRouting(f *Faker) string { _ = "STUB: not implemented"; return "" }

func AchAccount() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) AchAccount() string { _ = "STUB: not implemented"; return "" }

func achAccount(f *Faker) string { _ = "STUB: not implemented"; return "" }

func BitcoinAddress() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BitcoinAddress() string { _ = "STUB: not implemented"; return "" }

func bitcoinAddress(f *Faker) string { _ = "STUB: not implemented"; return "" }

func BitcoinPrivateKey() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BitcoinPrivateKey() string { _ = "STUB: not implemented"; return "" }

func bitcoinPrivateKey(f *Faker) string { _ = "STUB: not implemented"; return "" }

func BankName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BankName() string { _ = "STUB: not implemented"; return "" }

func bankName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func BankType() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) BankType() string { _ = "STUB: not implemented"; return "" }

func bankType(f *Faker) string { _ = "STUB: not implemented"; return "" }

func addPaymentLookup() { _ = "STUB: not implemented"; return }
