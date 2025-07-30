class NumArray {
    int[] n;

    public NumArray(int[] nums) {
    n= new int[nums.length];
    int sum =nums[0];
    n[0]=nums[0];
    for (int i =1 ; i<nums.length;i++){
        sum= sum + nums[i];
        n[i]= sum;
     } 
     
    }
    
    public int sumRange(int left, int right) {
        if (left==0){
            return n[right];
        }
        return n[right]-n[left-1];
    }
}

/**
 * Your NumArray object will be instantiated and called as such:
 * NumArray obj = new NumArray(nums);
 * int param_1 = obj.sumRange(left,right);
 */