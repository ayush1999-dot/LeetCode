class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        Set<Integer> unique = new HashSet<>();
        List<Integer> dupli = new ArrayList<>();    
        for(int num:nums){
            if (unique.contains(num)){
                dupli.add(num);
            }
            unique.add(num);

        }
        return dupli;
    }
}