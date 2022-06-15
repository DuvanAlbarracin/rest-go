package main

import (
	jsonparse "encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
    ID string
}

type Book struct{
    ID string `json:"id, omitempty"`
    Name string `json:"name,omitempty"`
    Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error{
    var books []Book
    pathAbs, _ := filepath.Abs("books.json")
    content, readerr := os.ReadFile(pathAbs)
    if readerr != nil{
        log.Println("Error:", readerr)
        os.Exit(1)
    }

    errmarshal := jsonparse.Unmarshal(content, &books)
    if errmarshal != nil{
        log.Println("Error marshal:", errmarshal)
        os.Exit(1)
    }

    for _, book := range books{
        if book.ID == args.ID{
            *reply = book
            break
        }
    }
    return nil
}

func main(){
    s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterService(new(JSONServer), "")
    r := mux.NewRouter()
    r.Handle("/rpc", s)
    log.Fatal(http.ListenAndServe(":1234", r))
}
