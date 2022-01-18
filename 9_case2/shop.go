package case2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAllShop() ([]ShopInfo, error) {
	var data []ShopInfo
	url := fmt.Sprintf("%s/shopinfo", base_url)
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

func GetShopInfoByShopID(shopID string) (ShopInfo, error) {
	var data ShopInfo
	url := fmt.Sprintf("%s/shopinfo/%s", base_url, shopID)
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
