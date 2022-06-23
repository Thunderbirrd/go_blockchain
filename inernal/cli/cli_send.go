package cli

import (
	"fmt"
	"github.com/Thunderbirrd/go_blockchain/inernal/blockchain"
	"github.com/Thunderbirrd/go_blockchain/inernal/server"
	"github.com/boltdb/bolt"
	"log"
)

func (cli *CLI) send(from, to string, amount int, nodeID string, mineNow bool) {
	if !blockchain.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !blockchain.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{Blockchain: bc}
	defer func(DB *bolt.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(bc.DB)

	wallets, err := blockchain.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

	tx := blockchain.NewUTXOTransaction(&wallet, to, amount, &UTXOSet)

	if mineNow {
		cbTx := blockchain.NewCoinbaseTX(from, "")
		txs := []*blockchain.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		server.SendTx(server.KnownNodes[0], tx)
	}

	fmt.Println("Success!")
}
