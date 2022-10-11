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



func main() {
	form,to,amount,err:=events.GetEvent()
	if err!=nil{
		log.Println(err)
	}
	db, err:=

}
