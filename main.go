package main

import (
	"bytes"
	//"errors"
	"fmt"
	//"os"
	"io/ioutil"
)

//type saveFile struct {
//	filepath string
//	separator string
//}

//func getFileData() (saveFile, error) {
//	if len(os.Args) < 2 {
//		return saveFile{}, errors.New("A filepath argument is required")
//	}
//	return
//}

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
	b, err := ioutil.ReadFile("./checkpoints/1614939630.ram")
	if err != nil {
		fmt.Println(err)
	}
	// Lives Left comes after "LEFT" string in byte array
	s := "LEFT"
	fmt.Println("Length of byte array is ", len(b))
	fmt.Println(getSearchString(b, s))
}


