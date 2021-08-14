package main

func findSubarrayWithMaxSum(data []int) [2]int {
	subarrayMaxSum := data[0]
	subarrayLeftLimit := 0
	subarrayRightLimit := 0
	partialSum := 0
	minusPosition := -1

	for key, value := range data {
		partialSum += value

		if partialSum > subarrayMaxSum {
			subarrayMaxSum = partialSum
			subarrayLeftLimit = minusPosition + 1
			subarrayRightLimit = key
		}

		if partialSum < 0 {
			partialSum = 0
			minusPosition = key
		}
	}

	result := [2]int{subarrayLeftLimit, subarrayRightLimit}

	return result
}
