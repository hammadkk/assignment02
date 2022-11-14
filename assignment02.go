package assignment02

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce int
	BlockData        []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {

    //First add the whole block data and take it's hash, then convert the long hash into a hex and convert the string as a number for nonce.
	//Everytime it will give a different nonce as the data will be changed.

	datablock := ""
	for i := 0; i < len(blockData); i++ {
		datablock +=  blockData[i].TransactionID+ blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)+blockData[i].Sender
	}

	hashh := fmt.Sprintf("%x", sha256.Sum256([]byte(datablock)))
	hex := hex.EncodeToString([]byte(hashh[:len(hashh)-60]))
	res, errorr := strconv.ParseInt(hex, 16, 64) //Converts string into a number

	if errorr != nil {
		panic(errorr)
	}
	return int(res)
	
}



func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
		blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}




func NewBlock(blockData []Transaction, chainHead *Block) *Block {

	previoushash := ""

	if chainHead == nil {
		previoushash = "NONE"
		} else {
		previoushash = chainHead.CurrentHash
		}


		newBlock := Block{
			Nonce:       GenerateNonce(blockData),
			PrevHash:    previoushash,
			PrevPointer: chainHead,
			BlockData:   blockData,
			CurrentHash: CalculateHash(blockData, GenerateNonce(blockData)),
				
		}

	return &newBlock
	
}

func ListBlocks(chainHead *Block) {

	inc :=1

	for chainHead != nil {

		
		fmt.Println("\n------------------------------Block no------------------------- : ",inc)
		fmt.Printf("\nNonce : %d", chainHead.Nonce)
		fmt.Printf("\nCurrent Hash : %s", chainHead.CurrentHash)
		fmt.Printf("\nPrev Hash : %s", chainHead.PrevHash)
		DisplayTransactions(chainHead.BlockData)//Block data contains all the transaction attributes so call the fucntion
	
		chainHead = chainHead.PrevPointer
		inc ++
	  }
}

func DisplayTransactions(blockData []Transaction) {

	for a := 0; a < len(blockData); a++ {

		fmt.Printf("\n \nTransaction no : %d", a + 1)
		fmt.Printf("\n Sender : %s", blockData[a].Sender)
		fmt.Printf("\n Reciever : %s", blockData[a].Receiver)
		fmt.Printf("\n Amount Sent: %d", blockData[a].Amount)
		fmt.Printf("\n Transaction id : %s\n", blockData[a].TransactionID)
		
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {

	sum := sha256.Sum256([]byte(sender + receiver + strconv.Itoa((amount))))

	transaction := Transaction{
		
		TransactionID: hex.EncodeToString(sum[:]),
		Sender:        sender,
		Receiver:      receiver,
		Amount:        amount,
	}
	return transaction
}

