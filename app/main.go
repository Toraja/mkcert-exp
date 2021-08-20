package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)
	http.ListenAndServe(":8080", mux)
}

func greet(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(response{Greeting: "How are you?"})
	fmt.Fprint(w, string(resp))
}

type response struct {
	Greeting string `json:"greeting"`
}
