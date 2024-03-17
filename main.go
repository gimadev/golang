package main

import "fmt"

func main() {

	numbers := [...]int {8, 10, 12, 90}

	numbers2 := [...]int {0, 1}

	res := []int {}

	i := 0
	j := 0

	for {

		if i == len(numbers) && j == len(numbers2) {
			break
		}

		if i < len(numbers) && j < len(numbers2) {
			if numbers[i] < numbers2[j] {
				res = append(res, numbers[i])
				i++
			} else {
				res = append(res, numbers2[j])
				j++
			}
		}

		if i < len(numbers) && j == len(numbers2) {
			res = append(res, numbers[i])
			i++
		}

		if i == len(numbers) && j < len(numbers2) {
			res = append(res, numbers2[j])
			j++
		}
			
	}

	fmt.Println(res)


}

