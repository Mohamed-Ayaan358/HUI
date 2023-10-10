// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// // Struct to represent a node in the parse tree
// type Node struct {
// 	Name     string
// 	Content  map[string]string
// 	Children []*Node
// }

// var html []string

// func readHtmlFromFile(fileName string) (string, error) {
// 	bs, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(bs), nil
// }

// func main() {
// 	fileName := "index.html"
// 	text, err := readHtmlFromFile(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create a new document using goquery
// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Find the body element
// 	bodySelection := doc.Find("body")

// 	// Check if the body element exists
// 	if bodySelection.Length() == 0 {
// 		fmt.Println("No <body> element found in the HTML.")
// 		return
// 	}

// 	// Get the HTML content of the body element
// 	bodyHTML, _ := bodySelection.Html()

// 	// Parse the body content
// 	Parser(bodyHTML)
// }

// func Parser(text string) {
// 	// Parse the <body> content here
// 	lines := strings.Split(text, "\n")
// 	fmt.Println(lines)

// }

// func createParseTree(selection *goquery.Selection) *Node {
// 	node := &Node{
// 		Name:    selection.Nodes[0].Data,
// 		Content: make(map[string]string),
// 	}

// 	selection.Children().Each(func(_ int, child *goquery.Selection) {
// 		childNode := createParseTree(child)
// 		node.Children = append(node.Children, childNode)
// 	})

// 	if len(selection.Nodes) > 0 {
// 		content := selection.Nodes[0].FirstChild
// 		if content != nil {
// 			node.Content["text"] = content.Data
// 		}
// 	}

// 	return node
// }
