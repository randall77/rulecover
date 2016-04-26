package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"flag"
)

func main() {
	flag.Parse()

	rulefile := flag.Arg(0)
	coverdata := flag.Arg(1)

	// read coverdata into map
	m := map[string]bool{}
	f, err := os.Open(coverdata)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "rewrite ") {
			continue
		}
		line = line[8:]
		m[line] = true
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}

	// now read rulefile, write output
	f, err = os.Open(rulefile)
	if err != nil {
		panic(err)
	}
	base := path.Base(rulefile)
	i := 0
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		i++
		comment, _ := regexp.MatchString("\\s*//", line)
		if m[fmt.Sprintf("%s:%d", base, i)] {
			fmt.Print("*** ")
		} else if strings.Contains(line, "->") && !comment {
			fmt.Print("--- ")
		} else {
			fmt.Print("    ")
		}
		fmt.Println(line)
	}
	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}
