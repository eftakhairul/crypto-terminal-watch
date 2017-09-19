// Package cmd provides terminal interface for crypto coins
package cmd


import ("os"
		"github.com/urfave/cli"
		"github.com/eftakhairul/crypto-terminal-watch/crypto"
	)


func Execute() {
	app := cli.NewApp()
	app.Name = "crypto-terminal-watch"
	app.Usage = "Checking crypto market"


	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "api",
			Value: "coinmarketcap",
			Usage: "API to fetch currency information",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "ethereum"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}

		ccm := crypto.New()
		coins, _ := ccm.GetCoinData(name)
		Render(coins)

		return nil
	}

	app.Run(os.Args)
}