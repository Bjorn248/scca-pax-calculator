package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"

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
				z.Next()
				if paxValue == false {
					className = strings.TrimSpace(string(z.Text()))
					if className == "" || className == "\n" || strings.Contains(className, "Results") {
						continue
					}
					paxValue = true
				} else {
					classPax = strings.Trim(strings.TrimSpace(string(z.Text())), ".")
					if classPax == "" || classPax == "\n" || strings.Contains(classPax, "Results") {
						continue
					}
					paxValue = false

					// Only add entries where the PAX value is a valid number
					if _, err := strconv.ParseFloat(classPax, 64); err == nil {
						paxIndices[className] = classPax
					}
				}
			}
		}
	}

	jsonString, err := json.MarshalIndent(paxIndices, "", "  ")
	if err != nil {
		log.Fatal("Could not Marhsal json", err)
	}

	return string(jsonString)
}

// generatePaxCalculator outputs a static site from go templates
// that allows you to compare autocross runs across different classes
func generateSite(paxMap string) {
	fmt.Println("Generating static site...")
	tpl, err := template.ParseFiles("templates/common.js.tmpl")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	outFile, err := os.Create("src/common.js")
	if err != nil {
		log.Fatal("Could not create file", err)
	}

	err = tpl.Execute(outFile, paxMap)
	if err != nil {
		log.Fatal("Could not execute template", err)
	}

	outFile.Close()
	fmt.Println("Done!")
}

func main() {
	paxMap := getPaxMap()
	generateSite(paxMap)
}
