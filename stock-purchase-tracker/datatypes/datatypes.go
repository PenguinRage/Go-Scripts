package datatypes

type Stock struct {
	Exchange string  `db:"exchange"`
	Code     string  `db:"code"`
	State    string  `db:"state"`
	Quantity int     `db:"quantity"`
	Currency string  `db:"currency"`
	Price    float64 `db:"price"`
}
