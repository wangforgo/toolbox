package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	readCfg()
	if len(os.Args) != 2 {
		log.Fatal("usage: lc id\nid is the leetcode problem number\n")
	}

	leetcodeNum,err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("illegal problem number: %v\n", os.Args[1])
	}

	var problemDir = workDir+`/`+strconv.Itoa(leetcodeNum)
	var problemFile = problemDir+`/`+strconv.Itoa(leetcodeNum)+`.go`

	err = os.Mkdir(problemDir, 0777)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(problemFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("package main\n\nfunc main() {\n\n\n}\n")
	if err != nil {
		log.Fatal(err)
	}

	exec.Command(goland, problemFile).Run()
}

var workDir string
var goland string
func readCfg() {
	wd, _ := os.Getwd()
	wd = strings.Replace(wd, `\`,`/`,-1)
	bs, err := ioutil.ReadFile(wd+"/leetcoder.config")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(bs),"\n")
	for i:= range lines {
		s := strings.TrimSpace(lines[i])
		if strings.HasPrefix(s,"goland") {
			goland = strings.Replace(strings.TrimSpace(s[strings.Index(s,"=")+1:]),"\\","/", -1)
		}
		if strings.HasPrefix(s, "dir") {
			workDir = strings.Replace(strings.TrimSpace(s[strings.Index(s,"=")+1:]),"\\","/", -1)
		}
	}
	if workDir == "" {
		log.Fatal("workDir should be set in config.txt: dir = C:/leetcode")
	}
	if goland == "" {
		log.Fatal("goland should be set in config.txt: goland = C:/program/goland.exe")
	}
}
