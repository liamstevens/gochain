package gochain

type Transaction struct {
	source      []byte
	destination []byte
	volume      int
	Block       //Anonymous member for referencing block attributes
}

type Ledger struct {
	length         int
	transactions   []Transaction
	previousLedger *Ledger
	id             int
}

func createTransaction(s []byte, d []byte, v int, mint bool) *Transaction {
	if mint {
		t := &Transaction{s, d, v, *CreateBlock(string(v), []byte{})}
		return t
	} else {
		t := &Transaction{s, d, v, *CreateBlock(string(v), s)}
		return t
	}
}

func createLedger(le int, t []Transaction, p *Ledger, i int) *Ledger {
	l := &Ledger{le, t, p, i}
	return l
}

//TODO: transaction validation
func (ledger *Ledger) addTransaction(t Transaction) {
	if len(ledger.transactions) < ledger.length {
		ledger.transactions = append(ledger.transactions, t)
	}
}

func (ledger *Ledger) validateTransaction(t Transaction) {

}
