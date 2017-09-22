package crypto

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (BASEURL = "https://api.coinmarketcap.com/v1"
	   DEFAULT_CURRENCY = "USD"
	   DEFAULT_CRYPTO_CURRENCY = "all"
	   DEFAULT_LIMIT = -1
)

type Coinmarketcap struct {
}

func New() *Coinmarketcap {
	return new(Coinmarketcap)
}

func (c Coinmarketcap) GetCoinData(cryptoCurrency string) ([]Coin, error) {
	response, err := makeRequest(makeUrlByParams(cryptoCurrency, DEFAULT_CURRENCY, DEFAULT_LIMIT))
	if err != nil {
		return nil, err
	}

	return processResponse(response, DEFAULT_CURRENCY)
}

func (c Coinmarketcap) GetAllCoinData(currency string, limit int)([]Coin, error) {
	response, err := makeRequest(makeUrlByParams(DEFAULT_CRYPTO_CURRENCY, currency, limit))
	if err != nil {
		return nil, err
	}

	return processResponse(response, currency)
}



func processResponse(response []byte, currency string) ([]Coin, error) {
	data := []map[string]string{}

	err := json.Unmarshal(response, &data)
	if err != nil {
		return nil, err
	}


	coins := make([]Coin, len(data))
	for i := 0; i < len(data); i++ {
		coins[i].Name = data[i]["name"]
		coins[i].Symbol = data[i]["symbol"]
		coins[i].Price, _ = strconv.ParseFloat(data[i][fmt.Sprintf("price_%s", strings.ToLower(currency))], 64)
		coins[i].Currency = currency
	}

	return coins, nil
}


func makeUrlByParams(cryptoCurrency string, currency string, limit int) string {
	var url string = fmt.Sprintf("%s/ticker/", BASEURL)

	if cryptoCurrency != DEFAULT_CRYPTO_CURRENCY {
		url = fmt.Sprintf("%s/%s/", url, cryptoCurrency)
	}

	params := []string{}

	if limit > 0 {
		params = append(params, fmt.Sprintf("limit=%v", limit))
	}

	if currency != DEFAULT_CURRENCY {
		params = append(params, fmt.Sprintf("convert=%s", strings.ToLower(currency)))
	}

	return fmt.Sprintf("%s?%s", url, strings.Join(params, "&"))
}



// HTTP Request Helper
func makeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := doReq(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}



// HTTP Client
func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}