class Solution {
    public int firstMissingPositive(int[] nums) {
    Set<Integer> check = new HashSet<>();
    for ( int num:nums){
       check.add(num) ;
    }

    for (int i = 1;i<=check.size();i++){
        if (!check.contains(i)){
            return i;
        }
      
        
    }
  return check.size()+1;
    }
}