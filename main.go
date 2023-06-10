/*package main

import (
	"encoding/json"
	"fmt"
)

// main function
func main() {

	// defining a map
	var result map[string][]string

	// string json
	jsonString := `{
		"47": ["47"],
		"цветок": ["цвето"]
	  }`

	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		// print out if error is not nil
		fmt.Println(err)
	}

	// printing details of map
	// iterate through the map
	for _, value := range result {
		fmt.Println(value[0])
	}
}*/

package main

import (
	"disp_bot/mail"
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func htmlToString(htmlString string) string {
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		panic(err)
	}

	var clearText string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			clearText += n.Data + "\n"
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
			if n.Type == html.ElementNode && n.Data == "p" {
				clearText += "\n"
			}
		}
	}
	f(doc)
	return clearText
}

func main() {
	var Mail mail.Mail

	Mail.Connect()
	defer Mail.Close()

	htmlStrings := Mail.GetEmail()
	htmlString := htmlStrings[len(htmlStrings)-1]

	fmt.Println("Subject: ", htmlToString(htmlString.Subject))
	fmt.Println("Body: ", htmlToString(htmlString.Body))

}
