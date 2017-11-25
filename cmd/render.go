package cmd

import (
	"fmt"
	"os"

	"github.com/eftakhairul/crypto-terminal-watch/crypto"
	"github.com/olekukonko/tablewriter"
)

// 	Render all coins in a nice tabular format in terminal
func Render(coins []crypto.Coin) {

	if len(coins) == 0 {
		fmt.Println("Sorry! Please try agian. To get info about BTC in CAD: ctw --currency CAD BTC")
		return
	}

	data := make([][]string, len(coins))

	for i := 0; i < len(coins); i++ {
		data[i] = []string{coins[i].Name, coins[i].Symbol, fmt.Sprintf("%.10f %s", coins[i].Price, coins[i].Currency)}
	}

	// Final Printing into console
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Symbol", "Price"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}
