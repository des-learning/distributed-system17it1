package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(map[string]string{"message": "Hello world"})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		var m Message
		json.Unmarshal(body, &m)
		res, _ := json.Marshal(map[string]string{"message": "Hello " + m.Message})
		fmt.Fprint(w, string(res))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
