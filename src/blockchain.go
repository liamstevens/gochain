package gochain

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

//Define basal struct for blockchain and methods for use in composition of token and ledgers
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// This will join our previous block's relevant info with the new blocks
	hash := sha256.Sum256(info)
	//This performs the actual hashing algorithm
	b.Hash = hash[:]
}

func CreateBlock(data []byte, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data []byte) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock([]byte("Genesis"), []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (b *Block) SignBlock(privKey *rsa.PrivateKey) []byte {
	rng := rand.Reader
	signature, err := rsa.SignPKCS1v15(rng, privKey, crypto.SHA256, b.Hash[:])
	if err != nil {
		return nil
	} else {
		return signature
	}
}

func (b *Block) VerifyBlock(pubKey *rsa.PublicKey, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, b.Hash[:], signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return false
	} else {
		return true
	}
}
