package main

import (
	"example/db"
	"example/service/product"
	"example/store"
	"example/store/pg"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()
	db, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() { db.Close() }()

	storeProduct := pg.NewProductImplDB(db)
	productService := product.NewProductService(db, storeProduct)

	r.POST("/products", func(c *gin.Context) {
		data := &store.ProductData{}
		if err := c.Bind(&data); err != nil {
			c.Error(err)
		}

		if err := productService.CreateProduct(c, data); err != nil {
			c.Error(err)
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
