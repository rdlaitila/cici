package node

import (
	"cici/router"
	"log"
)

func init() {
	log.Println("Init:", resourceBase)

	router.GlobalRouter.HandleFunc(resourceBase, getNode).Methods("GET")
}
