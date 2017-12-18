// Crypto package with common coin interface for all crypto related API
package crypto

const (
	DEFAULT_CURRENCY           = "USD"
	DEFAULT_CRYPTO_CURRENCY_ID = "all"
	DEFAULT_LIMIT              = -1
)

//Coin struct
type Coin struct {
	Name      string
	Symbol    string
	Price     float64
	Currency  string
	Volume24  float64
	Maxsupply float64
}

//Crypto interface
type Crypto interface {
	GetCoinData(cryptoCurrency string, currency string) ([]Coin, error)
	GetCoinDataByCurrency(currency string, limit int) ([]Coin, error)
}
