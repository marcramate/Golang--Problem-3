package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const URLsend = "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
const URLcheck = "https://mock-node-wgqbnxruha-as.a.run.app/check/"

type payloadMock struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type Response struct {
	Tx_hash string `json:"tx_hash"`
}

type Status struct {
	Tx_status string `json:"tx_status"`
}

func main() {

	payload :=
		payloadMock{
			Symbol:    "ETH",
			Price:     4500,
			Timestamp: 1678912345,
		}

	txhash, err := sendTransection(payload)
	if err != nil {
		fmt.Println("Faild transection Data:", err)
		return
	}

	fmt.Println("Transaction ", txhash)

	monitorStatus(txhash, 5*time.Second)

}

func sendTransection(payload payloadMock) (string, error) {
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	resp, err := http.Post(URLsend, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var txResponse Response
	err = json.NewDecoder(resp.Body).Decode(&txResponse)

	return txResponse.Tx_hash, nil
}

func checkStatus(txHash string) (string, error) {
	resp, err := http.Get(URLcheck + txHash)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var resstatus Status
	err = json.NewDecoder(resp.Body).Decode(&resstatus)
	if err != nil {
		return "", err
	}

	return resstatus.Tx_status, nil
}

func monitorStatus(txHash string, interval time.Duration) {
	for {
		status, err := checkStatus(txHash)
		if err != nil {
			fmt.Println("Check Status NotFound:", err)
			return
		}
		fmt.Println("Status :", status)

		switch status {
		case "CONFIRMED":
			fmt.Println("Transaction has been processed and confirmed")
			return
		case "FAILED":
			fmt.Println("Transaction failed to process")
			return
		case "PENDING":
			fmt.Println("Transaction is awaiting processing")
		case "DNE":
			fmt.Println("Transaction does not exist")
		}
		time.Sleep(interval)
	}
}
