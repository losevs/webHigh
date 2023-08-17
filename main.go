package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

var file, err = os.Open("text.txt")

func init() {
	if err != nil {
		log.Fatalln(err)
	}
}

var scanner = bufio.NewScanner(file)

func init() {
	scanner.Split(bufio.ScanWords)
}

func Highlight(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello."))
}

func main() {
	http.HandleFunc("/", Highlight)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
