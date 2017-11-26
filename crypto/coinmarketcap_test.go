// Unit test for CoinMarketCap API
package crypto

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eftakhairul/crypto-terminal-watch/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetCoinDataForCoinMarketCap(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{
					"id": "bitcoin", 
					"name": "Bitcoin", 
					"symbol": "BTC", 
					"rank": "1", 
					"price_usd": "7829.1", 
					"price_btc": "1.0", 
					"24h_volume_usd": "3948610000.0", 
					"market_cap_usd": "130683039694", 
					"available_supply": "16691962.0", 
					"total_supply": "16691962.0", 
					"max_supply": "21000000.0", 
					"percent_change_1h": "-2.71", 
					"percent_change_24h": "-1.95", 
					"percent_change_7d": "18.48", 
					"last_updated": "1511236753"
			}
	]`)
	}))
	defer testServer.Close()

	coinmarketcap := crypto.NewCoinmarketcap()
	coinmarketcap.SetBaseURL(testServer.URL)
	coins, err := coinmarketcap.GetCoinData("BTC", "USD")

	assert.NoError(t, err)

	assert.Equal(t, 7829.1, coins[0].Price)
	assert.Equal(t, "Bitcoin", coins[0].Name)
	assert.Equal(t, "BTC", coins[0].Symbol)
}

func TestGetAllCoinDataForCoinMarketCap(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[
			{
					"id": "bitcoin", 
					"name": "Bitcoin", 
					"symbol": "BTC", 
					"rank": "1", 
					"price_usd": "9440.11", 
					"price_btc": "1.0", 
					"24h_volume_usd": "4948740000.0", 
					"market_cap_usd": "157676382589", 
					"available_supply": "16702812.0", 
					"total_supply": "16702812.0", 
					"max_supply": "21000000.0", 
					"percent_change_1h": "1.07", 
					"percent_change_24h": "8.79", 
					"percent_change_7d": "19.17", 
					"last_updated": "1511721553"
			}, 
			{
					"id": "ethereum", 
					"name": "Ethereum", 
					"symbol": "ETH", 
					"rank": "2", 
					"price_usd": "457.497", 
					"price_btc": "0.0485508", 
					"24h_volume_usd": "1111250000.0", 
					"market_cap_usd": "43905991693.0", 
					"available_supply": "95970010.0", 
					"total_supply": "95970010.0", 
					"max_supply": null, 
					"percent_change_1h": "-0.42", 
					"percent_change_24h": "-2.4", 
					"percent_change_7d": "28.76", 
					"last_updated": "1511721553"
			}, 
			{
					"id": "bitcoin-cash", 
					"name": "Bitcoin Cash", 
					"symbol": "BCH", 
					"rank": "3", 
					"price_usd": "1603.7", 
					"price_btc": "0.170189", 
					"24h_volume_usd": "1026510000.0", 
					"market_cap_usd": "26979887043.0", 
					"available_supply": "16823525.0", 
					"total_supply": "16823525.0", 
					"max_supply": "21000000.0", 
					"percent_change_1h": "0.26", 
					"percent_change_24h": "-0.42", 
					"percent_change_7d": "35.86", 
					"last_updated": "1511721570"
			}, 
			{
					"id": "ripple", 
					"name": "Ripple", 
					"symbol": "XRP", 
					"rank": "4", 
					"price_usd": "0.25052", 
					"price_btc": "0.00002659", 
					"24h_volume_usd": "127260000.0", 
					"market_cap_usd": "9675801495.0", 
					"available_supply": "38622870411.0", 
					"total_supply": "99993173757.0", 
					"max_supply": "100000000000", 
					"percent_change_1h": "0.16", 
					"percent_change_24h": "-1.05", 
					"percent_change_7d": "8.41", 
					"last_updated": "1511721541"
			}, 
			{
					"id": "bitcoin-gold", 
					"name": "Bitcoin Gold", 
					"symbol": "BTG", 
					"rank": "5", 
					"price_usd": "344.613", 
					"price_btc": "0.0365713", 
					"24h_volume_usd": "159528000.0", 
					"market_cap_usd": "5745779447.0", 
					"available_supply": "16673136.0", 
					"total_supply": "16773136.0", 
					"max_supply": "21000000.0", 
					"percent_change_1h": "0.08", 
					"percent_change_24h": "-1.45", 
					"percent_change_7d": "148.88", 
					"last_updated": "1511721577"
			}
	]`)
	}))
	defer testServer.Close()

	coinmarketcap := crypto.NewCoinmarketcap()
	coinmarketcap.SetBaseURL(testServer.URL)
	coins, err := coinmarketcap.GetAllCoinData("USD", 5)

	assert.NoError(t, err)

	assert.Equal(t, 5, len(coins))
	assert.Equal(t, "Bitcoin", coins[0].Name)
	assert.Equal(t, "BTC", coins[0].Symbol)

	assert.Equal(t, "Bitcoin Gold", coins[4].Name)
	assert.Equal(t, "BTG", coins[4].Symbol)
}
