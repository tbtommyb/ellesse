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
	fmt.Printf("%-[1]*[2]s %[3]*[4]d %-[5]*[6]s\n", modeWidth, e.file.Mode(), sizeWidth, e.file.Size(), nameWidth, name)
}

func printHeader() error {
	fmt.Printf("%-[1]*[2]s %[3]*[4]s %-[5]*[6]s\n", modeWidth, "Mode", sizeWidth, "Size", nameWidth, "Name")
	return nil
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
	sizeWidth = 8
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
