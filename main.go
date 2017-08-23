package main

import (
  "io/ioutil"
  "fmt"
  "log"
  "os"
)

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
  for _, entry := range entries {
    fmt.Println(entry.Name())
  }
}
