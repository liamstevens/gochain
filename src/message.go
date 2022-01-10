package gochain

//Struct for the purpose of constructing messages.
type Message struct {
	id          int
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
