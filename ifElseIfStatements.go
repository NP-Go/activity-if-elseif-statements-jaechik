package main

import (
	"fmt"
	"strconv"
)

var userInput string

func main() {
	fmt.Println("Please input a number")
	fmt.Scanln(&userInput)
	num, _ := strconv.ParseInt(userInput, 10, 0)
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	if num%2 == 0 {
		fmt.Println("Even")
		if num > 9 {
			fmt.Println("It has more than 1 digit")
		} else {
			fmt.Println("It has more than 1 digit")
		}
	} else {
		fmt.Println("Odd")
		if num > 9 {
			fmt.Println("It has more than 1 digit")
		} else {
			fmt.Println("It has more than 1 digit")
		}
	}
}
