package controllers

import (
	"go-web/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		qtty := r.FormValue("qtty")

		priceConvertedToFloat64, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Convert to float64 error.")
		}

		qttyConvertedToInt, err := strconv.Atoi(qtty)
		if err != nil {
			log.Println("Convert to int error.")
		}

		models.Insert(name, description, priceConvertedToFloat64, qttyConvertedToInt)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	pId := r.URL.Query().Get("id")
	idConvert, err := strconv.Atoi(pId)

	if err != nil {
		log.Panic(err.Error())
	}

	models.Delete(idConvert)
	http.Redirect(w, r, "/", http.StatusFound)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	pId := r.URL.Query().Get("id")
	idConvert, err := strconv.Atoi(pId)

	if err != nil {
		log.Panic(err.Error())
	}

	product := models.Edit(idConvert)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		qtty := r.FormValue("qtty")

		priceConvertedToFloat64, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Convert to float64 error. (PRICE)")
		}

		idConvertedToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Convert to int error. (ID)")
		}

		qttyConvertedToInt, err := strconv.Atoi(qtty)
		if err != nil {
			log.Println("Convert to int error. (QTTY)")
		}

		models.Update(name, description, priceConvertedToFloat64, qttyConvertedToInt, idConvertedToInt)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
