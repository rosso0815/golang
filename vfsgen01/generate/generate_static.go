package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/shurcooL/vfsgen"
)

func main() {
	log.Println("@@@ assets_generate start")
	var cwd, _ = os.Getwd()
	templates := http.Dir(filepath.Join(cwd, "static"))
	if err := vfsgen.Generate(templates, vfsgen.Options{
		Filename:     "assets/assets_vfsdata.go",
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
	}); err != nil {
		log.Fatalln(err)
	}
}
