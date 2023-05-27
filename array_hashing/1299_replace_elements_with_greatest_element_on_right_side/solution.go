package replaceelementswithgreatestelementonrightside1299


func replaceElements(arr []int) []int {
    max := -1
	for i := len(arr)-1; i >=0; i-- {
		temp := arr[i]
		arr[i] = max
		max = maxInt(max, temp)
	}
	return arr
}

func maxInt(x,y int) int {
	if x > y {
		return x
	}
	return y
}