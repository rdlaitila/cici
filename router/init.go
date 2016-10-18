package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	log.Println("Init: router")

	r := mux.NewRouter()
	r.HandleFunc("/", nil)
	r.HandleFunc("/products", nil)
	r.HandleFunc("/articles", nil)
	http.Handle("/", r)

	GlobalRouter = r
}
