package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "os"
  "github.com/fatih/color"
)

type colorConfig struct {
  dir (func(a ...interface{}) string)
}

type fileInfo struct {
  file os.FileInfo
  config *colorConfig
}

func (e *fileInfo) print() (error) {
  name := e.file.Name()
  if e.file.IsDir() {
    name = e.config.dir(name)
  }
  fmt.Printf("%-12s %10d %-20s\n", e.file.Mode(), e.file.Size(), name)
  return nil
}

func printHeader() (error) {
  fmt.Printf("%-12s %10s %-20s\n", "Mode", "Size", "Name")
  return nil
}

func main() {
  var dirname string

  if len(os.Args) == 1 {
    dirname = "."
  } else {
    dirname = os.Args[1]
  }
  conf := &colorConfig{
    dir: color.New(color.FgYellow).SprintFunc(),
  }

  entries, err := ioutil.ReadDir(dirname)
  if err != nil {
    log.Fatal(err)
  }
  printHeader()
  for _, entry := range entries {
    file := &fileInfo{
      file: entry,
      config: conf,
    }
    file.print()
  }
}
