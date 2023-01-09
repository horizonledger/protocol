package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type Msg struct {
	Type   string    `json:"type"`
	Value  string    `json:"value"`
	Sender uuid.UUID `json:"uuid,omitempty"`
	Time   time.Time `json:"time,omitempty"`
}

func parseMessageFromBytes(p []byte) Msg {

	msg := Msg{}
	err := json.Unmarshal(p, &msg)
	if err != nil {
		log.Println("couldnt parse message")
		// fmt.Printf("type: %s\n", msg.Type)
		// fmt.Printf("value: %s\n", msg.Value)
	}
	return msg
}

func parseMessageToBytes(msg Msg) []byte {

	msgByte, err := json.Marshal(msg)
	if err != nil {
		log.Println("couldnt parse message")
		// fmt.Printf("type: %s\n", msg.Type)
		// fmt.Printf("value: %s\n", msg.Value)
	}

	return msgByte

}
