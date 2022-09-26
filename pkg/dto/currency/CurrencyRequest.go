package currency

type Request struct {
	Title string `validate:"required,min=3,max=15" json:"title"`
	IsoCode string `validate:"omitempty" json:"iso_code"`
}