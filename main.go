package main

import (
	"docker-watcher/dockerHost"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.Methods("GET").Path("/test").HandlerFunc(GetTest)
	router.Methods("GET").Path("/images").HandlerFunc(GetImages)
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	log.Println("Test GET called")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Test"))
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	log.Println("Images GET called")
	w.Header().Set("Content-Type", "application/json")
	images := dockerHost.GetImages()
	jsonData, err := json.Marshal(images)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonData)
}

func main() {
	router := mux.NewRouter()
	setupRouter(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
