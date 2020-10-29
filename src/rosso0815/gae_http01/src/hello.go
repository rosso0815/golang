package hello

import (
	"html/template"
	"net/http"
	"path"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "controllerIndex 1")
	User := struct {
		UserName string
		Title    string
	}{
		UserName: "PfisterA",
		Title:    "PfisterA-Title",
	}

	tmpl, err := template.New("").Delims("[[", "]]").ParseGlob(path.Join("templates", "*.html"))
	if err != nil {
		log.Infof(ctx, "error ", err)
	} else {
		log.Infof(ctx, "no error ")
	}
	tmpl.ExecuteTemplate(w, "index.html", User)
}

func handlerStatic(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Serving the static page = ", r.URL.Path)

	http.ServeFile(w, r, "./"+r.URL.Path)
}

//-----------------------------------------------------------------------------
// main function in google gae
func init() {

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/static/", handlerStatic)

}

//-----------------------------------------------------------------------------
