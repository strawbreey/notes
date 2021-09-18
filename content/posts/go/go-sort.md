---
title: "Go Sort"
date: 2021-09-10T19:40:55+08:00
draft: false
---


### sort

```go
package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	sort.Ints(nums)
	num := nums[len(nums)/2]
	return num
}


```