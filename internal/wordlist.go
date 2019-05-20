package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sync/errgroup"
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

func init() {
	// clean old data
	fd, err := os.Open(dirName)
	if err == nil {
		if err := os.RemoveAll(dirName); err != nil {
			log.Fatal(err)
		}
	}
	defer fd.Close()

	// create new word list dir
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func main() {
	eg, ctx := errgroup.WithContext(context.Background())
	for path, name := range langs {
		path, name := path, name
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			resp, err := http.Get(url + path + ".txt")
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			src, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			content := fmt.Sprintf("package wordlist\n var %s = [2048]string{", name)
			res := strings.Split(string(src), "\n")
			for _, v := range res {
				if v == "" {
					continue
				}
				content += `"` + v + `", `
			}
			content += "}"
			return ioutil.WriteFile(dirName+"/"+path+".go", []byte(content), os.ModePerm)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("gofmt", "-w", dirName)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Update successful!")
}
