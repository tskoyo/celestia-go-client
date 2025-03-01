package main

import (
	"context"
	"fmt"

	client "github.com/celestiaorg/celestia-openrpc"
)

func main() {
	rootCtx := context.Background()
	addr := "http://192.168.0.56:2665"
	token := ""

	client, err := client.NewClient(rootCtx, addr, token)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	peerAddr, err := client.P2P.Info(rootCtx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Peer addr: ", peerAddr.ID.String())
}
