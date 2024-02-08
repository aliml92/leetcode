package carfleet

import "sort"

func carFleet(target int, position []int, speed []int) int {
	result := 1

	size := len(position)
	if size == 1 {
		return result
	}
    
    pairSet := make([][2]int, size)
    for i:=0; i < size; i++ {
        pairSet[i] = [2]int{position[i], speed[i]}
    }
    
    sort.Slice(pairSet, func(i, j int) bool {
		return pairSet[i][0] > pairSet[j][0]
	})

    i, j := 0, 1
	ti := float32(target-pairSet[i][0])/float32(pairSet[i][1])
	for j < size {
		tj := float32(target-pairSet[j][0])/float32(pairSet[j][1])
		if tj <= ti {
		} else {
			result++
			i++
			ti = tj
		}
		j++
	} 

    return result    
}