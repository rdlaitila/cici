package file

import (
	"cici/router"
	"log"
)

func init() {
	log.Println("Init:", resourceBase)

	router.GlobalRouter.HandleFunc(resourceBase, get).Methods("GET")
}
