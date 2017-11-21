// Package cmd provides terminal interface for crypto coins
package cmd

import (
	"os"

	"github.com/eftakhairul/crypto-terminal-watch/crypto"
	"github.com/urfave/cli"
)

func Execute() {
	var app = cli.NewApp()
	app.Name = "crypto-terminal-watch  (ctw)"
	app.Usage = "Checking crypto market"
	app.Version = "0.0.1"

	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Eftakhairul Islam",
			Email: "eftakhairul@gmail.com",
		},
	}
	app.Copyright = "(c) 2017 Eftakhairul Islam"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "currency, c",
			Value: "USD",
			Usage: "The currency you want see the price",
		},
		cli.IntFlag{
			Name:  "limit, l",
			Value: 100,
			Usage: "Not implemented yet",
		},
	}

	app.Action = func(c *cli.Context) error {
		var coins []crypto.Coin
		var crypto = crypto.New()
		var cryptoCurrency = "ALL"
		var currencyForConversation = "USD"

		if c.NArg() > 0 {
			cryptoCurrency = c.Args().Get(0)
		}

		if c.String("currency") != "" {
			currencyForConversation = c.String("currency")
		}

		if cryptoCurrency == "ALL" {
			coins, _ = crypto.GetAllCoinData(currencyForConversation, 10)
		} else {
			coins, _ = crypto.GetCoinData(cryptoCurrency, currencyForConversation)
		}

		Render(coins)
		return nil
	}

	app.Run(os.Args)
}
