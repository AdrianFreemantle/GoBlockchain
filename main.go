package main

import (
	"fmt"
	"log"

	"github.com/AdrianFreemantle/goblockchain/block"
	"github.com/AdrianFreemantle/goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	//Wallet
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 10)

	//Blockchain
	blockchain := block.NewBlockChain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 10, walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added: ", isAdded)

	//Mining
	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("A %d\n", blockchain.CalcuateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("B %d\n", blockchain.CalcuateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("M %d\n", blockchain.CalcuateTotalAmount(walletM.BlockchainAddress()))
}
