package main

import (
	"fmt"
	"net/url"
	"os"

	"go-japscan-scrapper/contentHtml"

	"github.com/akamensky/argparse"
)

// Check if link is good formatted
// Check if the link is from japscan.to
func isLinkValid(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != "" && u.Hostname() == "www.scan-op.com"
}

// Remove invalid links
func removeInvalidLinks(links []string) []string {
	for i, elem := range links {
		// append if link is valid
		if isLinkValid(elem) == false {
			links = append(links[:i], links[i+1:]...)
		}
	}

	return links
}

func parseFlag() ([]string, error) {
	parser := argparse.NewParser("print", "Prints provided string to stdout")
	s := parser.List("l", "link", &argparse.Options{Required: true, Help: "link from scan-op.com"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	return *s, err
}

func main() {
	// get all links from arg
	links, err := parseFlag()
	if err != nil {
		return
	}

	// remove all invalid links
	links = removeInvalidLinks(links)

	if len(links) == 0 {
		fmt.Println("No valid links founds")
		return
	}

	for _, elem := range links {
		contentHtml.GetHtmlContent(elem)
	}

}
