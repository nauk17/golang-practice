package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/quannv132/practice_0/proto"

	"google.golang.org/grpc"
)

var client proto.BlockchainClient

func main() {
	addFlag := flag.Bool("add", false, "add new block")
	listFlag := flag.Bool("list", false, "get block chain")

	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot dial server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)

	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("Unable to add block: %v", err)
	}

	log.Printf("New block hash : %s", block.Hash)
}

func getBlockchain() {
	bc, err := client.GetBlockChain(context.Background(), &proto.GetBlockchainRequest{})
	if err != nil {
		log.Fatalf("Unable to add block: %v", err)
	}
	log.Println("blocks:")
	for _, b := range bc.Blocks {
		log.Printf("hash: %s, prevBlockHash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
	}
}
