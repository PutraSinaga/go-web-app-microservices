package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "time"
    "encoding/json"
    detail "github.com/PutraSinaga/go-web-app-microservices/src/go-microservices/detail"
)


func healthHandler(w http.ResponseWriter, r *http.Request){
    log.Println("checking health")
    response := map[string]string{
        "Status" : "ok", 
        "time" : time.Now().String(), 
    }

    json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request){
 
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    w.WriteHeader(http.StatusOK)
}

func detailHandler(w http.ResponseWriter, r *http.Request){

    fmt.Println("lihat detail..")

    hostname, err := detail.GetHostname()
    if err != nil {
        panic(err)
    }

    fmt.Println(hostname)

    IP, _ := detail.GetIP()
    fmt.Println(hostname, IP)

 
}


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/health", healthHandler)
    r.HandleFunc("/", rootHandler)
    r.HandleFunc("/detail", detailHandler)
    


    log.Fatal(http.ListenAndServe(":8000", r))
}