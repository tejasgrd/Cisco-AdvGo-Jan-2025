package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

var products []Product = []Product{
	{101, "Pen", 10},
	{102, "Pencil", 5},
	{103, "Marker", 50},
}

type AppServer struct {
	routes map[string]func(w http.ResponseWriter, r *http.Request)
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
}

func (appServer *AppServer) AddRoute(pattern string, handlerFn func(w http.ResponseWriter, r *http.Request)) {
	appServer.routes[pattern] = handlerFn
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, exists := appServer.routes[r.URL.Path]; exists {
		handler(w, r)
		return
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}

/* *************** App Specific ************** */
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!\n")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	case http.MethodPost:
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}
		products = append(products, newProduct)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The list of customers will be served!\n")
}

func main() {
	appServer := NewAppServer()
	appServer.AddRoute("/", IndexHandler)
	appServer.AddRoute("/products", ProductsHandler)
	appServer.AddRoute("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Println(err)
	}
}
