package main

import "fmt"

type GiveAway struct {
    loud		string
		louder	string
    loudest	string
}

func shoutMe(shout GiveAway) {
    fmt.Println(shout.loud + shout.louder + shout.loudest)
}

func main() {
    var shout GiveAway
    shout.loud = "Me"
		shout.louder = "eee"
    shout.loudest = "eeeeee"
		count := 10
		// s := shoutMe(shout)
		for index := 0; index < count; index++ {
			fmt.Printf("%d - ", index+1)
			shoutMe(shout)
		}
}



// import (
// 	"fmt"
// //	"strings"
// 	"net/http"
// )
//
//
// func main()  {
// 	_, x := http.Get("www.uci.edu")
// 	fmt.Println(x)
// }

// func main() {
//   s := make([]int, 0, 3)
//   s = append(s, 100)
//   fmt.Println(len(s), cap(s))
// }

// type P struct {
//     x string
// y int
// }
// func main() {
//   b := P{"x", -1}
//   a := [...]P{P{"a", 10},
//         P{"b", 2},
//         P{"c", 3}}
//   for _, z := range a {
//     if z.y > b.y {
//       b = z
//     }
//   }
//   fmt.Println(b.x)
// }

// func main() {
//   x := map[string]int {
//     "ian": 1, "harris": 2}
//   for i, j := range x {
//     if i == "harris" {
//       fmt.Print(i, j)
//     }
//   }
// }


// func main() {
//   x := [...]int {1, 2, 3, 4, 5}
//   y := x[0:2]
//   z := x[1:4]
//   fmt.Print(len(y), cap(y), len(z), cap(z))
// }


// func main() {
//   x := [...]int {4, 8, 5}
//   y := x[0:2]
//   z := x[1:3]
//   y[0] = 1
//   z[1] = 3
//   fmt.Print(x)
// }


// func main() {
//   x := []int {4, 8, 5}
//   y := -1
//   for _, elt := range x {
//     if elt > y {
//       y = elt
//     }
//   }
//   fmt.Print(y)
// }


// func main()  {
// 	// var i,v int
// 	// var a [5]int
// 	// var b [5]int = [5]{1,2,3,4,5}
// 	c := [...]int{5,7,3,4,5}
// 	// a[1] = 456
// 	// fmt.Println(a[1],a[2],a[3],a[4],"P")
// 	// s1 := c[1:3]
// 	s2 := c[2:5]
// 	for i,v := range s2 {
// 		fmt.Printf("Index -> %d, Value -> %d \n", i, v)
// 	}
// 	fmt.Println(len(s2), cap(s2))
// }





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
