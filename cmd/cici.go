package main

import (
	_ "cici/resources/core/filesystem/file"
	_ "cici/resources/core/filesystem/node"
	_ "cici/router"
	"log"
	"net/http"
)

func main() {
	log.Println("cici started")
	http.ListenAndServe("0.0.0.0:8181", nil)
}
