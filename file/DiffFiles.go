package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file1Name := "/Users/bytedance/Desktop/a"
	file2Name := "/Users/bytedance/Desktop/b"

	file1, err := os.Open(file1Name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	file2, err := os.Open(file2Name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer func() {
		_ = file1.Close()
		_ = file2.Close()
	}()

	reader1 := bufio.NewReader(file1)
	reader2 := bufio.NewReader(file2)
	var list1, list2 []string
	stringMap := make(map[string]bool, 100)
	for {
		a, _, c := reader1.ReadLine()
		if c == io.EOF {
			break
		}
		list1 = append(list1, string(a))
		stringMap[string(a)] = false
	}

	for {
		a, _, c := reader2.ReadLine()
		if c == io.EOF {
			break
		}
		list2 = append(list2, string(a))
		stringMap[string(a)] = false
	}
	for _, v1 := range list1 {
		for _, v2 := range list2 {
			if v1 == v2 {
				stringMap[v1] = true
			}
		}
	}
	for s := range stringMap {
		if !stringMap[s] {
			fmt.Println(s)
		}
	}
}
