package gofakeit

type ProductInfo struct {
	Name        string   `json:"name" xml:"name"`
	Description string   `json:"description" xml:"description"`
	Categories  []string `json:"categories" xml:"categories"`
	Price       float64  `json:"price" xml:"price"`
	Features    []string `json:"features" xml:"features"`
	Color       string   `json:"color" xml:"color"`
	Material    string   `json:"material" xml:"material"`
	UPC         string   `json:"upc" xml:"upc"`
	Audience    []string `json:"audience" xml:"audience"`
	Dimension   string   `json:"dimension" xml:"dimension"`
	UseCase     string   `json:"use_case" xml:"use_case"`
	Benefit     string   `json:"benefit" xml:"benefit"`
	Suffix      string   `json:"suffix" xml:"suffix"`
}

func Product() *ProductInfo { _ = "STUB: not implemented"; return nil }

func (f *Faker) Product() *ProductInfo { _ = "STUB: not implemented"; return nil }

func product(f *Faker) *ProductInfo { _ = "STUB: not implemented"; return nil }

func ProductName() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductName() string { _ = "STUB: not implemented"; return "" }

func productName(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductDescription() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductDescription() string { _ = "STUB: not implemented"; return "" }

func productDescription(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductCategory() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductCategory() string { _ = "STUB: not implemented"; return "" }

func productCategory(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductFeature() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductFeature() string { _ = "STUB: not implemented"; return "" }

func productFeature(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductMaterial() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductMaterial() string { _ = "STUB: not implemented"; return "" }

func productMaterial(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductUPC() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductUPC() string { _ = "STUB: not implemented"; return "" }

func productUPC(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductAudience() []string { _ = "STUB: not implemented"; return nil }

func (f *Faker) ProductAudience() []string { _ = "STUB: not implemented"; return nil }

func productAudience(f *Faker) []string { _ = "STUB: not implemented"; return nil }

func ProductDimension() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductDimension() string { _ = "STUB: not implemented"; return "" }

func productDimension(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductUseCase() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductUseCase() string { _ = "STUB: not implemented"; return "" }

func productUseCase(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductBenefit() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductBenefit() string { _ = "STUB: not implemented"; return "" }

func productBenefit(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductSuffix() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductSuffix() string { _ = "STUB: not implemented"; return "" }

func productSuffix(f *Faker) string { _ = "STUB: not implemented"; return "" }

func ProductISBN(opts *ISBNOptions) string { _ = "STUB: not implemented"; return "" }

func (f *Faker) ProductISBN(opts *ISBNOptions) string { _ = "STUB: not implemented"; return "" }

type ISBNOptions struct {
	Version   string
	Separator string
}

func productISBN(f *Faker, opts *ISBNOptions) string { _ = "STUB: not implemented"; return "" }

func addProductLookup() { _ = "STUB: not implemented"; return }
