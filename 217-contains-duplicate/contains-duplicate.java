class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashSet<Integer> duplicate_check = new HashSet<>();
        for (int num :nums){
            if(duplicate_check.contains(num)){
                return true;// Found Duplicate
            }
            duplicate_check.add(num);
            
        }
        return false;
    }
}