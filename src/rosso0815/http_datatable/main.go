package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"pfistera/http_datatable/controller"
)

// FileSystem custom file system handler
type FileSystem struct {
	fs http.FileSystem
}

// Open opens file
func (fs FileSystem) Open(path string) (http.File, error) {
	log.Println("Open", path)
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := fs.fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func main() {

	port := flag.String("p", "8081", "port to serve on")

	//dirStatics := flag.String("d", ".", "the directory of static file to host")
	//dirNode := flag.String("n", "node_modules", "the directory of static file to host")
	flag.Parse()

	//fileServerDirectory := http.FileServer(FileSystem{http.Dir(*directory)})
	//http.Handle("/statics", http.FileServer(http.Dir(*dirStatics)))
	//http.Handle("/node_modules", http.FileServer(http.Dir(*dirNode)))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./statics"))))
	http.Handle("/node_modules/", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))

	//http.Handle("/", http.StripPrefix(strings.TrimRight("/statics/", "/"), fileServer))
	//http.Handle("/node_modules/", http.StripPrefix(strings.TrimRight("/node_modules/", "/"), fileServer))

	//log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))

	// add routes
	http.HandleFunc("/api", controller.APIControllerindex)
	//http.HandleFunc("/test01.html", controllerTest01)
	//http.HandleFunc("/datatables.html", controllerDatatables)

	// add static files
	//fs := http.FileServer(http.Dir("bower_components"))
	//http.Handle("/bower_components/", http.StripPrefix("/bower_components/", fs))

	//fs = http.FileServer(http.Dir("node_modules"))
	//http.Handle("/node_modules/", http.StripPrefix("/node_modules/", fs))

	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/", http.StripPrefix("/static/", fs))

	// mux router
	//r := mux.NewRouter()
	//r.HandleFunc("/", controller.APIControllerindex)
	//r.HandleFunc("/datatables.html", controllerDatatables)
	//r.HandleFunc("/javascript.html", controllerJavascript)
	//r.HandleFunc("/vuetable.html", controllerVuetable)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)
	//log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

/*
func controllerTest01(w http.ResponseWriter, r *http.Request) {
	log.Print("controllerTest01")
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "test01.html")
	tmpl, _ := template.New("").Delims("[[", "]]").ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func controllerJavascript(w http.ResponseWriter, r *http.Request) {
	log.Print("controllerJavascript")
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "javascript.html")
	tmpl, _ := template.New("").Delims("[[", "]]").ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func controllerDatatables(w http.ResponseWriter, r *http.Request) {
	log.Print("controllerDatatables")
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "datatables.html")
	tmpl, _ := template.New("").Delims("[[", "]]").ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func controllerVuetable(w http.ResponseWriter, r *http.Request) {
	log.Print("controllerVuetable")
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "vuetable.html")
	tmpl, _ := template.New("").Delims("[[", "]]").ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}
*/

//------------------------------------------------------------------------------
