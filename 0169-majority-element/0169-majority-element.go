func majorityElement(nums []int) int {
    count:=0
    num:=-1
    for  i:=0;i<len(nums);i++ {
        if count ==0{
            num = nums[i]
            count++
        }else {
            if nums[i]== num{
                count++
            }else{
                count--
            }
        }
    }
        if count<1{
            return -1
        }
        count =0
        for  i:=0;i<len(nums);i++{
            if num == nums[i]{
                count++
            }
        }
        if count >len(nums)/2{
            return num
        }else{
            return -1
        }

    }

