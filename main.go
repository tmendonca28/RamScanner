package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func getSearchString(data []byte, searchTerm string) int {
	results := make(map[string]int, 0)

	index := len(data)
	tmp := data
	for true {
		match := bytes.LastIndex(tmp[0:index], []byte(searchTerm))
		if match == -1 {
			break
		}else{
			index = match
			results[searchTerm] = match + 5
		}
	}
	byteOffset := results["LEFT"]
	return byteOffset
}

func main() {
	saveFiles, err:= walkMatch("./checkpoints", "*.ram")
	if err != nil {
		fmt.Print(err)
	}
	for _, saveFile := range saveFiles {
		s, err := ioutil.ReadFile(saveFile)
		searchTerm := "LEFT"
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(saveFile)
		fmt.Println(getSearchString(s, searchTerm))
		fmt.Println(s[getSearchString(s, searchTerm)])
		fmt.Println()
	}
}


