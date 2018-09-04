package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

// PART 13 - https://www.youtube.com/watch?v=dySAX8VZ2TU
// PART 15 - https://www.youtube.com/watch?v=2umnbZxdu-s 


type Sitemapindex struct {
	// SLICE of Locations - string type
	// "sitemap>loc" - 'loc' tag underneath the 'sitemap' tag
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles [] string `xml:"url>news>title"`
	Keywords [] string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}
/*
	Get information from the Internet
*/
func main() {
	var s Sitemapindex
	var n News
	news_map := make(map[string]NewsMap)

	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	
	//string_body := string(bytes)
	//fmt.Println("Printing string_body: \n", string_body)
	
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		// don't care about the value
		for i, _ := range n.Titles {
			//fmt.Println(i)
			news_map[n.Titles[i]] = NewsMap{n.Keywords[i], n.Locations[i]}
		}
	}

	for idx, data := range news_map{
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}
