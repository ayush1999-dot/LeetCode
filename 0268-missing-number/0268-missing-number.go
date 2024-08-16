func missingNumber(nums []int) int {
var sum1 int
    for i,v := range nums{
        sum1 +=(i+1)-v
    }
    return sum1
}