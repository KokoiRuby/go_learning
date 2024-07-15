package main

import "fmt"

func main() {
	//var matrix [2][3]int
	//matrix = [2][3]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//}
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(&matrix[0][0], &matrix[0][1], &matrix[0][2])
	fmt.Println(&matrix[1][0], &matrix[1][1], &matrix[1][2])

	for i, v1 := range matrix {
		for j, v2 := range v1 {
			fmt.Printf("matrix[%v][%v] = %v \n", i, j, v2)
		}
	}

	numbers := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	row3 := []int{7, 8, 9}

	numbers = append(numbers, row1, row2, row3)

	for i := range numbers {
		fmt.Println(numbers[i])
	}

}
