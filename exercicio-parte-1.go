package main

import "fmt"

func main() {
	divisibleByThree := []int{}

		for i := 1; i <= 100; i++ {
				if(i % 3 == 0) {
					divisibleByThree = append(divisibleByThree, i)

				} 
		}
		fmt.Println(divisibleByThree) 
}