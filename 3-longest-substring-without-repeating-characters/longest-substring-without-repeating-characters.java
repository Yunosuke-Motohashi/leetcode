class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> lastIndex = new HashMap<>();
        int left = 0, length = 0;
        for (int i=0; i<s.length(); i++) {
            char c = s.charAt(i);
            int preIdx = lastIndex.getOrDefault(c, -1);
            if (left <= preIdx) {
                left = lastIndex.get(c) + 1;
            }
            lastIndex.put(c, i);
            length = Math.max(length, i - left + 1);

        }
        return length;
    }
}