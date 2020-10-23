package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rpc.energyweb.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	block, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block number", block)

	record := make(map[string]bool, 0)
	bn := new(big.Int).SetUint64(block)
	for {
		b, err := client.BlockByNumber(context.Background(), bn.Sub(bn, new(big.Int).SetUint64(1)))
		if err != nil {
			log.Fatal(err)
		}
		key := string(b.Extra()[:])
		if record[key] {
			log.Println("found duplicate : ", key)
			log.Printf(" %d validators found", len(record))
			break
		}
		record[key] = true
		fmt.Println("block extra details", key)
	}

	for k := range record {
		fmt.Printf("%s\n", k)
	}
}
