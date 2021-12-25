package main

import (
	"AlliasolDevis/handler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/", handler.DevisHome)
	r.HandleFunc("/devis", handler.DevisHandler).Methods("POST", http.MethodOptions)
	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
