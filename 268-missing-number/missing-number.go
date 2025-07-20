func missingNumber(nums []int) int {
    n:= len(nums)
    sum:= (n*(n+1))/2
    temp:= 0
    for _,v:=range nums {
        temp+= v
    }
    return sum-temp
}