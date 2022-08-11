package usecase

import (
	"sqlredis-server/config"
)

const (
	getColourOfProduct = "SELECT product_color FROM product WHERE product_name = ?"
)

func GetColourOfProduct(productName string) (string, error) {
	var price *string

	err := config.MysqlConnection.QueryRow(getColourOfProduct, productName).Scan(&price)

	return *price, err
}
