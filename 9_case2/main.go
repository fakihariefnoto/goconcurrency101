package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var client http.Client = http.Client{}

const base_url = "https://61dd5a81f60e8f001766874a.mockapi.io"

func Init() {
	shops, err := GetAllShop()

	if err != nil {
		panic(err)
	}

	fetchStandartNoConcurrency(shops)
	fetchWithConcurrency(shops)
}

func fetchStandartNoConcurrency(shops []ShopInfo) {
	defer caculateTime(time.Now(), "standart w/o concurrency")

	shop := shops[0]

	products, err := GetAllProduct(shop.ID)

	if err != nil {
		panic(err)
	}

	accounts, err := GetAllAccount(shop.ID)

	if err != nil {
		panic(err)
	}

	for _, product := range products {
		GetProductInfoByProductID(product.ShopID, product.ID)
	}

	for _, account := range accounts {
		GetAccountInfoByAccountID(account.ShopID, account.ID)
	}
}

func fetchWithConcurrency(shops []ShopInfo) {
	defer caculateTime(time.Now(), "with concurrency")

	wgShop := sync.WaitGroup{}
	wgShop.Add(2)
	shop := shops[0]

	go func(shopid string) {
		products, err := GetAllProduct(shopid)
		wgProduct := sync.WaitGroup{}
		wgProduct.Add(len(products))

		if err != nil {
			panic(err)
		}

		for _, product := range products {
			go func(shopid, productid string) {
				GetProductInfoByProductID(shopid, productid)
				wgProduct.Done()
			}(product.ShopID, product.ID)
		}

		wgProduct.Wait()
		wgShop.Done()
	}(shop.ID)

	go func(shopid string) {
		accounts, err := GetAllAccount(shopid)
		wgAccount := sync.WaitGroup{}
		wgAccount.Add(len(accounts))

		if err != nil {
			panic(err)
		}

		for _, account := range accounts {
			go func(shopid, accountid string) {
				GetAccountInfoByAccountID(shopid, accountid)
				wgAccount.Done()
			}(account.ShopID, account.ID)
		}

		wgAccount.Wait()
		wgShop.Done()
	}(shop.ID)

	wgShop.Wait()
}

func caculateTime(start time.Time, module string) {
	fmt.Printf("%s took %v\n", module, time.Since(start))
}
