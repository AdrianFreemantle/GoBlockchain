package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "my_address"
	bc := NewBlockChain(myBlockchainAddress)
	bc.Print()

	bc.AddTransaction("A", "B", 10)
	bc.Mining()

	bc.AddTransaction("C", "D", 95)
	bc.AddTransaction("E", "F", 12)
	bc.Mining()

	bc.Print()

	fmt.Printf("My balance: %d\n", bc.CalcuateTotalAmount(myBlockchainAddress))
	fmt.Printf("A balance: %d\n", bc.CalcuateTotalAmount("A"))
	fmt.Printf("B balance: %d\n", bc.CalcuateTotalAmount("B"))
	fmt.Printf("C balance: %d\n", bc.CalcuateTotalAmount("C"))
	fmt.Printf("D balance: %d\n", bc.CalcuateTotalAmount("D"))
	fmt.Printf("E balance: %d\n", bc.CalcuateTotalAmount("E"))
	fmt.Printf("F balance: %d\n", bc.CalcuateTotalAmount("F"))
}

func NewBlockChain(myBlockchainAddress string) {
	panic("unimplemented")
}
