package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	MINING_DIFFICULTY = 3
	MINING_SENDER     = "THE BLOCKCHAIN"
	MINING_REWARD     = 1.0
)

type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp:      %d\n", b.timestamp)
	fmt.Printf("nonce:          %d\n", b.nonce)
	fmt.Printf("previous_hash:  %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

func NewBlockChain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value int64) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions,
			NewTransaction(t.senderBlockchainAddress,
				t.recipientBlockchainAddress,
				t.value))
	}
	return transactions
}

func (bc *Blockchain) ValidateProof(nonce int, previousHash [32]byte, transactions []*Transaction, diffculty int) bool {
	zeros := strings.Repeat("0", diffculty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:diffculty] == zeros
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) CalcuateTotalAmount(blockchainAdress string) int64 {
	var totalAmount int64 = 0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			if blockchainAdress == t.recipientBlockchainAddress {
				totalAmount += t.value
			}
			if blockchainAdress == t.senderBlockchainAddress {
				totalAmount -= t.value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidateProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce += 1
	}
	return nonce
}

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      int64
}

func NewTransaction(sender string, recipient string, value int64) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address   	%s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address	%s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                          %d\n", t.value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string `json:"recipient_blockchain_address"`
		Value                      int64  `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

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
