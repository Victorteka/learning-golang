package main

import (
  "os"
)

func main()  {
  finfo, err := os.Stat("filename.txt")
  if err != nil {
      // no such file or dir
      return
  }
  if finfo.IsDir() {
      // it's a file
  } else {
      // it's a directory
  }
}
