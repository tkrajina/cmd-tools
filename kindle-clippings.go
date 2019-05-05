package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

var metadataReg = regexp.MustCompile(`\-\s+Your\s(.*?)\s+on\s+(.*?)\s*\|.*`)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var toFile string

	flag.StringVar(&toFile, "to", "", "Append to file")
	flag.Parse()

	for _, filename := range flag.Args() {
		if !strings.HasSuffix(filename, ".txt") {
			filename = path.Join(filename, "documents/My Clippings.txt")
		}
		bytes, err := ioutil.ReadFile(filename)
		panicIfErr(err)

		res := parseKindleNotes(string(bytes))
		fmt.Println(res)

		if len(toFile) > 0 {
			f, err := os.OpenFile(toFile, os.O_APPEND|os.O_WRONLY, 0600)
			panicIfErr(err)
			_, err = f.WriteString(res)
			panicIfErr(err)
			_ = f.Close()

			fmt.Printf("Appended to %s\n", toFile)
			fmt.Printf("Empty %s [y/N]?", filename)
			var answer string
			fmt.Scanf("%s", &answer)
			if answer == "y" {
				fmt.Println("Emptying " + filename)
				err = ioutil.WriteFile(filename, []byte{}, 777)
				panicIfErr(err)
			}
		}
	}
}

func parseKindleNotes(str string) string {
	var keys []string
	var texts = map[string][]string{}

	var res bytes.Buffer
	var unknown bytes.Buffer
	parts := strings.Split(str, "==========")
	for _, part := range parts {
		part = strings.TrimLeft(part, " \n\r\t")
		lines := strings.Split(part, "\n")
		if len(lines) < 3 {
			unknown.WriteString(part)
		} else {
			title := lines[0]
			metadata := lines[1]
			text := strings.TrimSpace(strings.Join(lines[2:], "\n"))
			//fmt.Println("title:", title)
			//fmt.Println("text:", text)

			if metadataReg.MatchString(metadata) {
				all := metadataReg.FindAllStringSubmatch(metadata, -1)
				typ := all[0][1]
				loc := all[0][2]

				if typ == "Note" {
					text = "Note: " + text
				}
				if typ == "Highlight" || typ == "Note" {
					typ = "Txt"
				}

				//fmt.Printf("type=%s location=%s\n", typ, loc)

				key := fmt.Sprintf("%s :: %s", strings.TrimSpace(title), strings.TrimSpace(typ))
				if _, found := texts[key]; !found {
					keys = append(keys, key)
					texts[key] = []string{}
				}
				texts[key] = append(texts[key], text)
			} else {
				unknown.WriteString(part)
			}
		}
	}

	sort.Strings(keys)

	for _, key := range keys {
		txt := strings.TrimSpace(strings.Join(texts[key], "\n"))
		if len(txt) == 0 {
			continue
		}
		res.WriteString("# " + key + "\n\n")
		res.WriteString(txt)
		res.WriteString("\n\n")
	}

	if unknown.Len() > 0 {
		res.WriteString("# Unknown\n\n")
		res.WriteString(unknown.String())
	}

	return res.String()
}
