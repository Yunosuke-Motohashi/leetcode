// import java.util.Arrays;

class Solution {
    public boolean containsDuplicate(int[] nums) {
        // System.out.println(Arrays.toString(nums));
        Map<Integer, Integer> count = new HashMap<>();
        boolean ans = false;
        for (int i=0; i<nums.length; i++) {
            int ni = nums[i];
            if (count.containsKey(ni)) {
                ans = true;
                break;
            } else {
                count.put(ni, 1);
            }
        }
        return ans;
    }
}