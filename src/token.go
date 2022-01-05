package gochain

type Token struct {
	value             float64
	hash              []byte
	previousHolderSig []byte
	newHolderSig      []byte
}
