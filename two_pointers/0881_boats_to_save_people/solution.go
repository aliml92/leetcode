package boatstosavepeople

import "slices"


func numRescueBoats(people []int, limit int) int {
	var result int
	
    n := len(people)
    if n == 1 {
        return 1
    }
	
    slices.Sort(people)
	
    i, j := 0, n-1
	for i < j {
		if people[i] + people[j] > limit {
			j--
		} else {
			i++
			j--
		}
		result++
	}

    if i == j {
        result++
    }

	return result
}