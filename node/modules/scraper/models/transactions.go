package models

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

// ResultTransactions is the struct for the response
type ResultTransactions struct {
	Number           string
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	Miner            string
	Difficulty       string
	TotalDifficulty  string
	ExtraData        string
	Size             string
	GasUsed          string
	TimeStamp        string
	Transactions     []TransactionStruct
}

//ResultString is used when the type returned is a direct string
type ResultString struct {
	Result string
}

//Account is the struct whcih represents each individual etehreum address
type Account struct {
	gorm.Model
	Address string
	TxRaw   pq.ByteaArray `gorm:"type:varchar(255)[]"`
}

// TransactionStruct is the struct for each transaction in a block
type TransactionStruct struct {
	BlockHash   string
	BlockNumber string
	From        string
	To          string
	Gas         string
	GasPrice    string
	Hash        string
	Value       string
}

//Transaction Inside an account
type Transaction struct {
	To          string
	Value       string
	BlockNumber string
}

//Store will save the transaction Struct for the corresponding account
func (tx *TransactionStruct) Store() map[string]interface{} {
	//first lets check if the account in the transaction struct has exists
	account := &Account{}
	err := GetDB().Table("accounts").Where("address = ?", tx.From).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			account.Address = tx.From
			transaction := &Transaction{To: tx.To, Value: tx.Value, BlockNumber: tx.BlockNumber}
			txRaw, _ := json.Marshal(transaction)
			account.TxRaw = append(account.TxRaw, txRaw)
			fmt.Println("Saved for account: " + tx.From)
			GetDB().Create(account)
			return map[string]interface{}{
				"message": "Succesfully created new accounts",
				"status":  true,
			}
		} else {
			return map[string]interface{}{
				"message": "Failed to Save or Get Accounts",
				"status":  false,
			}
		}
	} else {
		transaction := &Transaction{To: tx.To, Value: tx.Value, BlockNumber: tx.BlockNumber}
		txRaw, _ := json.Marshal(transaction)
		account.TxRaw = append(account.TxRaw, txRaw)
		fmt.Println("Saved for account: " + tx.From)
		GetDB().Save(account)
		return map[string]interface{}{
			"message": "Succesfully saved to accounts",
			"status":  true,
		}
	}

}
