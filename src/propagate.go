package gochain

import (
	"bufio"
	"log"
	"os"
)

var peers []string

func loadConfig(conf string) {
	file, err := os.Open(conf)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		peers = append(peers, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//Interprocess communications for consensus.
//When the server process receives transactions, validate any new transactions and add valid transactions to current ledger.
//
