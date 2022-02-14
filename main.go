package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msg := "REQUEST FROM SAMPLE REMOTE PROVIDER ( PORT : 8080 )"
	response := map[string]string{
		"message": msg,
	}
	m, _ := json.Marshal(response)
	fmt.Println("MESSAGE SENT TO CFA PRIVATE")
	w.Write(m)
}

func main() {
	http.HandleFunc("/", fetchData)
	http.ListenAndServe(":8080", nil)
}
