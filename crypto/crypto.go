package crypto


import "errors"

//Crypto interface
type Crypto interface {
	GetCoinData() []Coin
	GetCoinDataByCurrency() []Coin
}

var cryptoCurrencies = map[string]string = map[string]string{
	"BTC": "bitcoin",
	"XBT": "bitcoin",
	"ETH": "ethereum",
	"BCH": "bitcoin-cash",
	"XBC": "bitcoin-cash",
	"XRP": "ripple",
	"DASH": "dash",
	"LTC": "litecoin",
}

//Coin struct
type Coin struct {
	Name   string
	Symbol string
	Price  float64
	Currency string
}


type CryptoApi struct {
	cryptoApi Crypto
}



func validateCryptoCurrency(name string) (string, error) {
	if value, ok := cryptoCurrencies[name]; ok {
		return value, nil
	}

	return nil, errors.New("can't work with 42")
}
