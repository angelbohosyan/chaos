package main

import "fmt"

func pp(x string) {
	for i:=1;i<100;i++ {
		fmt.Println(i)
	}
}
func multiply(a,b int) int {
	sum:=0
	for i:=0;i<b;i++ {
		sum += a;
	}
	return sum
}

func main() {
	fmt.Println(multiply(3,3))
}
