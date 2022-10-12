package events

import (
	token "ERC20/contracts"
	db2 "ERC20/db"
	"context"
	_ "database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"strings"
)

const (
	transferEvent = "Transfer"
	mintEvent     = "Mint"
)

func GetEvent(db *sqlx.DB) {
	eventsQuery := db2.NewEvents(db)
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/758643e59476416e93ab5a4d873b5ccd")
	if err != nil {
		log.Println(err)

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
		log.Println(err)
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Name: Transfer\n")

			var event token.ContractsTransfer

			err := contractAbi.UnpackIntoInterface(&event, "Transfer", vLog.Data)
			if err != nil {
				log.Println(err)
			}

			event.From = common.HexToAddress(vLog.Topics[1].Hex())
			event.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("From: %s\n", event.From.Hex())
			fmt.Printf("To: %s\n", event.To.Hex())
			fmt.Printf("Tokens: %s\n", event.Value.String())

			fmt.Printf("\n\n") // pointer to event log
			if event.From.Hex() == "0x0000000000000000000000000000000000000000" {
				err := eventsQuery.Create(mintEvent, event.From.Hex(), event.To.Hex(), event.Value.String())
				if err != nil {
					log.Println(errors.Wrap(err, "Insert error"))
				}
			} else {
				err := eventsQuery.Create(transferEvent, event.From.Hex(), event.To.Hex(), event.Value.String())
				if err != nil {
					log.Println(errors.Wrap(err, "Insert error"))
				}
			}

		}

	}

}
