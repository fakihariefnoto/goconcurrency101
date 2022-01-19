package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAllAccount(shopID string) ([]AccountInfo, error) {
	var data []AccountInfo
	url := fmt.Sprintf("%s/shopinfo/%s/saldoinfo", base_url, shopID)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return data, err
	}

	response, err := client.Do(req)

	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	byteResponse, err := io.ReadAll(response.Body)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(byteResponse, &data)

	return data, err
}

func GetAccountInfoByAccountID(shopID, accountID string) (AccountInfo, error) {
	var data AccountInfo
	url := fmt.Sprintf("%s/shopinfo/%s/saldoinfo/%s", base_url, shopID, accountID)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return data, err
	}

	response, err := client.Do(req)

	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	byteResponse, err := io.ReadAll(response.Body)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(byteResponse, &data)

	return data, err
}
