package models

import (
	"go-web/database"
	"log"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Qtty        int
}

func GetAllProducts() []Product {
	db := database.Connect()

	sql := "SELECT * FROM products ORDER BY id ASC"

	selectAll, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	var productList []Product

	for selectAll.Next() {
		product := Product{}

		err = selectAll.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Qtty)
		if err != nil {
			panic(err.Error())
		}

		productList = append(productList, product)
	}

	defer db.Close()

	return productList
}

func Insert(name, description string, price float64, qtty int) {
	db := database.Connect()

	sql := "INSERT INTO products (name, description, price, qtty) values ($1,$2,$3,$4)"

	insertData, err := db.Prepare(sql)

	if err != nil {
		log.Panic(err.Error())
	}

	insertData.Exec(name, description, price, qtty)

	defer db.Close()
}

func Delete(id int) {
	db := database.Connect()

	sql := "DELETE from products WHERE id = $1"

	insertData, err := db.Prepare(sql)

	if err != nil {
		log.Panic(err.Error())
	}

	insertData.Exec(id)

	defer db.Close()
}

func Edit(id int) Product {
	db := database.Connect()

	editProduct, err := db.Query("SELECT * from products WHERE id = $1", id)

	if err != nil {
		log.Panic(err.Error())
	}

	ProductToUpdate := Product{}

	for editProduct.Next() {
		var id, qtty int
		var name, description string
		var preco float64

		err = editProduct.Scan(&id, &name, &description, &preco, &qtty)
		if err != nil {
			log.Panic(err.Error())
		}

		ProductToUpdate.Id = id
		ProductToUpdate.Name = name
		ProductToUpdate.Description = description
		ProductToUpdate.Price = preco
		ProductToUpdate.Qtty = qtty
	}

	defer db.Close()
	return ProductToUpdate
}

func Update(name, description string, price float64, qtty, id int) {
	db := database.Connect()

	sql := "UPDATE products set name=$1, description=$2, price=$3, qtty=$4 WHERE id = $5"

	updateData, err := db.Prepare(sql)

	if err != nil {
		log.Panic(err.Error())
	}

	updateData.Exec(name, description, price, qtty, id)

	defer db.Close()
}
