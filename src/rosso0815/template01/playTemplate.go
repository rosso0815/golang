package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func urgentNote(text string) string {
	return fmt.Sprintf("URGENT >> %s << DONE", text)
}

func main() {

	log.Println("@@@ start")

	type Inventory struct {
		Material string
		Count    uint
	}

	type Report struct {
		Inventory Inventory
		Index     string
	}

	const tmplText = `
{{with .Inventory -}}
	{{ .Material }} items are made of {{ .Count }}
{{- end}}

{{23 -}} < {{- 45}}
{{ .Index | urgentNote }}
done
`
	var err error

	report := Report{Inventory: Inventory{Material: "wool", Count: 17888}, Index: "HELLO"}

	fmap := template.FuncMap{"urgentNote": urgentNote}

	tmpl := template.New("test")

	tmpl = tmpl.Funcs(fmap)

	/* latest 2 actions are parse and execute */
	tmpl, err = tmpl.Parse(tmplText)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, report)
	if err != nil {
		panic(err)
	}

	log.Println("@@@ end")

}
