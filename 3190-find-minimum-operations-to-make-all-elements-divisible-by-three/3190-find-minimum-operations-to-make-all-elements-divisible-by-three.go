func minimumOperations(nums []int) int {
    sum:=0
    for _,v:=range nums{
        if v%3 !=0{
            sum++
        }
    }
    return sum
}