package cli

import (
	"fmt"
	"github.com/Thunderbirrd/go_blockchain/inernal/blockchain"
)

func (cli *CLI) createWallet(nodeID string) {
	wallets, _ := blockchain.NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
