package main

import (
	"fmt"
	"go-japscan-scrapper/contentHtml"
	"net/url"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/sirupsen/logrus"
)

// Check if link is good formatted
// Check if the link is from japscan.to
func isLinkValid(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != "" && (u.Hostname() == "www.scan-op.com" || u.Hostname() == "scan-op.com")
}

// Remove invalid links
func removeInvalidLinks(links []string) []string {
	for i, elem := range links {
		// append if link is valid
		if isLinkValid(elem) == false {
			links = append(links[:i], links[i+1:]...)
			logrus.Warn(elem, " is invalid. Skipping !")
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

func downloadManga(url string, path string) bool {
	// get content html
	htmlPage := contentHtml.GetHtmlContent(url)
	// get all images in []string
	imageUrl := contentHtml.GetLinks(htmlPage)
	if imageUrl == nil {
		return false
	}

	for _, elem := range imageUrl {
		// get image number
		split := strings.Split(elem, "/")
		imageName := path + "/" + split[(len(split)-1)]
		imageName = strings.TrimSpace(imageName)

		// Download in folder
		err := contentHtml.DownloadFile(strings.TrimSpace(elem), imageName)
		if err != nil {
			logrus.Errorln(err)
		} else {
			logrus.Info("\033[1;1;33m", imageName, " \033[1;1;32mDownload success !\033[0m")
		}
	}

	return true
}

func createFolders(link string) string {

	split := strings.Split(link, "/")
	path := split[4] + "/" + split[5]
	os.MkdirAll(path, os.ModePerm)

	return path
}

func main() {
	// Setup logrus
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	logrus.SetFormatter(Formatter)

	// get all links from arg
	links, err := parseFlag()
	if err != nil {
		return
	}

	// remove all invalid links
	links = removeInvalidLinks(links)

	if len(links) == 0 {
		logrus.Errorln("\033[1;1;31mNo valid links found\033[0m")
		return
	}

	// iter on each link given in arg
	for _, elem := range links {
		logrus.Info("\033[1;1;33mDownloading :\033[0m", "\n")

		path := createFolders(elem)

		if downloadManga(elem, path) == true {
			fmt.Println("")
			logrus.Info("\033[1;1;33m", elem, "\033[1;1;32m Finish !\033[0m")
			fmt.Println("")
		}
	}
}
