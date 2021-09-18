package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Product struct {
	ProductID      int    `json:"id"`
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

func Init() {
	filename := "products.json"
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		log.Fatal("Unable to load products")
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Unable to load products")
	}

	products := make([]*Product, 0)

	err = json.Unmarshal(content, &products)

	if err != nil {
		log.Fatal(err)
	}

}
