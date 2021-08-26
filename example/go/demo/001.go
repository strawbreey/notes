package main

import "fmt"

func main() {
//    fmt.Println("Hello, World!")
   var a [3]int
//    fmt.Println(a[0])
//    fmt.Println(a[2])

   for i, v := range a {
	//    fmt.Println("%d %d", i, v)
	   fmt.Printf("%d %d\n", i, v)
   }
}