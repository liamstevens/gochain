package gochain

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
)

//Generate initial UNL - contact curator server and select a random subset of all nodes provided.
//This ensures a good spread of UNL configurations but does mean that additional configuration
//should be done if this mechanism is undesired.

//Generate Node UID - used for stamping of messages sent by client and server
func generateNodeUID(filename string) {
	nodeUID := make([]byte, 32)
	confFile, ferr := os.Open(filename)
	if ferr != nil {
		fmt.Println("error:", ferr)
		return
	}
	_, err := rand.Read(nodeUID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	w := bufio.NewWriter(confFile)
	w.Write(nodeUID)
	confFile.Close()
}
