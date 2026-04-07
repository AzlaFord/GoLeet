package twosum

func TwoSum(arr []int, target int) []int {

	m := make(map[int]int)
	for i, num := range arr {

		complement := target - num
		_, ok := m[complement]
		if ok {
			return []int{m[complement], i}
		} else {
			m[num] = i
		}
	}
	return nil
}
