package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

var (
	fileInput string
	dirOutput string
)

func main() {
	flag.StringVar(&fileInput, "inputFile", "dayone.md", "day one 归档文件")
	flag.StringVar(&dirOutput, "outputDir", ".", "标准 markdown 格式的导出目录")
	flag.Parse()

	file, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}

	diaryEntries, err := importDayOne(file)
	if err != nil {
		panic(err)
	}

	fileEntries := exportDayOne(diaryEntries)

	for _, fileEntry := range fileEntries {
		file, err := os.Create(path.Join(dirOutput, fileEntry.FileName))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		_, err = file.WriteString(fileEntry.Content)
		if err != nil {
			panic(err)
		}
	}
}

type DiaryEntry struct {
	Date     time.Time
	Location string
	Weather  string
	Photo    string
	Content  []string
}

type FileEntry struct {
	FileName string
	Content  string
}

type FSM struct {
	state FSMState
	entry *DiaryEntry
}

type FSMState string

const (
	Init     FSMState = "init"
	Metadata FSMState = "metadata"
	Content  FSMState = "content"
)

const dayOneTimeLayout = "January 2, 2006 at 15:04"

var beijingTZ = time.FixedZone("Beijing Time", int((8 * time.Hour).Seconds()))

func (f *FSM) Push(line string) *DiaryEntry {
	if f.state == Init {
		sections := strings.SplitN(line, "\t", 3)
		if sections[1] != "Date:" {
			panic(fmt.Errorf("need date line=%s", line))
		}

		f.entry = &DiaryEntry{}
		t, err := time.ParseInLocation(dayOneTimeLayout, sections[2], beijingTZ)
		if err != nil {
			panic(err)
		}
		f.entry.Date = t

		f.state = Metadata
	} else if f.state == Metadata {
		if line == "" {
			f.state = Content
			return nil
		}

		sections := strings.SplitN(line, "\t", 3)
		if sections[1] == "Location:" {
			f.entry.Location = sections[2]
		} else if sections[1] == "Weather:" {
			f.entry.Weather = sections[2]
		} else if sections[1] == "Photo:" {
			f.entry.Photo = sections[2]
		} else if sections[1] == "Tags:" {
		} else if sections[1] == "Starred" {
		} else {
			panic(fmt.Errorf("metadata invalid line=%s", line))
		}
	} else if f.state == Content {
		sections := strings.SplitN(line, "\t", 3)
		if len(sections) > 2 && sections[1] == "Date:" {
			f.entry.Content = f.entry.Content[:len(f.entry.Content)-1]
			curEntry := f.entry

			f.entry = &DiaryEntry{}
			t, err := time.ParseInLocation(dayOneTimeLayout, sections[2], beijingTZ)
			if err != nil {
				panic(err)
			}
			f.entry.Date = t

			f.state = Metadata

			return curEntry
		}

		f.entry.Content = append(f.entry.Content, line)
	}
	return nil
}

func importDayOne(input io.Reader) ([]DiaryEntry, error) {
	scanner := bufio.NewScanner(input)
	entries := []DiaryEntry{}
	fsm := FSM{
		state: Init,
	}
	for scanner.Scan() {
		line := scanner.Text()
		if entry := fsm.Push(line); entry != nil {
			entries = append(entries, *entry)
		}
	}
	if fsm.entry != nil {
		if len(fsm.entry.Content) > 1 && len(fsm.entry.Content[len(fsm.entry.Content)-1]) == 0 {
			fsm.entry.Content = fsm.entry.Content[:len(fsm.entry.Content)-1]
		}
		entries = append(entries, *fsm.entry)
	}
	return entries, scanner.Err()
}

const diaryTmplText = `---
title: {{.Title}}
updated: {{.Date}}
created: {{.Date}}
{{- if .Location}}
location: {{.Location}}
{{- end}}
{{- if .Weather}}
Weather: {{.Weather}}
{{- end}}
---
{{- if .Photo}}

![{{.Photo}}](../../assets/{{.PhotoEscaped}})
{{- end}}

{{.Content -}}
`

type DiaryTmplArgs struct {
	Title        string
	Date         string
	Location     string
	Weather      string
	Photo        string
	PhotoEscaped string
	Content      string
}

func escape(s string) string {
	s = strings.ReplaceAll(s, " ", "%20")
	s = strings.ReplaceAll(s, "(", "%28")
	s = strings.ReplaceAll(s, ")", "%29")

	return s
}

func exportDayOne(entries []DiaryEntry) []FileEntry {
	diaryTmpl, err := template.New("diary").Parse(diaryTmplText)
	if err != nil {
		panic(err)
	}

	fileEntries := []FileEntry{}

	for _, entry := range entries {
		args := DiaryTmplArgs{
			Title:        entry.Date.Format("2006-01-02 15:04"),
			Date:         entry.Date.Format(time.RFC3339),
			Location:     entry.Location,
			Weather:      entry.Weather,
			Photo:        entry.Photo,
			PhotoEscaped: escape(entry.Photo),
			Content:      strings.Join(entry.Content, "\n"),
		}

		b := new(strings.Builder)
		if err := diaryTmpl.Execute(b, args); err != nil {
			panic(err)
		}
		fileEntry := FileEntry{
			FileName: fmt.Sprintf("%s.md", args.Title),
			Content:  b.String(),
		}

		fileEntries = append(fileEntries, fileEntry)
	}
	return fileEntries
}
