package numberof1bits191


func hammingWeight(num uint32) int {
    var result int
    var mask uint32 = 1
    for num != 0 {
        if num & mask == 1 {
            result++
        }
        num >>= 1        
    }
    return result
}