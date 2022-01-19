package main

type ShopInfo struct {
	ID          string  `json:"id"`
	Account     string  `json:"account"`
	Rating      float32 `json:"rating"`
	Description string  `json:"description"`
}

type ProductInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Status      int    `json:"status"`
	CreateDate  string `json:"createdAt"`
	ShopID      string `json:"shopinfoId"`
}

type AccountInfo struct {
	ID            string `json:"id"`
	Name          string `json:"account_name"`
	Account       string `json:"account"`
	AccountNumber string `json:"number"`
	Amount        string `json:"amount"`
	CurrencyCode  string `json:"currency_code"`
	CreditCard    string `json:"cc_number"`
	ShopID        string `json:"shopinfoId"`
}
