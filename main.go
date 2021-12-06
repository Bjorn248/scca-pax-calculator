package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// Holds all the class names to pax indices
	paxIndices := make(map[string]string)

	url := "https://www.solotime.info/pax/"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d \n", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Printf("+%v", paxIndices)
}
