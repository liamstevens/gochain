package gochain

import (
	"time"
)

//Struct for the purpose of constructing messages.
type Message struct {
	id          int
	nodeId      []byte
	messageType int
	timestamp   []byte
	contents    []byte
}

//constants for messageType usage
const PING = 0x00
const PONG = 0x01
const NEW_LEDGER = 0x03
const NEW_TX = 0x04
const AGREEMENT = 0x05
const UNLREQ = 0x06
const UNLRESP = 0x07

func constructMessage(id int, nodeid []byte, mtype int, cont []byte) Message {
	m := Message{id, nodeid, mtype, []byte(time.Now().String()), cont}
	return m
}
