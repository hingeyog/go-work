package main

import ("fmt"
		"net/http")

// https://www.youtube.com/watch?v=QSLvvxzTa7o

func index_handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<h1>Het there!</h1>
					<p>Go is fast!</p>
					<p>... and simple!</p>`)

/*
	fmt.Fprintf(w, "<h1>Het there!</h1>")
	fmt.Fprintf(w, "<p>Go is fast!</p>")
	fmt.Fprintf(w, "<p>... and simple!</p>")
	fmt.Fprintf(w, "<p>YOu %s even add %s</p>", "can", "<strong>variables</strong>")
*/

}

func main() {
	//<base URL>, <function to handle>
	http.HandleFunc("/", index_handler)

	//start the server
	//<port number> , <SOMETHING??>
	http.ListenAndServe(":8000", nil)
}