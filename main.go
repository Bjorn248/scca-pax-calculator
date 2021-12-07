package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// getPaxMap returns a json object (as a string) of all classes
// to their pax index
func getPaxMap() string {
	// Holds all the class names to pax indices
	paxIndices := make(map[string]string)

	url := "https://www.solotime.info/pax/"

	fmt.Println("Fetching HTML from https://www.solotime.info/pax/...")
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d \n", res.StatusCode)
	}

	fmt.Println("Finished fetching HTML")

	z := html.NewTokenizer(res.Body)

	paxValue := false
	className := ""
	classPax := ""

	fmt.Println("Parsing HTML...")

	for {
		if z.Next() == html.ErrorToken {
			// Returning io.EOF indicates success.
			if z.Err() == io.EOF {
				fmt.Println("Finished parsing HTML")
				break
			} else {
				log.Fatal(z.Err())
			}
		}
		token := z.Token()
		if token.Type == html.StartTagToken {
			if token.Data == "td" {
				if string(z.Raw()) == `<td align="left" valign="middle">` {
					z.Next()
					if paxValue == false {
						className = strings.TrimSpace(string(z.Text()))
						if className == "" || className == "\n" {
							continue
						}
						paxValue = true
					} else {
						classPax = strings.Trim(strings.TrimSpace(string(z.Text())), ".")
						if classPax == "" || classPax == "\n" {
							continue
						}
						paxValue = false

						paxIndices[className] = classPax
					}
				}
			}
		}
	}

	jsonString, err := json.Marshal(paxIndices)
	if err != nil {
		log.Fatal("Could not Marhsal json", err)
	}

	return string(jsonString)
}

// generatePaxCalculator outputs a static site from go templates
// that allows you to compare autocross runs across different classes
func generateSite(paxMap string) {
	fmt.Println("Generating static site...")
}

func main() {
	paxMap := getPaxMap()
	fmt.Println(paxMap)

	generateSite(paxMap)
}
