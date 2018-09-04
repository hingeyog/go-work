package main

/*
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>
*/

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

// part 1 - https://www.youtube.com/watch?v=ccANcNk8Dac
// part 2 - https://www.youtube.com/watch?v=-PATP8IZq5A

type SitemapIndex struct {
	// [5] <type> == ARRAY --> Fixed size
	// [ ] <type> == SLICE --> Not fixed size
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

// Print location as strings
func (l Location) String() string {
	// Sprintf formats according to a format specifier and returns the resulting string.
	return fmt.Sprintf(l.Loc)
}

/*
	Get information from the Internet
*/
func main() {
	// use '_' to define a variable which you don't intend to use
	// '_' here is used for 'err' response
	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println("Printing string_body: \n", string_body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println("\n\nPrinting all the Locations: \n", s.Locations)

	//fmt.Println("\n\nPrinting all the Locations: \n", s.Locations)

	fmt.Println("\n\nPrinting all the Locations on separate formatted lines:")	
	// 'range' function iterates over the structure and returns the index value and whatever actual value is
	for _, Location := range s.Locations {
			fmt.Printf("%s\n", Location)
	}

}