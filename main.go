package main

import (
	"log"
	"net/http"
)

func Highlight(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello."))
}

func main() {
	http.HandleFunc("/", Highlight)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
