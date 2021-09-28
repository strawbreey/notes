package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	arr := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	num := make([]string, 0)
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
		for _, v := range words[i] {
			if strings.ContainsRune(arr[0], v) || strings.ContainsRune(arr[1], v) || strings.ContainsRune(arr[2], v) {

			} else {
				fmt.Printf("23333")
				break
			}
		}
		num = append(num, words[i])

	}
	return num
}

func main() {
	nums := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Print(findWords(nums))
}
