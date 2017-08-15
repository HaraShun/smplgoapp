package main

import (
        "fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Path: %s", r.URL.Path[1:])
        fmt.Println("Request: ", r.URL.Path[1:])
}

func main() {
        http.HandleFunc("/",handler)
        fmt.Println("Server Ready...")
        http.ListenAndServe(":8080",nil)
}
