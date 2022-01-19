package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAllProduct(shopID string) ([]ProductInfo, error) {
	var data []ProductInfo
	url := fmt.Sprintf("%s/shopinfo/%s/productinfo", base_url, shopID)
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

func GetProductInfoByProductID(shopID, productID string) (ProductInfo, error) {
	var data ProductInfo
	url := fmt.Sprintf("%s/shopinfo/%s/productinfo/%s", base_url, shopID, productID)
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
