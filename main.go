package main

import (
	"docker-watcher/container"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func setupRouter(router *mux.Router) {
	router.Methods("GET").Path("/test").HandlerFunc(GetTest)
}

func GetTest(w http.ResponseWriter, r *http.Request) {
	log.Println("Test GET called")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Test"))
}

// Resource: https://medium.com/backendarmy/controlling-the-docker-engine-in-go-d25fc0fe2c45
func main() {
	//router := mux.NewRouter()
	//setupRouter(router)
	//log.Fatal(http.ListenAndServe(":8080", router))
	err := container.ListContainer()
	if err != nil {
		return
	}
}
