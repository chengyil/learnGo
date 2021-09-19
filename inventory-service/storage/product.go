package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Product struct {
	ProductID      int    `json:"productID"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var (
	products = struct {
		sync.RWMutex
		m map[int]*Product
	}{m: make(map[int]*Product)}
)

func Find() []*Product {
	prods := make([]*Product, 0)
	for _, product := range products.m {
		prods = append(prods, product)
	}
	return prods
}

func Init() {
	filename := "./storage/products.json"
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		log.Fatal("Unable to load products", err)
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Unable to load products", err)
	}

	loaded := make([]*Product, 0)

	if err = json.Unmarshal(content, &loaded); err != nil {
		log.Fatal(err)
	}

	products.Lock()
	defer products.Unlock()
	for _, product := range loaded {
		products.m[product.ProductID] = product
	}
}
