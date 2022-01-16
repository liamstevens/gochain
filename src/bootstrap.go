package gochain

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"net"
	"os"
)

type Node struct {
	nodeUID  []byte
	msgCount int
	UNL      []string
}

func newNode(uid, count, list) *Node {
	n := Node{nodeUID: uid, msgCount: count, UNL: list }
	return &n
}

selfNode := newNode(generateNodeUID(),0,[]string) //probably want a node runtime with functions to fetch from config instead.


//Generate initial UNL - contact curator server and select a random subset of all nodes provided.
//This ensures a good spread of UNL configurations but does mean that additional configuration
//should be done if this mechanism is undesired.
func fetchUNL(file string) {
	var config []string
	confFile, ferr := os.Open(file)
	if ferr != nil {
		fmt.Println("error:", ferr)
		return
	}
	scanner := bufio.NewScanner(confFile)
	for scanner.Scan() {
		config = append(config, scanner.Text())
	}
	//Contact curator server and save response
	curatorConn, _ := net.Dial("tcp", config[0])
	fmt.Fprint(curatorConn, constructMessage(0,self.nodeUID,UNLREQ,[]byte))

}

//Generate Node UID - used for stamping of messages sent by client and server
func generateNodeUID(filename string) []byte {
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
	return nodeUID
}


