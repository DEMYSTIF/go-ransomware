package main

import (
	"fmt"
	"os"
)

func main() {
	content := []string{}

	f, err := os.Open("./")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		if !v.IsDir() && v.Name() != "ransomware.go" {
			content = append(content, v.Name())
		}
	}

	fmt.Println(content)
}
