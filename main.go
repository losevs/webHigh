package main

import (
	"bufio"
	"fmt"
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
	chText := "<h1>"
	for scanner.Scan() {
		chText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", scanner.Text()[:1], scanner.Text()[1:])
	}
	chText += "</h1>"
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(chText))
}

func main() {
	http.HandleFunc("/", Highlight)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
