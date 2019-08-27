package main

import (
	"fmt"
	"os"
)

// Check if link is good formatted
// Check if the link is from japscan.to
func isLinkValid() {

}

// Get all valid links
func readArgs() [][]string {

	argsSize := len(os.Args[1:])

	links := make([][]string, argsSize)

	for i, elem := range os.Args[1:] {
		// append if link is valid
		if isLinkValid(links[i]) == true {
			links[i] = append(links[i], elem)
		}
	}

	return links
}

func main() {
	links := readArgs()
	fmt.Printf("%v\n", links)

}

//req, err := http.NewRequest("GET", "https://www.japscan.to/lecture-en-ligne/shingeki-no-kyojin/101/4.html", nil)
//if err != nil {
//// handle err
//}
//req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
//req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
//req.Header.Set("Accept-Language", "en-US,en;q=0.5")
//req.Header.Set("Referer", "https://www.japscan.to/lecture-en-ligne/shingeki-no-kyojin/101/4.html")
//req.Header.Set("Connection", "keep-alive")
//req.Header.Set("Cookie", "__cfduid=db1d20b3eab825e6868ecba2c86147dd31566657424; session=ki6n1ocqe1uinbvd61nb25arh3; cf_clearance=f194d620c9409857206b1f20e014726d546ac8d9-1566895356-14400-150")
//req.Header.Set("Upgrade-Insecure-Requests", "1")
//req.Header.Set("Te", "Trailers")

//resp, err := http.DefaultClient.Do(req)
//if err != nil {
//// handle err
//}
//defer resp.Body.Close()

//file, err := os.Create("test.html")
//if err != nil {
//panic(err)
//}
//resp.Write(file)
//fmt.Printf("len = %d\n", resp.StatusCode)
