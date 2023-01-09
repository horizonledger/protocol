package main

import (
	"encoding/json"
	"log"
)

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
