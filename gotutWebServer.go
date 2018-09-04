package main

import ("fmt"
		"net/http")

/*
 	w: writer
 	r: request
 */
func index_handler(w http.ResponseWriter, r *http.Request) {
	//Fprintf - ???
	/*
		C:\__TOYOTA\__YOGESH\TEST WORKSPACE\go-work>godoc fmt Fprintln 
		use 'godoc cmd/fmt' for documentation on the fmt command
		func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
    	Fprintln formats using the default formats for its operands and writes
    	to w. Spaces are always added between operands and a newline is
    	appended. It returns the number of bytes written and any write error
    	encountered.
	*/
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

/*
 	w: writer
 	r: request
 */
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webpage by yogi!")
}

func main() {
	//<base URL>, <function to handle>
	http.HandleFunc("/", index_handler)
	
	http.HandleFunc("/about/", about_handler)

	//start the server
	//<port number> , <SOMETHING??>
	http.ListenAndServe(":8000", nil)
}