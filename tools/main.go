package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const url = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/"
const dirName = "wordlist"

// path => export_name
var langs = map[string]string{
	"chinese_simplified":  "ChineseSimplified",
	"chinese_traditional": "ChineseTraditional",
	"english":             "English",
	"french":              "French",
	"italian":             "Italian",
	"japanese":            "Japanese",
	"korean":              "Korean",
	"spanish":             "Spanish",
}

const text = `package wordlist

var {{.Name}} = []string{ 
{{ range .Array }}{{if .}} "{{.}}", {{end}} 
{{ end}}
}
`

type Data struct {
	Name  string
	Array []string
}

var filetpl = template.Must(template.New("data").Parse(text))

func main() {
	// create new word list dir
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	for path, name := range langs {
		err := func() error {
			log.Printf("Getting %s's wordlist\n", name)
			resp, err := http.Get(fmt.Sprintf("%s%s.txt", url, path))
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			src, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			wfpath := fmt.Sprintf("%s/%s.go", dirName, path)
			wfile, err := os.OpenFile(wfpath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
			if err != nil {
				return err
			}
			defer wfile.Close()

			log.Printf("Writing %s's world list\n", name)
			data := Data{Array: strings.Split(string(src), "\n"), Name: name}
			return filetpl.Execute(wfile, data)
		}()

		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Update successful!")
}
