package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	if a := 17; a%17 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	a := strconv.Itoa(17)
	fmt.Println(a)

	for a > 1 {
		fmt.Println("It has more than 1 digit")
	}

}
