package cli

import (
	"fmt"
	"github.com/Thunderbirrd/go_blockchain/inernal/blockchain"
	"github.com/boltdb/bolt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !blockchain.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(bc.DB)

	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
