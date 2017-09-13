package crypto


import (coinApi "github.com/miguelmota/go-coinmarketcap")


type Coinmarketcap struct {
}

func New() *Coinmarketcap {
	return new(Coinmarketcap)
}


func (c Coinmarketcap) GetCoinData() []Coin {
}

func (c Coinmarketcap) GetCoinDataByCurrency(cryptoCurrency string) Coin {

}