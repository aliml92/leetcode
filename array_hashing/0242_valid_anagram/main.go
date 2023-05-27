package validanagram242


func isAnagram(s string, t string) bool {
	smap := make(map[byte]int)
	tmap := make(map[byte]int)
	
	if len(s) != len(t){
		return false
	}
	for i := 0; i < len(s); i++ {
		sv, ok := smap[s[i]]
		if !ok{
			sv = 0
		}
		smap[s[i]] = 1 + sv 
		tv, ok := tmap[t[i]]
		if !ok {
			tv = 0
		}
		tmap[t[i]] = 1 + tv 

	}
	for k, v := range smap{
		tv, ok := tmap[k]
		if ok {
			if v != tv {
				return false
			} 
		} else {
			return false
		}
	}
	return true

}