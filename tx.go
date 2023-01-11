package protocol

import (
	"encoding/json"
	"log"
)

// uname-register
// uname-transfer
// dname-register. dns record contanis public key. verify and make onchain tx
// transfer-value. money transaction

type NameTx struct {
	Type   string `json:"type"`
	Action string `json:"action"`
	Name   string `json:"name"`
	//can sender just be a name? could require name like random123. what if ownership changes
	SenderID     string `json:"senderID,omitempty"`
	SenderPubkey string `json:"senderPubkey,omitempty"`
	Signature    string `json:"signature,omitempty"`
	//Time time.Time `json:"time,omitempty"`
}

func ParseTxToBytes(tx NameTx) []byte {

	txByte, err := json.Marshal(tx)
	if err != nil {
		log.Println("couldnt parse message")
	}

	return txByte

}

func ParseTxFromBytes(txb []byte) NameTx {

	tx := NameTx{}
	err := json.Unmarshal(txb, &tx)
	if err != nil {
		log.Println("couldnt parse message")
	}
	return tx

}
