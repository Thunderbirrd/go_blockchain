package main

import (
	"github.com/Thunderbirrd/go_blockchain/inernal/blockchain"
	cliPackage "github.com/Thunderbirrd/go_blockchain/inernal/cli"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			log.Panic(err)
		}
	}(bc.DB)

	cli := cliPackage.CLI{BC: bc}
	cli.Run()
}
