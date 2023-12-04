package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/",Index)
	
	fs := http.FileServer(http.Dir("./uploads"))
	http.Handle("/pix.gif", http.StripPrefix("/", fs))

	http.ListenAndServe(":5647",nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprintf(w,"Online")
}

func enableCors(w *http.ResponseWriter) {
(*w).Header().Set("Access-Control-Allow-Origin", "https://xfy.onrender.com/")
}
