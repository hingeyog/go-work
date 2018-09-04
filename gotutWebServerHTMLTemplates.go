package main

import ("fmt"
		"net/http"
		"html/template")

// part 16 - https://www.youtube.com/watch?v=Uo9MSE2Gbzs
// HTML Templates
// part 17 - https://www.youtube.com/watch?v=GnLHI_nekm8

type NewsAggPage struct {
	Title string
	News string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<h1>Het there!</h1>
					<p>Go is fast!</p>
					<p>... and simple!</p>`)
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// Initializing the page
	p := NewsAggPage{Title : "Amazing News Aggregator!", News : "Some news!"}

	// t - template, _ - error
	t, _ := template.ParseFiles("basictemplating.html")

	// execute the tempalte with writer and values (which is the page)
	t.Execute(w, p)

	//fmt.Println(t.Execute(w, p)) --> will actually print the errors if any in executing the template 
	// e.g. <h1>{{  .Titles }}</h1> instead of <h1>{{  .Title }}</h1> in html template
} 

func main() {
	//<base URL>, <function to handle>
	http.HandleFunc("/", index_handler)

	http.HandleFunc("/agg/", newsAggHandler)

	//start the server
	//<port number> , <SOMETHING??>
	http.ListenAndServe(":8000", nil)
}