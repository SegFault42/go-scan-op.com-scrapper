package contentHtml

import (
	"io"
	"net/http"
	"os"

	"github.com/anaskhan96/soup"
	"github.com/sirupsen/logrus"
)

func DownloadFile(url string, filePath string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return err
}

func GetLinks(html string) []string {
	// Takes the HTML string as an argument, returns a pointer to the DOM constructed
	doc := soup.HTMLParse(html)

	// Find all images
	parse := doc.Find("div", "id", "all")
	if parse.Error != nil {
		logrus.Error("Unable to parse link")
		return nil
	}
	links := parse.FindAll("img")

	var s []string
	// store all links in []string
	for _, link := range links {
		s = append(s, link.Attrs()["data-src"])
	}

	return s
}

func GetHtmlContent(link string) string {

	// Get request
	resp, err := soup.Get(link)
	if err != nil {
		panic(err)
	}

	return resp
}
