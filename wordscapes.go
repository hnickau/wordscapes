package main

import (
  "fmt"
  "os"
	"os/exec"
	"runtime"
  "strings"
  "strconv"
  "log"
  "io"
)

func main() {
  args := make([] string, 1)
  stdins := make([] io.WriteCloser, 1)
  cs := make(map[string]int)
  letters := strings.Split(os.Args[1],"")
  minw := 3
  maxw := 7
  if len(os.Args) > 2 {
    minw,_ = strconv.Atoi(os.Args[2])
  }
  if len(os.Args) > 3 {
    maxw,_ = strconv.Atoi(os.Args[3])
  }
  for _,l := range letters {
    cs[l] = cs[l] + 1
  }
  fmt.Println(cs)
  args[0] = "^[" 
  for l, c := range cs {
    args[0] += l
    c = c
  }
  args[0] = fmt.Sprintf("%s]\\{%d,%d\\}$",args[0],minw,maxw)
  //args[0] = "hello"
  fmt.Println(args[0])
  wordsdir := "/users/nickau/local/"
  if runtime.GOOS == "darwin" {
    wordsdir = "/Users/nickau/local/" 
  }
	cmd := exec.Command("grep", args[0], wordsdir+"words")
  cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)

	err := cmd.Run()
	if false { //err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}  
}


