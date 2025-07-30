func subarraySum(nums []int, k int) int {
count :=0
prefix:=0
sumMap :=make(map[int]int)
sumMap[0]=1
for i:=0;i<len(nums);i++{
    prefix +=nums[i]

    if val,ok:= sumMap[prefix-k];ok{
        count+=val
    }
    sumMap[prefix]+=1
}
return count
}