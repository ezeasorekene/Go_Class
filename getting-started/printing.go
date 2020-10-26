package main

import (
	"fmt"
//	"strings"
)

func main()  {
	// var i,v int
	// var a [5]int
	// var b [5]int = [5]{1,2,3,4,5}
	c := [...]int{5,7,3,4,5}
	// a[1] = 456
	// fmt.Println(a[1],a[2],a[3],a[4],"P")
	// s1 := c[1:3]
	s2 := c[2:5]
	for i,v := range s2 {
		fmt.Printf("Index -> %d, Value -> %d \n", i, v)
	}
	fmt.Println(len(s2), cap(s2))
}





//Pointers and Deallocation example
// func foo() *int {
//    x := 1
//    return &x
// }

//func main() {
//    var y *int
//    y = foo()
//    fmt.Println(*y)
//}

/**
func main() {
    var y *int
    y = foo()
    fmt.Printf("%d", *y)
}

func main() {
  s := strings.Replace("ianianian", "ni", "in", 2)
  fmt.Println(s)
}


func main() {
  x:=7
  switch {
    case x>3:
      fmt.Printf("1")
    case x>5:
      fmt.Printf("2")
    case x==7:
      fmt.Printf("3")
    default:
      fmt.Printf("4")
  }
}
**/

// func main() {
//   var xtemp int
//   x1 := 0
//   x2 := 1
//   for x:=0; x<5; x++ {
//     xtemp = x2
//     x2 = x2 + x1
//     x1 = xtemp
//   }
//   fmt.Println(x2)
// }
