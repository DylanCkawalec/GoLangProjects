package main

import (
	"fmt"
)

type numberSet []int

func main() {
	

}

func createNewNumberSet(numberSet []int) (int, int) {
	var max int = numberSet[0]
	var min int = numberSet[0]
	for _, value := range numberSet {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func (ns numberSet) odd_even(int) string {
	message := ""
	for check, value := range ns {
		//check the index and see if it's even or odd
		if check%2 == 0 {
			value = value
			message = "is even"
			fmt.Println(message)
		} else {
			value = value
			message = "is odd"
			fmt.Println(message)
		}
	}
	return message
}

func (ns numberSet) printEvenOrOdd() (index int, teller string) {
	for _, value := range ns {
		fmt.Println(index, value)
		for _, oe := range ns {
			if oe%2 == 0 {
				fmt.Println(teller, "is even")
			} else {
				fmt.Println(teller, "is odd")
			}
		}
	}
	return index, teller

}
