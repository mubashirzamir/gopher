package arrays

func Sum(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}

	return sum
}

func SumAll(arrays ...[]int) []int {
	result := []int{}
	for _, arr := range arrays {
		result = append(result, Sum(arr))
	}

	return result
}

func SumAllTails(arrays ...[]int) []int {
	beheadedArrays := [][]int{}

	for _, arr := range arrays {
		beheaded := []int{}
		if len(arr) > 0 {
			beheaded = arr[1:]
		}

		beheadedArrays = append(beheadedArrays, beheaded)
	}

	return SumAll(beheadedArrays...)
}
