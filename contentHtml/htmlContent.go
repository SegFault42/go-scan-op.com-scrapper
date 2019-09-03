package contentHtml

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

func GetHtmlContent(link string) {

	// HTTP get request
	//resp, err := http.Get(link)
	//if err != nil {
	//panic(err)
	//}
	//defer resp.Body.Close()

	//// Convert response as string
	//dataInBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//panic(err)
	//}
	//pageContent := string(dataInBytes)

	//titleStartIndex := strings.Index(pageContent, `<div id="all" style=" display: none; ">`)
	//if titleStartIndex == -1 {
	//fmt.Println("No title element found")
	//os.Exit(0)
	//}

	//fmt.Printf("%v\n", titleStartIndex)
	//titleStartIndex += 39

	//titleEndIndex := strings.Index(pageContent, `<div id="ppp" style="">`)
	//if titleEndIndex == -1 {
	//fmt.Println("No closing tag for title found.")
	//os.Exit(0)
	//}
	//fmt.Printf("%v\n", titleEndIndex)

	//pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	//// Print out the result
	//fmt.Printf("Page title: %s\n", pageTitle)
	resp, err := soup.Get(link)
	if err != nil {
		panic(err)
	}

	doc := soup.HTMLParse(resp)

	fmt.Println(doc.Text())
	// <div id="all" style=" display: none; ">
	links := doc.Find("div", "id", "all").FindAll("img")
	for _, link := range links {
		fmt.Println(link.Attrs()["data-src"])
	}
}
