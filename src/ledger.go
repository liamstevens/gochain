package gochain

type Transaction struct {
	source      []byte
	destination []byte
	volume      int
}

type Ledger struct {
	transactions   []Transaction
	previousLedger *Ledger
	id             int
}

func createTransaction()
