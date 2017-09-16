package crypto

import ("fmt"
		coinApi "github.com/miguelmota/go-coinmarketcap")


type Coinmarketcap struct {
}

func New() *Coinmarketcap {
	return new(Coinmarketcap)
}


func (c Coinmarketcap) GetCoinData(name string) {

	fmt.Println(name)


	// Get global market data
	marketInfo, err := coinApi.GetCoinData(name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(marketInfo)
	}
}
