package cli

import (
	"fmt"
	"github.com/Thunderbirrd/go_blockchain/inernal/blockchain"
)

func (cli *CLI) reindexUTXO(nodeID string) {
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
