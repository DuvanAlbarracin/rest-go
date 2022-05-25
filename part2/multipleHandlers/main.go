package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main(){
    multipleMux := http.NewServeMux()

    multipleMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, rand.Float32())
    })

    multipleMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, rand.Intn(100))
    })

    port := ":8000"
    fmt.Printf("Starting server on port %s", port)
    log.Fatal(http.ListenAndServe(port, multipleMux))
}
