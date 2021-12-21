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

func createTransaction(s []byte, d []byte, v int) *Transaction {
	t := &Transaction{s, d, v}
	return t
}

func createLedger(t []Transaction, p *Ledger, i int) *Ledger {
	l := &Ledger{t, p, i}
	return l
}
