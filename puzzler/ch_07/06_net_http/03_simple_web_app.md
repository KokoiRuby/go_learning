```go
const form = `
    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello</h1>")
}

func FuncServer(w http.ResponseWriter, r *http.Request) {
	// set header
    w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
        // return form
		io.WriteString(w, form)
	case "POST":
        // parse submitted form 
		io.WriteString(w, r.FormValue("in"))
	}
}

// http://localhost:8080/test1
// http://localhost:8080/test2
func main() {
    http.Handle("/test1", http.HandlerFunc(SimpleServer))
	http.Handle("/test2", http.HandlerFunc(FuncServer))
    if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
```

