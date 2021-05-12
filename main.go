package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/happyeric77/golangSolidityOracle/contractCtl"
	"github.com/happyeric77/golangSolidityOracle/oracleServer"
	"github.com/joho/godotenv"
)

func loadChainInfo() (*ethclient.Client, *ecdsa.PrivateKey, error) {
	client, err := ethclient.Dial(fmt.Sprintf("wss://kovan.infura.io/ws/v3/%s", os.Getenv("INFURA_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		return nil, nil, err
	}
	return client, privateKey, nil
}

func main() {
	godotenv.Load("./.env")

	client, privateKey, err := loadChainInfo()
	if err != nil {
		log.Fatal(err)
	}
	_ = client
	_ = privateKey

	// contractInstance, ca, txa, err := contractCtl.Deploy(client, privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _ = contractInstance
	// _ = ca
	// _ = txa

	oracleInstance, err := contractCtl.Load(client)
	if err != nil {
		log.Fatal()
	}

	currentPrice, err := oracleInstance.CurrencyPrice(nil)
	if err != nil {
		log.Fatal("ERR: main.go Get currentPrice fail")
	}
	fmt.Printf("Current Price: %v\n", currentPrice)

	auth, err := contractCtl.TxOps(client, privateKey, 0)
	if err != nil {
		log.Fatal(fmt.Sprintf("ERR (contractCtl.Deply): %v\n", err))
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go oracleServer.SubscribeEventLog(client, oracleInstance, privateKey)

	tx, err := oracleInstance.RequestCurrencyPrice(auth, big.NewInt(10))
	fmt.Printf("Request oracle to update price: %x", tx.Hash())

	wg.Wait()
}
