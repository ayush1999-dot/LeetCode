func findDuplicates(nums []int) []int {
check:=make(map[int]bool)
arr:= []int{}
for _,num := range nums{
    if check[num]{
        arr = append(arr,num)
    }
    check[num]= true

}
return arr
}