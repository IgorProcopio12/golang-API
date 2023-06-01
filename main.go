package main

import (
	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int32
	Code        string
	Description string
	Price       float64
}

var products []Product

func main() {
	e := echo.New()
	e.GET("/products", getProducts)
	e.POST("/products", createProduct)
	e.Logger.Fatal(e.Start(":8080"))
}

func getProducts(c echo.Context) error {
	return c.JSON(200, products)
}

func createProduct(c echo.Context) error {
	product := new(Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	products = append(products, *product)
	return c.JSON(200, products)
}
