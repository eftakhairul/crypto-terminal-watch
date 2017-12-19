package cmd

import (
	"fmt"
	"os"
	"time"

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
		data[i] = []string{coins[i].Name,
			coins[i].Symbol,
			fmt.Sprintf("%.10f %s", coins[i].Price, coins[i].Currency),
			fmt.Sprintf("%.2f %s", coins[i].Volume24, coins[i].Currency),
			fmt.Sprintf("%v %s", coins[i].Maxsupply, coins[i].Symbol),
		}
	}

	// Final Printing into console
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Symbol", "Price", "Volume 24h", "Max Supply"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()

	fmt.Println("Updated at : ", time.Now().Format("2006-01-02 15:04:05"))
}
