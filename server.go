package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)


/* web endpoints */
func (a *App) GetHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request key: %v\n", "test")
}



func (a *App) PutHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Delete key: %v to %v\n", "test")
}


func (a *App) initRoutes() {
	a.router.HandleFunc("/put/{key}", a.PutHandler).Methods("POST")
	a.router.HandleFunc("/get/{key}", a.GetHandler).Methods("POST")
	a.router.HandleFunc("/delete/{key}", a.DeleteHandler).Methods("DELETE")
}


