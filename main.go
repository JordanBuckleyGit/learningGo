package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}
    
func main() {
    http.HandleFunc("/", helloWorld)
    fmt.Println("Server starting on port 8080...")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:",err)
    }
}
