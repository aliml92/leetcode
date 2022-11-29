package groupanagrams49


func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)
	for _, s := range strs {
		var count [26]int
		for _, c := range s {
			count[int(c)-int('a')]++
		}
		m[count] = append(m[count], s) 
	}
	var r [][]string
	for _, v := range m {
		r =append(r, v)
	}
	return r
}