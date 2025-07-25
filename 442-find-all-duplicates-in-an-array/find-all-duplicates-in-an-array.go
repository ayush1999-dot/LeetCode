func findDuplicates(nums []int) []int {
check:=make([]bool, len(nums))
arr:= []int{}
for  i := 0 ;i<len(nums);i++{
    if check[nums[i]-1]{
        arr = append(arr,nums[i])
    }
    check[nums[i]-1]= true

}
return arr
}