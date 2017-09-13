package crypto

//Crypto interface
type Crypto interface {
	GetCoinData() []Coin
	GetCoinDataByCurrency() []Coin
}


//Coin struct
type Coin struct {
	name   string
	symbol string
	price  float64
	currency string
}


type CryptoApi struct {
	cryptoApi Crypto
}

func NewCrypto(c Crypto) Crypto {
	var cryptoApi Crypto =  CryptoApi{cryptoApi: c}
	return cryptoApi;
}




