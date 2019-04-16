package main

import (
  "fmt"
  "os"
	"os/exec"
	"io/ioutil"
	"runtime"
  "strings"
  "strconv"
  "log"
  "sort"
//  "io"
)

type byLength []string
// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. 
func (s byLength) Len() int {
  return len(s)
}
func (s byLength) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
  if len(s[i]) == len(s[j]) {
    return s[i] < s[j]
  } 
  return len(s[i]) < len(s[j]) 
}



func main() {
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
//  fmt.Println(cs)
  args := "^[" 
  for l, _ := range cs {
    args += l
  }
  args = fmt.Sprintf("%s]\\{%d,%d\\}$",args,minw,maxw)
  wordsdir := "/users/nickau/local/"
  if runtime.GOOS == "darwin" {
    wordsdir = "/Users/nickau/local/" 
  }
	cmd := exec.Command("grep", args, wordsdir+"words")
	
  stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

  results, _ := ioutil.ReadAll(stdout)
  cmd.Wait()

  for l, c := range cs {
    regex := l
    for c > 0 {
      regex += ".*" + l
      c = c - 1
    } 
    cmd = exec.Command("grep", "-v", regex)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
	    log.Fatal(err)
    }
    stdin, err := cmd.StdinPipe()
    if err != nil {
	    log.Fatal(err)
    }
	  if err := cmd.Start(); err != nil {
		  log.Fatal(err)
	  }
    
    fmt.Fprintf(stdin, "%s", string(results)) 
    stdin.Close()
    
    results, _ = ioutil.ReadAll(stdout)
    cmd.Wait()
  }
    
  lines := strings.Split(string(results), "\n")
  sort.Sort(byLength(lines))
  result := strings.Join(lines,"\n")
  fmt.Println(result)
}


