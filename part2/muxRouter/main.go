package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category is : %v\n", vars["category"])
    fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main(){
    router := mux.NewRouter()
    router.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("articleRoute")

    // prints the path of the named route
    url, err := router.Get("articleRoute").URL("category", "books", "id", "123")
    if err != nil{
        panic(err)
    }
    fmt.Printf(url.Path)

    server := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }
    log.Fatal(server.ListenAndServe())
}
