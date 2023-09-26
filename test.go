package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://ws.apothem.network/ws")
	fmt.Println("Main Called")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to WSS")

	addresses := []common.Address{
		common.HexToAddress("0x9abf23ba30374a935188508af5d30a59e647a9ba"),
		common.HexToAddress("0x82142e063c6860a266490d0f40a695a3754c00e7"),
		common.HexToAddress("0x6df8cc35b894dd28a94919b93162fff2a59a8de9"),
		common.HexToAddress("0x4fdb6ef10478d8f2093cd55ea094701881bdc28d"),
		common.HexToAddress("0xc90cbcc24b7b52dab7f667ad4324e320453d9752"),
	}

	topics := [][]common.Hash{
		{
			common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
			common.HexToHash("0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93"),
		},
		{
			common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
			common.HexToHash("0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93"),
		},
		{
			common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
			common.HexToHash("0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93"),
		},
		{
			common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
			common.HexToHash("0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93"),
		},
		{
			common.HexToHash("0xd8d7ecc4800d25fa53ce0372f13a416d98907a7ef3d8d3bdd79cf4fe75529c65"),
			common.HexToHash("0xa7842b9ec549398102c0d91b1b9919b2f20558aefdadf57528a95c6cd3292e93"),
		},
	}

	query := ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    topics,
	}
	fmt.Println("Query that is getting subscribed", query)
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	fmt.Println("Subscription error status:", err)
	fmt.Println("Subscription data", sub)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
