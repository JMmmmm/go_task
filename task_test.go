package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubarrayWithMaxSum(t *testing.T) {
	assert := assert.New(t)

	testData := []int{10, -3, -12, 8, 42, 1, -7, 0, 3}

	assert.Equal(findSubarrayWithMaxSum(testData), [2]int{3, 5})
}
