package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          string  `json: "id"`
	Name        string  `json: "name"`
	Price       float64 `json: "price"`
	Inventory   int     `json: "inventory"`
	Code        string  `json: "code"`
	IsPublished bool    `json: "isPublished"`
	CreatedAt   string  `json: "createdAt"`
}

func main() {
	app := gin.Default()

	data := []byte(`[{"id":"Claudio1", "name":"caneta", "price":1.50, "invetory":30, "code":"v12", "isPublished": false, "createdAt":"20/12/2022"},{"id":"Claudio2", "name":"caneta", "price":1.50, "invetory":30, "code":"v12", "isPublished": false, "createdAt":"20/12/2022"}]`)
	err := os.WriteFile("products.json", data, 0644)
	if err != nil {
		panic(err)
	}

	var product []Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		panic(err)
	}

	app.GET("/person", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, product)
	})
	app.Run(":8080")
}
