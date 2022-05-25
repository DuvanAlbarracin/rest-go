package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request){
    queryParams := r.URL.Query()
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Got parameter for id: %s\n", queryParams.Get("id"))
    fmt.Fprintf(w, "Got parameter for category: %s\n", queryParams.Get("category"))
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/articles", QueryHandler)
    server := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    log.Fatal(server.ListenAndServe())
}

