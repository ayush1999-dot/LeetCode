func singleNumber(nums []int) int {
	mapnums := make(map[int]int,)

	for _, v := range nums {
mapnums[v]+= 1
	}

	for key, v := range mapnums {
		if v == 1 {
			return key
		}
	}

	return 0

}