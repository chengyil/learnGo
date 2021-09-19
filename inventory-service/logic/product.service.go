package logic

import (
	"inventory-service/dto"
	"inventory-service/storage"
)

func GetProductsList() []*dto.Product {
	productDtos := make([]*dto.Product, 0)
	products := storage.Find()
	for _, product := range products {
		productDtos = append(productDtos, &dto.Product{
			ProductID:      product.ProductID,
			ProductName:    product.ProductName,
			Manufacturer:   product.Manufacturer,
			Sku:            product.Sku,
			Upc:            product.Upc,
			PricePerUnit:   product.PricePerUnit,
			QuantityOnHand: product.QuantityOnHand,
		})
	}
	return productDtos
}
