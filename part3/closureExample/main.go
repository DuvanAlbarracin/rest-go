package main

import (
    "fmt"
)

func generator() func() int{
    i := 0
    return func() int{
        i++
        return i
    }
}

func main (){
    numberGen := generator()
    for i := 0; i < 4; i++{
        fmt.Print(numberGen(), "\t")
    }
}
