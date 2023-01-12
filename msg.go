package protocol

//https://gist.github.com/benjyz/97d904ff0fd5fab96d666db973be9e3d
//https://stackoverflow.com/questions/29111777/is-it-possible-to-partially-decode-and-update-json-go
//json.RawMessage
//https://en.wikipedia.org/wiki/Multiple_dispatch

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

// generic message object which can hold messages and tx
type Gen struct {
	Type string `json:"type"`
	// the value is a msg or tx struct
	Value json.RawMessage `json:"value,omitempty"`
	//Sender uuid.UUID       `json:"uuid,omitempty"`
	Sender string    `json:"sender,omitempty"`
	Time   time.Time `json:"time,omitempty"`
}

type Msg struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	//Sender uuid.UUID `json:"uuid,omitempty"`
	Sender string    `json:"sender,omitempty"`
	Time   time.Time `json:"time"`
}

//Amount   int    `json:"amount"`
//Sender   string `json:"sender"`
//Receiver string `json:"receiver"`
//Nonce    int    `json:"nonce"`

func ParseMessageFromBytes(p []byte) Msg {

	msg := Msg{}
	err := json.Unmarshal(p, &msg)
	if err != nil {
		log.Error("couldnt ParseMessageFromBytes ", string(p))
		// fmt.Printf("type: %s\n", msg.Type)
		// fmt.Printf("value: %s\n", msg.Value)
	}
	return msg
}

func ParseMessageToBytes(msg Msg) []byte {

	msgByte, err := json.Marshal(msg)
	if err != nil {
		log.Error("couldnt ParseMessageToBytes ", msg)
		// fmt.Printf("type: %s\n", msg.Type)
		// fmt.Printf("value: %s\n", msg.Value)
	}

	return msgByte

}

func ParseGenToBytes(msg Gen) []byte {

	msgByte, err := json.Marshal(msg)
	if err != nil {
		log.Error("couldnt ParseGenToBytes ", msg)
	}

	return msgByte

}

func ParseGenFromBytes(p []byte) Gen {

	msg := Gen{}
	err := json.Unmarshal(p, &msg)
	if err != nil {
		log.Error("couldnt ParseGenFromBytes ", string(p))
	}
	return msg
}

func MsgToGen(msg Msg) Gen {
	jxmsg := ParseMessageToBytes(msg)
	gen := Gen{Type: "Msg", Value: jxmsg}
	return gen
}

func TxToGen(tx NameTx) Gen {
	jxmsg := ParseTxToBytes(tx)
	gen := Gen{Type: "NameTx", Value: jxmsg}
	return gen
}
