package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	a = Solution(a, 6)
	fmt.Println(a)

}

func Solution(inputArray []int, countRotation int) []int {
	lengthInputArray := len(inputArray)
	for rotationIncrement := 0; rotationIncrement < countRotation; rotationIncrement++ {
		lastElement := inputArray[lengthInputArray-1]

		for i := lengthInputArray - 2; i >= 0; i-- {
			inputArray[i+1] = inputArray[i]
		}

		inputArray[0] = lastElement
	}
	return inputArray
}
