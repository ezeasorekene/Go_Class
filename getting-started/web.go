package main

import (
  "fmt"
  "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request)  {
  fmt.Fprint(w, "<h1>This is my first Go web page</h1>")
}

func main(){
  http.HandleFunc("/", handlerFunc)
  http.ListenAndServe("127.0.0.1:9000", nil)
}
