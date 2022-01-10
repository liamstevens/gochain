package gochain

//Struct for the purpose of constructing messages.

type Message struct {
	id          int
	messageType int
	timestamp   []byte
	contents    []byte
}
