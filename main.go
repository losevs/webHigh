package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
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
	chText := "<div style='width: 600px; margin: auto; padding-top: 20px'><h2>"
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
	chText += "</h2></div><br><div style='width: 600px; margin: auto; padding-top: 20px; text-align: center'><form action='/hand'><label for='text'>Text</label><br><textarea style='height: 190px; width: 558px;' id='text' name='text' value=''></textarea><br><input type='submit' value='Submit'></form></div>"
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(chText))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	textSlice := strings.Split(text, " ")
	NewText := "<div style='width: 600px; margin: auto; padding-top: 20px'><h2>"
	for _, word := range textSlice {
		length := len(word)
		switch {
		case length < 4:
			NewText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", word[:1], word[1:])
		case length == 4:
			NewText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", word[:2], word[2:])
		case length <= 6:
			NewText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", word[:3], word[3:])
		default:
			NewText += fmt.Sprintf("<span style='color: black'>%s</span><span style='color: gray'>%s</span> ", word[:4], word[4:])
		}
	}
	NewText += "</h2></div><br><div style='width: 600px; margin: auto; padding-top: 20px; text-align: center'><form action='/hand'><label for='text'>Text</label><br><textarea style='height: 190px; width: 558px;' id='text' name='text' value=''></textarea><br><input type='submit' value='Submit'></form></div>"
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(NewText))
}

func main() {
	http.HandleFunc("/", Highlight)
	http.HandleFunc("/hand", Handler)
	log.Fatalln(http.ListenAndServe(":80", nil))
}
