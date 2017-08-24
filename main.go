package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"log"
	"os"
)

var (
	showColors bool
	dirname    string
	termWidth  int
	modeWidth  int
	sizeWidth  int
	nameWidth  int
)

type colorConfig struct {
	dir     (func(a ...interface{}) string)
	symlink (func(a ...interface{}) string)
}

type fileInfo struct {
	file   os.FileInfo
	colors *colorConfig
}

func truncate(s string, n int) string {
	length := len(s)
	if length > n {
		s = s[:n-3] + "..."
	}
	return s
}

func outputLine(vals ...interface{}) {
	fmt.Printf("%-[1]*[2]v %[3]*[4]v %-[5]*[6]v\n", modeWidth, vals[0], sizeWidth, vals[1], nameWidth, vals[2])
}

func (e *fileInfo) print() {
	name := truncate(e.file.Name(), nameWidth)
	if showColors {
		if e.file.IsDir() {
			name = e.colors.dir(name)
		}
		if e.file.Mode()&os.ModeSymlink != 0 {
			name = e.colors.symlink(name)
		}
	}
	outputLine(e.file.Mode(), e.file.Size(), name)
}

func printHeader() {
	outputLine("Mode", "Size", "Name")
}

func init() {
	var err error
	flag.BoolVar(&showColors, "colors", false, "Toggle colours in output")
	flag.Parse()

	dirProvided := flag.NArg() > 0
	if dirProvided {
		dirname = flag.Arg(0)
	} else {
		dirname = "."
	}

	termWidth, _, err = terminal.GetSize(0)
	if err != nil {
		log.Fatal(err)
	}
	modeWidth = 12
	sizeWidth = 10
	nameWidth = termWidth - modeWidth - sizeWidth - 2 // -2 to allow for spaces
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
