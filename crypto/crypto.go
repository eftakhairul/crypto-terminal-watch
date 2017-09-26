package crypto


import "errors"

//Crypto interface
type Crypto interface {
	GetCoinData() []Coin
	GetCoinDataByCurrency() []Coin
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



