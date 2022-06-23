package cli

import (
	"fmt"
	"github.com/Thunderbirrd/go_blockchain/inernal/blockchain"
	"github.com/boltdb/bolt"
	"log"
)

func (cli *CLI) getBalance(address, nodeID string) {
	if !blockchain.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(bc.DB)

	balance := 0
	pubKeyHash := blockchain.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
