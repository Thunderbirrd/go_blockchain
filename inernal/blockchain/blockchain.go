package blockchain

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.DB"
const blocksBucket = "blocks"

type Blockchain struct {
	tip []byte
	DB  *bolt.DB
}

type BcIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		lastHash = bucket.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Fatalf("Error while viewing DB: %s", err.Error())
	}

	newBlock := NewBlock(data, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			return err
		}
		bc.tip = newBlock.Hash
		return nil
	})

	if err != nil {
		log.Fatalf("Error while updating DB: %s", err.Error())
	}

}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatalf("Error while opening DB file: %s", err.Error())
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		if bucket == nil {
			genesis := NewGenesisBlock()
			bucket, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}

			err = bucket.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}

			err = bucket.Put([]byte("l"), genesis.Hash)
			if err != nil {
				return err
			}
			tip = genesis.Hash
		} else {
			tip = bucket.Get([]byte("l"))
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error while updating DB: %s", err.Error())
	}

	bc := Blockchain{tip, db}
	return &bc
}

func (bc *Blockchain) Iterator() *BcIterator {
	bci := &BcIterator{bc.tip, bc.DB}

	return bci
}

func (i *BcIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bucket.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Fatalf("Error while viewing DB: %s", err.Error())
	}

	i.currentHash = block.PrevBlockHash

	return block
}
