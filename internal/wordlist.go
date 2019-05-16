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
	"sync"
)

const url = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/"

var langs = []string{
	"chinese_simplified", "chinese_traditional",
	"english", "french", "italian",
	"japanese", "korean", "spanish",
}

const dirName = "wordlist"

func main() {

	{
		// clean old data
		fd, err := os.Open(dirName)
		if err == nil {
			if err := os.RemoveAll(dirName); err != nil {
				log.Fatal(err)
			}
		}
		fd.Close()

		// create new word list dir
		if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	errchan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		defer close(errchan)
		wg := new(sync.WaitGroup)
		wg.Add(len(langs))
		for idx, lang := range langs {
			go func(idx int, lang string) {
				defer wg.Done()
				select {
				case <-ctx.Done():
					return
				default:
				}
				resp, err := http.Get(url + lang + ".txt")
				if err != nil {
					errchan <- err
					return
				}
				defer resp.Body.Close()

				src, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					errchan <- err
					return
				}

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

				content := fmt.Sprintf("package wordlist\n var %s = []string{", exportName)

				res := strings.Split(string(src), "\n")
				for _, v := range res {
					if v == "" {
						continue
					}
					content += `"` + v + `", `
				}

				content += "}"

				if err := ioutil.WriteFile(dirName+"/"+lang+".go", []byte(content), os.ModePerm); err != nil {
					errchan <- err
					return
				}
			}(idx, lang)
		}
		wg.Wait()
	}()

	var err error
	for v := range errchan {
		if v != nil {
			err = v
			cancel()
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("gofmt", "-w", dirName)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Update successful!")
}
