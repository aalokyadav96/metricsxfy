package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",Index)
	http.ListenAndServe(":5647",nil)
}


func Index(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		fmt.Fprintf(w,"200")
		return
	}
	fmt.Fprintf(w,"Online")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "https://xfy.onrender.com")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
