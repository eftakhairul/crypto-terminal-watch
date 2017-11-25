// Package cmd provides terminal interface for crypto coins
package cmd

import (
	"fmt"
	"os"

	"github.com/eftakhairul/crypto-terminal-watch/crypto"
	"github.com/urfave/cli"
)

// Run the command app
// Take input from command line and prints all info into tabular format
func Execute() {
	var app = cli.NewApp()
	app.Name = "crypto-terminal-watch  (ctw)"
	app.Usage = "Checking crypto market \n Example: ctw --c CAD BTC"
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
			Value: 10,
			Usage: "Amount of crypto currency info in one request",
		},
	}

	app.Action = func(c *cli.Context) error {
		var err error
		var coins []crypto.Coin

		var lookupCryptoCurrency = "ALL"
		var crypto = crypto.NewCoinmarketcap()

		if c.NArg() > 0 {
			lookupCryptoCurrency = c.Args().Get(0)
		}

		if lookupCryptoCurrency == "ALL" {
			coins, err = crypto.GetAllCoinData(c.String("currency"), c.Int("limit"))
		} else {
			coins, err = crypto.GetCoinData(lookupCryptoCurrency, c.String("currency"))
		}

		if err != nil {
			printOnError(err)
			return nil
		}

		Render(coins)
		return nil
	}

	app.Run(os.Args)
}

// Print error
func printOnError(err error) {
	fmt.Println("Something is went wrong. Please send an email to: eftakhairul@gmail.com")
}
