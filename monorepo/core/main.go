package main

import (
    "fmt"
    "net/http"
    "os"
    "strings"
)

func main() {
    serverPort := os.Getenv("PORT");
    serviceName := os.Getenv("SERVICE_NAME");

    fmt.Printf("[%s]: is running on port %s\n", strings.ToUpper(serviceName), serverPort)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("This is the "+ serviceName +" service"));
    });

    err := http.ListenAndServe(":"+serverPort, nil);

    if err != nil {
        fmt.Printf("Error starting server: %v\n", err);
    }
}
