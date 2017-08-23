package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "os"
)

func printEntry(entry os.FileInfo) (error) {
  fmt.Printf("%-12s %10d %-20s\n", entry.Mode(), entry.Size(), entry.Name())
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

  entries, err := ioutil.ReadDir(dirname)
  if err != nil {
    log.Fatal(err)
  }
  printHeader()
  for _, entry := range entries {
    printEntry(entry)
  }
}
