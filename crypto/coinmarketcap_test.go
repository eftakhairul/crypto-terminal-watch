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
}
