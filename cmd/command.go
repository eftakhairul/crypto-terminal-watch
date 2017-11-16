// Package cmd provides terminal interface for crypto coins
package cmd

import (
	"fmt"
	"os"

	"github.com/eftakhairul/crypto-terminal-watch/crypto"
	"github.com/urfave/cli"
)

func Execute() {
	app := cli.NewApp()
	app.Name = "crypto-terminal-watch"
	app.Usage = "Checking crypto market"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "currency, c",
			Value: "USD",
			Usage: "The currency you want see the price",
		},
	}

	app.Action = func(c *cli.Context) error {
		ccm := crypto.New()
		name := "ALL"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}

		if c.String("currency") == "spanish" {
			fmt.Println("Hola", name)
		}

		coins, _ := ccm.GetCoinData(name)
		Render(coins)

		return nil
	}

	app.Run(os.Args)
}
