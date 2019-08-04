package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type TaskGroup struct {
	Kind      string `json:"kind"`
	IsDefault bool   `json:"isDefault"`
}

type Task struct {
	Label   string     `json:"label"`
	Type    string     `json:"type"`
	Command string     `json:"command"`
	Group   *TaskGroup `json:"group,omitempty"`
}

type Tasks struct {
	Version string `json:"version"`
	Tasks   []Task `json:"tasks"`
}

var makefileTaskRegexp = regexp.MustCompile("^[\\w\\d\\-_\\.]+:.*$")

func main() {
	var write bool
	flag.BoolVar(&write, "w", false, "Write to .vscode/tasks.json")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "No Makefile file given")
		os.Exit(1)
	}
	byts, err := ioutil.ReadFile(flag.Arg(0))
	panicIfErr(err)

	tasks := Tasks{}
	tasks.Version = "2.0.0"

	n := 0
	for _, line := range strings.Split(string(byts), "\n") {
		if makefileTaskRegexp.MatchString(line) {
			taskName := strings.Split(line, ":")[0]

			task := Task{}
			task.Label = fmt.Sprintf("make %s", taskName)
			task.Type = "shell"
			task.Command = fmt.Sprintf("make %s", taskName)
			if n == 0 {
				task.Group = &TaskGroup{
					Kind:      "build",
					IsDefault: true,
				}
			}

			tasks.Tasks = append(tasks.Tasks, task)
			n++
		}
	}

	jsonByts, err := json.MarshalIndent(tasks, "", "    ")
	panicIfErr(err)

	if write {
		panicIfErr(ioutil.WriteFile(".vscode/tasks.json", jsonByts, 0x700))
	} else {
		fmt.Println("Copy-paste this in your .vscode/tasks.json")
		fmt.Println(string(jsonByts))
	}
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
