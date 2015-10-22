package main

import "fmt"

 func add(x, y int) int{
 return x+  y
 }

func main(){
n:=add(1,2)

fmt.Println(n)

 m := map[int]string{
   1: "one",
   256: "twohundredandfiftysix",
   1024: "onethousandtwentyfour",
 }

 for k, v := range m {
fmt.Printf("%d: %s\n", k, v)
 }
}
