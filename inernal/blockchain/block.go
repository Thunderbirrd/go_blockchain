package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatalf("Error while encoding block: %s", err.Error())
	}

	return result.Bytes()
}

func DeserializeBlock(bytesBlock []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(bytesBlock))

	err := decoder.Decode(&block)
	if err != nil {
		log.Fatalf("Error while decoding block: %s", err.Error())
	}

	return &block
}
