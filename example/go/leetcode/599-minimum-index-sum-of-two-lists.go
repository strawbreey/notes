package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {

	m := make(map[string]int)

	for k, v := range list1 {
		m[v] = -k - 1
	}

	for k, v := range list2 {

		if m[v] < 0 {
			m[v] = k - m[v]
		}
	}

	min := 10000
	key := ""
	for k, v := range m {
		if v > 0 && v < min {
			min = v
			key = k
		}
	}

	return []string{key}
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}

	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}

	fmt.Print(findRestaurant(list1, list2))
}
