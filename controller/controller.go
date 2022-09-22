package controller

import (
	"docker-watcher/dockerHost"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetLsForContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("GetLsForContainer GET called")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	containerName := vars["containerName"]
	ls := dockerHost.GetLsForContainer(containerName)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(ls))
}

func GetImageVersions(w http.ResponseWriter, r *http.Request) {
	log.Println("GetImageVersions GET called")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	imageName := vars["imageName"]
	versions := dockerHost.GetImageVersions(imageName)
	jsonResponse, err := json.Marshal(versions)
	if err != nil {
		log.Println("Unable to marshal json", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}
