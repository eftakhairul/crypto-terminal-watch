package cmd


import (
	"os"
	"fmt"
	"github.com/eftakhairul/crypto-terminal-watch/crypto"
	"github.com/olekukonko/tablewriter"
)

func Render(coins []crypto.Coin) {


	data :=  make( [][]string, len(coins))

	for i := 0; i < len(coins); i++ {
		data[i] = make([]string, 3)
		data[i][0] = coins[i].Name
		data[i][1] =coins[i].Symbol
		data[i][2] = fmt.Sprintf("%.10f %s", coins[i].Price, coins[i].Currency)
		fmt.Println(coins[i].Currency)
	}


	// Final Printing into console
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Symbol", "Price"})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
}