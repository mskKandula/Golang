package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timeStamp    time.Time
	transactions []string
	prevHash     []byte
	hash         []byte
}

func newBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()
	return &Block{
		timeStamp:    currentTime,
		transactions: transactions,
		prevHash:     prevHash,
		hash:         newHash(currentTime, transactions, prevHash),
	}

}

func newHash(currentTime time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, currentTime.String()...)

	for _, transaction := range transactions {
		input = append(input, transaction...)
	}

	hash := sha256.Sum256(input)

	return hash[:]
}

func printBlockInfo(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timeStamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.hash)
	printTransactions(block)
}

func printTransactions(block *Block) {

	fmt.Println("\tTransactions :")

	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v : %q\n", i, transaction)
	}
}

func main() {

	genesisTransactions := []string{"mohan sent krishna 50 bitcoins", "krishna sent mohan 30 bitcoins"}
	genesisBlock := newBlock(genesisTransactions, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInfo(genesisBlock)

	block2Transactions := []string{"sai sent mohan 30 bitcoins"}
	block2 := newBlock(block2Transactions, genesisBlock.hash)
	fmt.Println("--- Second Block ---")
	printBlockInfo(block2)

	block3Transactions := []string{"krishna sent mohan 45 bitcoins", "mohan sent sai 10 bitcoins"}
	block3 := newBlock(block3Transactions, block2.hash)
	fmt.Println("--- Third Block ---")
	printBlockInfo(block3)

}
