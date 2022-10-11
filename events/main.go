package events

import (
	token "ERC20/contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"strings"
)

func GetEvent() (string, string, string, error) {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/758643e59476416e93ab5a4d873b5ccd")
	if err != nil {
		return "", "", "", errors.Wrap(err, "connection error")

	}

	fmt.Println("we have a connection")

	contractAddress := common.HexToAddress("0x24cd898BBeb565b9EF692c4C062147a3ADc0A8b6")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2027094),
		ToBlock:   big.NewInt(20270940),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(token.ContractsABI)))
	if err != nil {
		return "", "", "", errors.Wrap(err, "MainABI error")
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return "", "", "", errors.Wrap(err, "Sub error")
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent token.ContractsTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				return "", "", "", errors.Wrap(err, "UnpackIntoInterface error")
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())

			fmt.Printf("\n\n") // pointer to event log
			return transferEvent.From.Hex(), transferEvent.To.Hex(), transferEvent.Value.String(), nil
		}
	}

}
