package logic

import "inventory-service/dto"

func GetProductsList() []*dto.Product {
	return []*dto.Product{
		{
			ProductID:      1,
			Manufacturer:   "IBM",
			Sku:            "sku",
			Upc:            "Upc",
			PricePerUnit:   "1.1",
			QuantityOnHand: 10,
			ProductName:    "Server",
		},
		{
			ProductID:      2,
			Manufacturer:   "IBM",
			Sku:            "sku",
			Upc:            "Upc",
			PricePerUnit:   "1.1",
			QuantityOnHand: 10,
			ProductName:    "Laptop",
		},
	}
}

// ProductID      int    `json:"productId"`
// Manufacturer   string `json:"manufacturer"`
// Sku            string `json:"sku"`
// Upc            string `json:"upc"`
// PricePerUnit   string `json:"pricePerUnit"`
// QuantityOnHand int    `json:"quantityOnHand"`
// ProductName    string `json:"productName"`
