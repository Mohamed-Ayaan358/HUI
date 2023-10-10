package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Struct to represent a node in the parse tree

type Node struct {
	Name     string
	Content  map[string]string
	Children []*Node
}

var css []string
var html []string

// myDictionary := make(map[string]int)

func readHtmlFromFile(fileName string) (string, error) {

	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func indexOf(slice []string, target string) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1 // Return -1 if the target is not found in the slice
}

func main() {
	fileName := "index.html"
	text, err := readHtmlFromFile(fileName)

	Parser(text)
	if err != nil {
		log.Fatal(err)
	}

}

func CSSextractor(CSSlines []string, CSSnum int) map[string]string {
	CSSselector := make(map[string]string)

	for _, lin := range CSSlines[CSSnum:] {
		trimmedl := strings.TrimSpace(lin)
		css = append(css, trimmedl)
		if "</style>" == trimmedl {
			break
		}
	}

	css = css[1 : len(css)-1]

	re1 := regexp.MustCompile(`\w+ {`)
	re2 := regexp.MustCompile(`}`)

	start := 0
	end := 0

	// var match string
	for n, val := range css {
		res1 := re1.FindString(val)
		res2 := re2.FindString(val)

		if res1 != "" {
			start = n
		}
		if res2 != "" {
			end = n
			tag := css[start][0 : len(css[start])-2]
			CSSselector[tag] = strings.Join(css[start+1:end], "")
		}

	}
	return CSSselector

}

func HTMLextractor(HTML string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(HTML))
	if err != nil {
		log.Fatal(err)
	}

	bodySelection := doc.Find("body")

	// Check if the body element exists
	if bodySelection.Length() == 0 {
		fmt.Println("No <body> element found in the HTML.")
		return
	}

	// Get the HTML content of the body element
	bodyHTML, _ := bodySelection.Html()

	// Parse the body content
	lines := strings.Split(bodyHTML, "\n")
	fmt.Println(lines)

}

func Parser(text string) {
	lines := strings.Split(text, "\n")

	for num, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if "<style>" == trimmedLine {
			CSSselector := CSSextractor(lines, num)
			for key, value := range CSSselector {
				val := strings.Split(value, ";")
				fmt.Println(key, val)
			}
		}

		if "<body>" == trimmedLine {
			HTMLextractor(text)
		}

		// fmt.Println(html)
	}

}
