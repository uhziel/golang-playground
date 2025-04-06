package main

import (
	"fmt"
	"os"

	"github.com/moby/patternmatcher"
	"github.com/moby/patternmatcher/ignorefile"
)

func main() {
	f, err := os.Open("gh1ignore")
	if err != nil {
		panic(err)
	}

	patterns, err := ignorefile.ReadAll(f)
	if err != nil {
		panic(err)
	}

	pm, err := patternmatcher.New(patterns)
	if err != nil {
		panic(err)
	}

	TestMatch(pm, ".rcon-cli.env")
	TestMatch(pm, "1.txt")
	TestMatch(pm, "logs")
	TestMatch(pm, "logs/1.log")
	TestMatch(pm, "logs/login/1.log")
}

func TestMatch(pm *patternmatcher.PatternMatcher, file string) {
	res, _ := pm.MatchesOrParentMatches(file)
	fmt.Println(file, res)
}
