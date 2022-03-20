package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x667a05757f2c27694ed46Af0205E2d392C5853c3")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2394201),
		ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		event := struct {
			Address   [32]byte
			Message   string
			Timestamp uint32
		}{}

		err := contractAbi.Unpack(&event, "Message", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(event.Address[:]))
		fmt.Println(string(event.Message))
		fmt.Println(strconv.Itoa(int(event.Timestamp)))

		//var topics [4]string
		//for i := range vLog.Topics {
		//	topics[i] = vLog.Topics[i].Hex()
		//}

		fmt.Println(topics[0])
	}
}
