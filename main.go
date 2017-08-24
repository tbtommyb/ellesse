package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
)

type colorConfig struct {
	dir     (func(a ...interface{}) string)
	symlink (func(a ...interface{}) string)
}

type fileInfo struct {
	file   os.FileInfo
	colors *colorConfig
}

func (e *fileInfo) print() {
	name := e.file.Name()
	if showColors {
		if e.file.IsDir() {
			name = e.colors.dir(name)
		}
		if e.file.Mode()&os.ModeSymlink != 0 {
			name = e.colors.symlink(name)
		}
	}
	fmt.Printf("%-12s %10d %-20s\n", e.file.Mode(), e.file.Size(), name)
}

var (
	showColors bool
	dirname    string
)

func printHeader() error {
	fmt.Printf("%-12s %10s %-20s\n", "Mode", "Size", "Name")
	return nil
}

func init() {
	flag.BoolVar(&showColors, "colors", false, "Toggle colours in output")
	flag.Parse()

	dirProvided := flag.NArg() > 0
	if dirProvided {
		dirname = flag.Arg(0)
	} else {
		dirname = "."
	}
}

func main() {
	colorConf := &colorConfig{
		dir:     color.New(color.FgYellow).SprintFunc(),
		symlink: color.New(color.FgRed).SprintFunc(),
	}

	entries, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	printHeader()
	for _, entry := range entries {
		file := &fileInfo{
			file:   entry,
			colors: colorConf,
		}
		file.print()
	}
}
