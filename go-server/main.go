package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sutin1234/react-go-api/config"
)

type TIME struct {
	Time string `json:"time"`
}

func main() {

	cfg, _ := config.LoadConfig()
	r := mux.NewRouter()

	r.HandleFunc(GetApiPath(cfg, "getTime"), IndexHandleFunc).Methods(http.MethodGet)

	err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), r)
	if err != nil {
		log.Fatalf("server failed %v", err)
	}
	fmt.Printf("server running on %v:%v/%v/%v", cfg.BaseUrl, cfg.Port, cfg.ApiPrefix, cfg.ApiVersion)
}

func IndexHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	timeNow := TIME{fmt.Sprint(now)}
	bytes, _ := json.Marshal(timeNow)

	_, _ = w.Write(bytes)
}

func GetApiPath(cfg *config.Config, path string) string {
	absPath := fmt.Sprintf("/%v/%v/%v", cfg.ApiPrefix, cfg.ApiVersion, path)
	fmt.Printf("Absolute path: %v", absPath)
	return absPath
}
