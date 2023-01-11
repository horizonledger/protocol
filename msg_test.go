package protocol

import (
	"testing"
)

func TestGen(t *testing.T) {

	gmsg := Gen{Type: "Msg"}
	msg := Msg{Type: "HNDCLIENT", Value: "hello"}
	x := ParseMessageToBytes(msg)
	gmsg.Value = x
	x = ParseGenToBytes(gmsg)
	//back
	gen := ParseGenFromBytes(x)

	if gen.Type != "Msg" {
		t.Error("error decoding")
	}

	z := ParseMessageFromBytes(gen.Value)

	if z.Type != "HNDCLIENT" {
		t.Error("error decoding")
	}

}
