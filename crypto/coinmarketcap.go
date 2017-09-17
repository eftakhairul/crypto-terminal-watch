package crypto

import ("fmt"
		coinApi "github.com/miguelmota/go-coinmarketcap")


type Coinmarketcap struct {
}

func New() *Coinmarketcap {
	return new(Coinmarketcap)
}


func (c Coinmarketcap) GetCoinData(name string) ([]Coin, error) {

	coins := make([]Coin, 1)

	// Get global market data
	marketInfo, err := coinApi.GetCoinData(name)
	if err != nil {
		fmt.Println(err)
		return coins, err
	} else {

		coins[0].Name = marketInfo.Name
		coins[0].Symbol = marketInfo.Symbol
		coins[0].Price = marketInfo.PriceUsd
		coins[0].Currency = "USD"

		return coins, nil
	}
}
