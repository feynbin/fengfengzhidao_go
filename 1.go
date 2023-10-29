package main

func main() {
	twosum([]int{2, 7, 11, 15}, 9)
	twosum([]int{3, 2, 4}, 6)
	twosum([]int{3, 3}, 6)
}

func twosum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
