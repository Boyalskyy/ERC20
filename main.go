package main

import (
	"ERC20/events"
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
	form,to,amount,err:=events.GetEvent()
	if err!=nil{
		log.Println(err)
	}
	db, err:=

}
