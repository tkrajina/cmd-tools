package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
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
	flag.Parse()

	for _, filename := range flag.Args() {
		bytes, err := ioutil.ReadFile(filename)
		panicIfErr(err)

		parseKindleNotes(string(bytes))
	}
}

func parseKindleNotes(str string) {
	var keys []string
	var texts = map[string][]string{}

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
				_ = loc
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
		if len(texts[key]) == 0 {
			continue
		}
		fmt.Println("#", key, "\n")
		for _, text := range texts[key] {
			fmt.Println(text)
		}
		fmt.Println("\n\n")
	}

	if unknown.Len() > 0 {
		fmt.Println("# Unknown")
		fmt.Println()
		fmt.Println(unknown.String())
	}
}
