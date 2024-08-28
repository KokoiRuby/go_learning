## [gowiki](https://go.dev/doc/articles/wiki/)

### Data Structure

一个 wiki 包含一连串相互连接的页面，每个页面都有一个 Title & Body。

```go
// wiki.go
type Page struct {
    Title string
    Body  []byte
}

// save persists page body into file.go named with page title.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// load reads page on disk given title
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{
		Title: title,
		Body:  body,
	}, nil
}

func TestServerAndLoadPage() {
    p1 := &Page{
		Title: "TestPage",
		Body:  []byte("This is a sample Page."),
	}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))  // This is a sample Page.
}
```

### Using net/http to server wiki pages

```go
cat > test.txt << EOF
Hello world
EOF
```

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// http://localhost:8080/view/test
func TestView() {
    http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Editing Pages

```go
const editForm = `
	<html>
		<body>
			<h1>Editing %s</h1>
				<form action="/save/%s" method="POST">
					<textarea name="body">%s</textarea><br>
					<input type="submit" value="Save"/>
				</form>
		</body>
	</html>
`

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{
			Title: title,
		}
	}
	fmt.Fprintf(w, editForm, p.Title, p.Title, p.Body)
}

// http://localhost:8080/edit/test
func TestEdit() {
    http.HandleFunc("/edit/", editHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### [html/template](https://pkg.go.dev/html/template)

将 HTML 以模板单独保存，而不是以源码的形式注入，利于维护和修改。

`.Title` and `.Body` dotted identifiers refer to `p.Title` and `p.Body`.

```bash
$ cat > edit.html << EOF
<html>
	<body>
		<h1>Editing {{.Title}}</h1>
			<form action="/save/{{.Title}}" method="POST">
				<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
				<div><input type="submit" value="Save"></div>
			</form>
	</body>
</html>
EOF
```

```go
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{
			Title: title,
		}
	}
	// fmt.Fprintf(w, editForm, p.Title, p.Title, p.Body)
    
    // read file.go & return *template.Template
	t, _ := template.ParseFiles("edit.html")
    // write html into http.ResponseWriter
	t.Execute(w, p)
}
```

Extract

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

renderTemplate(w, "edit", p)
renderTemplate(w, "view", p)
renderTemplate(w, "save", p)
```

Handling non-existent pages, if the page does not exist, it shall **redirect to edit** page.

```go
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}
```

### Error Handling

```go
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	// err handle
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	// err handle
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{
		Title: title,
		Body:  []byte(body),
	}
	err := p.save()
	// err handle
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
```

### Template caching

Inefficiency evey time a page is rendered.

Good practice: render all templates at init into a single `*Template`.

```go
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
```

### Validation

Security Flaw: a user can supply an arbitrary path to be read/written on the server.

To mitigate this, we can write a function to validate the title with a regex.

```go
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid page title")
	}
	return m[2], nil
}

func *Handler(w http.ResponseWriter, r *http.Request) {
    title, err := getTitle(w, r)
	if err != nil {
		return
	}
    ...
}
```

### Introducing Function Literals and Closures

A wrapper function that *takes a function of the above type*, and returns a function of type `http.HandlerFunc`

```go
func viewHandler(w http.ResponseWriter, r *http.Request, title string)
func editHandler(w http.ResponseWriter, r *http.Request, title string)
func saveHandler(w http.ResponseWriter, r *http.Request, title string)

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

http.HandleFunc("/view/", makeHandler(viewHandler))
http.HandleFunc("/edit/", makeHandler(editHandler))
http.HandleFunc("/save/", makeHandler(saveHandler))
```

