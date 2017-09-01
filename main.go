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

type colorSprintFunc func(a ...interface{}) string

type colorConfig struct {
	dir     colorSprintFunc
	symlink colorSprintFunc
}

func truncate(s string, n int) string {
	length := len(s)
	if length > n {
		s = s[:n-3] + "..."
	}
	return s
}

// interface{} is used to allow values of different types to be outputted
func outputLine(val1, val2, val3 interface{}) {
	fmt.Printf("%-[1]*[2]v %[3]*[4]v %-[5]*[6]v\n", modeWidth, val1, sizeWidth, val2, nameWidth, val3)
}

func printEntry(f os.FileInfo, c *colorConfig) {
	name := truncate(f.Name(), nameWidth)
	if showColors {
		if f.IsDir() {
			name = c.dir(name)
		}
		if f.Mode()&os.ModeSymlink != 0 {
			name = c.symlink(name)
		}
	}
	outputLine(f.Mode(), f.Size(), name)
}

func printHeader() {
	outputLine("Mode", "Size", "Name")
}

// TODO: Move to be with var declaration
func init() {
	var err error
	flag.BoolVar(&showColors, "colors", false, "Toggle colours in output")
	flag.Parse()

	dirname = flag.Arg(0)
	if dirname == "" {
		dirname = "."
	}

	termWidth, _, err = terminal.GetSize(0)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Set constants in the var
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
		printEntry(entry, colorConf)
	}
}
