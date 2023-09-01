package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type FOO struct {
	Foo string `json:"foo"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", IndexHandleFunc).Methods(http.MethodGet)

	err := http.ListenAndServe(":3001", r)
	if err != nil {
		fmt.Printf("server failed %v", err)
		return
	}
	fmt.Printf("server running on http://localhost:3001")
}

func IndexHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	foo := FOO{fmt.Sprintf("foo %v", time.Now().UTC().String())}
	bytes, _ := json.Marshal(foo)

	_, _ = w.Write(bytes)
}
