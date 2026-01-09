package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the main string: ")
	mainStr, _ := reader.ReadString('\n')
	mainStr = strings.TrimSpace(mainStr)

	fmt.Print("Enter the substring to search: ")
	subStr, _ := reader.ReadString('\n')
	subStr = strings.TrimSpace(subStr)

	if mainStr == "" || subStr == "" {
		fmt.Println("Please provide non-empty strings")
		return
	}

	if strings.Contains(mainStr, subStr) {
		fmt.Printf("\"%s\" is a substring of \"%s\"\n", subStr, mainStr)
	} else {
		fmt.Printf("\"%s\" is NOT a substring of \"%s\"\n", subStr, mainStr)
	}
}
