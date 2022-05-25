package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

type UUID struct{}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request){
    if r.URL.Path == "/"{
        genRandomUUID(w, r)
        return
    }
    http.NotFound(w, r)
    return
}

func genRandomUUID(w http.ResponseWriter, r *http.Request){
    c := 10
    b := make([]byte, c)
    _, err := rand.Read(b)
    if err != nil{
        panic(err)
    }
    fmt.Fprintf(w, fmt.Sprintf("%x", b) + "\n")
}

func main(){
    mux := &UUID{}
    port := ":8000"
    fmt.Printf("Starting server on port %s", port)
    log.Fatal(http.ListenAndServe(port, mux))
}
