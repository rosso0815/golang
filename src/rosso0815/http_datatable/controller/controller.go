package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

//------------------------------------------------------------------------------
func init() {
	log.Println("### ControllerApi init done")
}

//APIControllerindex lets play
func APIControllerindex(w http.ResponseWriter, r *http.Request) {
	log.Print("controllerIndex")

	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

	tmpl, err := template.New("").Delims("[[", "]]").ParseFiles(lp, fp)
	if err != nil {
		log.Print("error ", err)
	} else {
		log.Print("no error ")
	}

	tmpl.ExecuteTemplate(w, "layout", nil)
}

//-----------------------------------------------------------------------------
