package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file1Name := "/Users/bytedance/Desktop/a"

	file1, err := os.Open(file1Name)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer func() {
		_ = file1.Close()
	}()

	reader1 := bufio.NewReader(file1)
	var list1 []string
	stringMap := make(map[string]bool, 100)
	for {
		a, _, c := reader1.ReadLine()
		if c == io.EOF {
			break
		}
		list1 = append(list1, string(a))
		stringMap[string(a)] = false
	}
	for _, v := range list1 {
		fmt.Printf("'%s',\n", v)
	}

}
