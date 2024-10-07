package main

import (
	"fmt"
)

func main() {
	var Sum int = 0
	var i int
	var itisprime bool
	var Num int
	fmt.Println("prime numbers between 1 and 10 are:-")
	for Num = 1; Num <= 10; Num++ {
		itisprime = true
		if Num < 2 {
			itisprime = false // if num is less than 2 its not a prime number
		} else {
			for i = 2; i*i <= Num; i++ {
				if Num%i == 0 {
					itisprime = false
					break
				}
			}
		}
		if itisprime {
			fmt.Println(Num)
			Sum += Num
		}
	}
	fmt.Println("The sum of prime numbers between 1 and 10 is:-", Sum)
}
