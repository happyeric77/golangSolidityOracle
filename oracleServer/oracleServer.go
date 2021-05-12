package oracleServer

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/happyeric77/golangSolidityOracle/contractCtl"
	"github.com/happyeric77/golangSolidityOracle/contracts/currencyOracle"
)

func SubscribeEventLog(client *ethclient.Client, oracleInstance *currencyOracle.CurrencyOracle, privateKey *ecdsa.PrivateKey) {
	topic := crypto.Keccak256([]byte("RequestCurrencyPrice(uint256)"))
	query := ethereum.FilterQuery{
		Topics: [][]common.Hash{
			{
				common.BytesToHash(topic),
			},
		},
	}

	// Subscribe certain log (by FilterLogs function)
	events := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, events)
	if err != nil {
		log.Fatal(err)
	}

	// parse the abi document
	parsedABI, err := abi.JSON(strings.NewReader(currencyOracle.CurrencyOracleABI))

	for {
		fmt.Println("Listening the event")
		select {
		case err := <-sub.Err():
			log.Fatal(err)
			break
		case evt := <-events:
			unpacked, err := parsedABI.Unpack("RequestCurrencyPrice", evt.Data)
			if err != nil {
				log.Fatalf("oracleServer.SubscribEventLogs: parsedABI.Unpack failed: %v", err)
			}

			unpackedValue, ok := unpacked[0].(*big.Int)
			if !ok {
				log.Fatal("oracleServer.SubscribEventLogs: unpackedValue is not a pointer of big int")
			}
			fmt.Printf("Currency code to update: %v\n", unpackedValue)

			auth, err := contractCtl.TxOps(client, privateKey, 0)
			if err != nil {
				log.Fatal(fmt.Sprintf("oracleServer.SubscribEventLogs:Fail to update auth: %v\n", err))
			}

			// TODO: Call currency api to get updated currency value
			resp, err := http.Get("https://ethgasstation.info/api/ethgasAPI.json?")
			if err != nil {
				log.Fatalf("ERR: http.Get currency API: %v ", err)
			}
			fmt.Println(resp.Status)
			type Currency struct {
				Fast big.Int `json:"fast"`
			}
			body, _ := ioutil.ReadAll(resp.Body)
			var curr Currency
			err = json.Unmarshal(body, &curr)
			if err != nil {
				log.Fatalf("ERR: Unmarshal data Fail: %v", err)
			}
			fmt.Printf("Price will be updated to: %s\n", curr.Fast.String())

			tx, err := oracleInstance.UpdateCurrencyPrice(auth, &curr.Fast)
			if err != nil {
				log.Fatalf("oracleServer.SubscribEventLogs: Fail to call UpdateCurrencyPrice(): %v", err)
			}
			fmt.Printf("Price is updated: %x\n", tx.Hash())

		}
	}
}
