package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main(){
    router := httprouter.New()
    router.ServeFiles("/static/*filepath",
        http.Dir("Documents/Go/rest-go/part2/static"))
    log.Fatal(http.ListenAndServe(":8000", router))
}
