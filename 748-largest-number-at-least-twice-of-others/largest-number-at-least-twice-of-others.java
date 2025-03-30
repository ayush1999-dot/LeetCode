class Solution {
    public int dominantIndex(int[] nums) {
        if (nums.length <2){
            return -1;
        }
        int max = nums[0],secondmax=-1,count =0,indexof = 0;
        for(int i = 1 ; i < nums.length; i++){
             if (nums[i]>max) {
                 indexof = i;
                secondmax = max;
                max = nums[i];
                count=0;
            }else if(nums[i]==max){
                 count++;
             }else if(nums[i]>secondmax){
                secondmax = nums[i];
            }
            
        }
        if(count==1){
            return -1;
            
        } else if (secondmax*2 > max) {
            return -1;
            
        }else {
            return indexof;
        }
    
    }
}