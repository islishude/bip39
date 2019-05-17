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

var langs = []string{
	"chinese_simplified", "chinese_traditional",
	"english", "french", "italian",
	"japanese", "korean", "spanish",
}

func getExportName(idx int) string {
	var exportName string
	switch idx {
	case 0:
		exportName = "ChineseSimplified"
	case 1:
		exportName = "ChineseTraditional"
	case 2:
		exportName = "English"
	case 3:
		exportName = "French"
	case 4:
		exportName = "Italian"
	case 5:
		exportName = "Japanese"
	case 6:
		exportName = "Korean"
	case 7:
		exportName = "Spanish"
	}
	return exportName
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
	for idx, lang := range langs {
		idx, lang := idx, lang
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			resp, err := http.Get(url + lang + ".txt")
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			src, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			exportName := getExportName(idx)
			content := fmt.Sprintf("package wordlist\n// %s is word list\n var %s = []string{", exportName, exportName)
			res := strings.Split(string(src), "\n")
			for _, v := range res {
				if v == "" {
					continue
				}
				content += `"` + v + `", `
			}
			content += "}"
			return ioutil.WriteFile(dirName+"/"+lang+".go", []byte(content), os.ModePerm)
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
