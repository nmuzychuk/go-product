package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestGetProducts(t *testing.T) {
	products = append(products, Product{ID: "1", Title: "Apple"})

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProducts)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code: actual: %v expected %v",
			status, http.StatusOK)
	}

	expected := `[{"id":"1","title":"Apple"}]
`
	if rr.Body.String() != expected {
		t.Errorf("Body: actual %v expected %v",
			rr.Body.String(), expected)
	}
}

func TestCreateProduct(t *testing.T) {
	productJson := `{"ID": "1", "Title": "Apple"}`
	product := strings.NewReader(productJson)

	req, err := http.NewRequest("POST", "/products", product)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProduct)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("POST /products wrong status code | got %v want %v",
			status, http.StatusOK)
	}
}
