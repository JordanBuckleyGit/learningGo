package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", helloWorld)
    fmt.Println("Server starting on port 8080...")

    err := http.ListenAndServer(":8080",nil)
