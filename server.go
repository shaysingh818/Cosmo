package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
)


/* web endpoints */
func (a *App) GetHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	key := vars["key"]

	content := a.RetrieveKey([]byte(key))

	// create response
	response := make(map[string]string)
	response["key"]  = key
	response["dateUpload"] = "random date"
	response["currentTime"] = "current time"
	response["content"] = string(content)

	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Write(jsonResp)
}



func (a *App) PutHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// get key from response
	vars := mux.Vars(r)
	key := vars["key"]

	fileData, err := ioutil.ReadAll(r.Body)
	CheckError(err)

	a.PutKey([]byte(key), fileData)

	// generate response
	response := make(map[string]string)
	response["message"] = "put key"
	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Write(jsonResp)

}


func (a *App) DeleteHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// get key from response
	vars := mux.Vars(r)
	key := vars["key"]

	result := a.DeleteKey([]byte(key))
	if result != true {
		fmt.Println("Something went wrong again")
	}

	response := make(map[string]string)
	response["message"] = "deleted key!"
	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Write(jsonResp)

}


func (a *App) ListAllHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result := a.ViewKeys()

	response := make(map[string]string)
	response["message"] = "success"
	response["keys"] = result
	jsonResp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Something went wrong")
	}

	w.Write(jsonResp)
}


func (a *App) initRoutes() {
	a.router.HandleFunc("/list", a.ListAllHandler).Methods("GET")
	a.router.HandleFunc("/get/{key}", a.GetHandler).Methods("GET")
	a.router.HandleFunc("/put/{key}", a.PutHandler).Methods("POST")
	a.router.HandleFunc("/delete/{key}", a.DeleteHandler).Methods("DELETE")
}


