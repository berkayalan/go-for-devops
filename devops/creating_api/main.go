package main

// We will use gin gonic module here: https://github.com/gin-gonic/gin

import (
	//"errors"
	"net/http"
	//"vendor/golang.org/x/net/route"
	"github.com/gin-gonic/gin"
)

type product struct {
	ProductCode  string  `json:"id"`
	ProductName  string  `json:"name"`
	MainCategory string  `json:"main"`
	SubCategory  string  `json:"sub"`
	Price        float64 `json:"price"`
	Brand        string  `json:"brand"`
	Active       bool    `json:"isactive"`
}

var products = []product{
	{ProductCode: "42178322", ProductName: "Lloyd 1.5 Ton 3 Star Inverter Split Ac", MainCategory: "Appliances", SubCategory: "Air Conditioners", Price: 2999.99, Brand: "Llold", Active: true},
	{ProductCode: "62432353", ProductName: "Skybags Brat Black 46 Cms Casual Backpack", MainCategory: "Bags & luggage", SubCategory: "Backpacks", Price: 659.99, Brand: "Skybags", Active: true},
	{ProductCode: "34253562", ProductName: "Glowic Unisex's Backpack for Men Women Boys", MainCategory: "Bags & luggage", SubCategory: "Backpacks", Price: 599.90, Brand: "Glowic", Active: true},
	{ProductCode: "53242635", ProductName: "WOW Skin Science Vitamin C Serum for Skin whitenening", MainCategory: "Beauty & health", SubCategory: "Beauty & Grooming", Price: 219.99, Brand: "WOW", Active: false},
}

func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var new_product product

	if err := c.BindJSON(&new_product); err != nil {
		return
	}

	products = append(products, new_product)
	c.JSON(http.StatusCreated, new_product)

}

func main() {

	router := gin.Default()
	router.GET("/products", GetProducts)
	router.POST("/products", CreateProduct)
	router.Run("localhost:8080")

}
