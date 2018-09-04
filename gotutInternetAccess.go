package main

import ("fmt"
		"net/http"
		"io/ioutil")

// part 1 - https://www.youtube.com/watch?v=ccANcNk8Dac

/*
	Get information from the Internet
*/
func main() {
	// use '_' to define a variable which you don't intend to use
	// '_' here is used for 'err' response
	resp, _ := http.Get("http://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println("string_body: ", string_body)
	resp.Body.Close()
}