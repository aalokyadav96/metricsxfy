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
	fmt.Fprintf(w,"Online")
}
