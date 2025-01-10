package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
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

type Middleware func(http.HandlerFunc) http.HandlerFunc

type AppServer struct {
	routes      map[string]http.HandlerFunc
	middlewares []Middleware
}

func NewAppServer() *AppServer {
	return &AppServer{
		routes:      make(map[string]http.HandlerFunc),
		middlewares: make([]Middleware, 0),
	}
}

func (appServer *AppServer) UseMiddleware(middleware Middleware) {
	appServer.middlewares = append(appServer.middlewares, middleware)
}

func (appServer *AppServer) AddRoute(pattern string, handlerFn http.HandlerFunc) {
	for i := len(appServer.middlewares) - 1; i >= 0; i-- {
		handlerFn = appServer.middlewares[i](handlerFn)
	}
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
	// simulating a time consuming operation to handle timeouts
	// processingTime := 2 // simulating success
	processingTime := 5 // simulating failure with timeout
	for range processingTime {
		select {
		case <-r.Context().Done():
			http.Error(w, "request timed out", http.StatusRequestTimeout)
			return
		default:
			fmt.Println("processing.....")
			time.Sleep(time.Second)
		}
	}
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

func logWrapper(handlerFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%d - %s - %s\n", r.Context().Value("trace-id"), r.Method, r.URL.Path)
		handlerFn(w, r)
	}
}

func createTimeoutWrapper(timeoutDuration time.Duration) func(handlerFn http.HandlerFunc) http.HandlerFunc {
	return func(handlerFn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			timeoutCtx, cancel := context.WithTimeout(r.Context(), timeoutDuration)
			defer cancel()
			reqWithTimeout := r.WithContext(timeoutCtx)
			handlerFn(w, reqWithTimeout)
		}
	}
}

func traceWrapper(handlerFn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		traceId := rand.Intn(2000)
		traceCtx := context.WithValue(r.Context(), "trace-id", traceId)
		handlerFn(w, r.WithContext(traceCtx))
	}
}

func main() {
	appServer := NewAppServer()
	appServer.UseMiddleware(traceWrapper)
	timeoutWrapper := createTimeoutWrapper(3 * time.Second)
	appServer.UseMiddleware(timeoutWrapper)
	appServer.UseMiddleware(logWrapper)
	appServer.AddRoute("/", IndexHandler)
	appServer.AddRoute("/products", ProductsHandler)
	appServer.AddRoute("/customers", CustomersHandler)
	if err := http.ListenAndServe(":8080", appServer); err != nil {
		log.Println(err)
	}
}
