class Solution {
    public int lengthOfLongestSubstring(String s) {
        Map<Character, Integer> dict = new HashMap<>();
        int left = 0, length = 0;
        for (int i=0; i<s.length(); i++) {
            char c = s.charAt(i);
            if (dict.containsKey(c) && left <= dict.get(c)) {
                left = dict.get(c) + 1;
            }
            dict.put(c, i);
            int tmpL = i - left + 1;
            if (tmpL > length) {
                length = tmpL;
            }

        }
        return length;
    }
}