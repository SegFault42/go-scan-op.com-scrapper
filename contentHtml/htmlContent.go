package contentHtml

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHtmlContent(link string) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}

	// Do request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
	fmt.Printf("status = %d\n", resp.StatusCode)
}
