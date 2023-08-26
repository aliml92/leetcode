package containsduplicate217

// containsDuplicate returns true if nums contains
// any duplicates, false if there is not any.
func containsDuplicate(nums []int) bool {
	// create map
	set := make(map[int]struct{})
	// loop over nums slice
	for _, i := range nums {
		// check if map already has i as a key
		_, ok := set[i]
		// if yes, it means nums slice has duplicate
		if ok {
			// return immediately, no need to check the
			// rest of slice elements
			return true
		}
		// if map does not contain i, store it as a map key
		set[i] = struct{}{}
	}
	// if there is no duplicate found, return false
	return false
}
