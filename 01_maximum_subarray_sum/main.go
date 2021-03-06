package maximum_subarray_sum

// The maximum sum subarray problem consists in finding the maximum sum of a contiguous
// subsequence in an array or list of integers:

// maxSequence [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// -- should be 6: [4, -1, 2, 1]

// Easy case is when the list is made up of only positive numbers and
// the maximum sum is the sum of the whole array.
// If the list is made up of only negative numbers, return 0 instead.

// Empty list is considered to have zero greatest sum.
// Note that the empty list or array is also a valid sublist/subarray.

func MaximumSubarraySumLoops(numbers []int) int {
	sum := 0

	var totNums = len(numbers)

	if totNums == 0 {
		return 0
	}

	for i := 0; i < totNums; i++ {
		if numbers[i] >= 0 {
			tempSum := 0
			for j := i; j < totNums; j++ {
				tempSum += numbers[j]
				if tempSum > sum {
					sum = tempSum
				}
			}
		}
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaximumSubarraySum(numbers []int) int {
	res, sum := 0, 0
	for i := range numbers {
		sum += numbers[i]
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}
