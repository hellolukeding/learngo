package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/index", index)
	http.HandleFunc("/", indexhtml)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}

/**pocess*/

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	indexData := indexdata{
		Username: "lukeding",
		Password: "123456",
	}
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexhtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// t := template.New("index.html")
	path, _ := os.Getwd()
	w.Write([]byte(path))
	// t, _ = t.ParseFiles(path + "index.html")
	// t.Execute(w, nil)

}

/**mockdata*/

type indexdata struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
