package main

import (
	"docker-watcher/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.Methods("GET").Path("/version/{imageName}").HandlerFunc(controller.GetImageVersions)
	router.Methods("GET").Path("/ls/{containerName}").HandlerFunc(controller.GetLsForContainer)
}

func main() {
	router := mux.NewRouter()
	setupRouter(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
