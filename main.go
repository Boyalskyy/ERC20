package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"strings"

	token "ERC20/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/758643e59476416e93ab5a4d873b5ccd")
	if err != nil {
		log.Fatal(errors.Wrap(err, "connection error"))
		return

	}

	fmt.Println("we have a connection")

	contractAddress := common.HexToAddress("0x20ec06104035d0f2F9846960a279AA0eCC298dbB")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2027094),
		ToBlock:   big.NewInt(20270940),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(token.MainMetaData.ABI)))
	if err != nil {
		log.Fatal(errors.Wrap(err, "MainABI error"))
		return
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Sub error"))
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Name: Transfer\n")

			var transferEvent token.MainTransfer

			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(errors.Wrap(err, "UnpackIntoInterface error"))
				return
			}

			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", transferEvent.From.Hex())
			fmt.Printf("To: %s\n", transferEvent.To.Hex())
			fmt.Printf("Tokens: %s\n", transferEvent.Value.String())

			fmt.Printf("\n\n") // pointer to event log
		}
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "FilterLogs error"))
	//	return
	//}

	//contractAbi, err := abi.JSON(strings.NewReader(string(token.MainMetaData.ABI)))
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "MainABI error"))
	//	return
	//}
	//
	//logTransferSig := []byte("Transfer(address,address,uint256)")
	//logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	//
	//for _, vLog := range logs {
	//	fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
	//	fmt.Printf("Log Index: %d\n", vLog.Index)
	//
	//	switch vLog.Topics[0].Hex() {
	//	case logTransferSigHash.Hex():
	//		fmt.Printf("Log Name: Transfer\n")
	//
	//		var transferEvent token.MainTransfer
	//
	//		err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
	//		if err != nil {
	//			log.Fatal(errors.Wrap(err, "UnpackIntoInterface error"))
	//			return
	//		}
	//
	//		transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
	//		transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
	//
	//		fmt.Printf("From: %s\n", transferEvent.From.Hex())
	//		fmt.Printf("To: %s\n", transferEvent.To.Hex())
	//		fmt.Printf("Tokens: %s\n", transferEvent.Value.String())
	//
	//		fmt.Printf("\n\n")
	//	}
	//
	//}
}
