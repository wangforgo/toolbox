package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const (
	kb = 1 << 10
	mb = kb << 10
)

type para struct {
	pDir     string
	pHidden  bool
	pMinSize int
	pMaxSize int
	pSuffix  string
	pType    string
}

var p = &para{}

func initFlags() {
	flag.StringVar(&(p.pDir), "dir", "", "directory to be searched")
	flag.BoolVar(&(p.pHidden), "hidden", false, "search hidden files only?")
	flag.IntVar(&(p.pMinSize), "minSize", 0, "min size (kb) of files to be searched")
	flag.IntVar(&(p.pMaxSize), "maxSize", -1, "max size (kb) of files to be searched")
	flag.StringVar(&(p.pSuffix), "suffix", "", "suffix of files to be searched")
	flag.StringVar(&(p.pType), "type", "", "internal type of files to be searched")

	flag.Parse()
	if p.pDir == "" {
		fmt.Printf("invalid directory\n")
		os.Exit(-1)
	}
	if p.pMinSize < 0 {
		p.pMinSize = 0
	}
	if p.pMaxSize < p.pMinSize {
		p.pMaxSize = -1
	}

	fmt.Printf("Flags parsed. Got:\n%+v\n", p)
}

func main() {

	initFlags()

	resultPath := p.pDir + "result.txt"
	f, err := os.Create(resultPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	filepath.Walk(p.pDir, func(path string, info os.FileInfo, err error) error {
		// check dir
		if info.IsDir() {
			return nil
		}

		// check hidden
		if p.pHidden && !isFileHidden(info) {
			return nil
		}

		// check size
		size := int(info.Size())
		if size < p.pMinSize*kb || p.pMaxSize > 0 && size > p.pMaxSize {
			return nil
		}

		// check suffix
		if p.pSuffix != "" && !strings.HasSuffix(info.Name(), p.pSuffix) {
			return nil
		}

		// check file type todo

		// record satisfied files
		f.WriteString(path + "\n")

		return nil
	})
}

func isJpg(path string) bool {

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return false
	}
	if len(bytes) < 2 {
		return false
	}
	return bytes[0] == 0xFF && bytes[1] == 0xD8
}

func isFileHidden(file os.FileInfo) bool {
	fa := reflect.ValueOf(file.Sys()).Elem().FieldByName("FileAttributes").Uint()
	attr := []byte(strconv.FormatUint(fa, 2))
	if attr[len(attr)-2] == '1' {
		return true
	}
	return false
}
