package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request){
    log.Println("Processing resquest...")
    w.Write([]byte("OK\n"))
    log.Println("Finished processing request!")
}

func main(){
    r := mux.NewRouter()
    r.HandleFunc("/", handle)
    loggedRouter := handlers.LoggingHandler(os.Stdout, r)
    log.Fatal(http.ListenAndServe(":8000", loggedRouter))
}
