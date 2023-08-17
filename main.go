package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"unicode"
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

func countLetters(str string) int {
	count := 0
	for _, letter := range str {
		if unicode.IsLetter(letter) {
			count++
		}
	}
	return count
}

func Highlight(w http.ResponseWriter, r *http.Request) {
	chText := "<h1>"
	for scanner.Scan() {
		length := countLetters(scanner.Text())
		switch {
		case length < 4:
			chText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", scanner.Text()[:1], scanner.Text()[1:])
		case length == 4:
			chText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", scanner.Text()[:2], scanner.Text()[2:])
		case length <= 6:
			chText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", scanner.Text()[:3], scanner.Text()[3:])
		default:
			chText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", scanner.Text()[:4], scanner.Text()[4:])
		}
	}
	chText += "</h1>"
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(chText))
}

func main() {
	http.HandleFunc("/", Highlight)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
