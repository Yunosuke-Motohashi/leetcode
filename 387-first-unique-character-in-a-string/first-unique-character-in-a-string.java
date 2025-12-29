class Solution {
    public int firstUniqChar(String s) {
        Map<Character, Integer> countMap = new HashMap<>();
        for (int i=0; i<s.length(); i++) {
            char ci = s.charAt(i);
            if (countMap.containsKey(ci)) {
                countMap.put(ci, countMap.getOrDefault(ci, 0) + 1);
            } else {
                countMap.put(ci, 1);
            }
        }

        int ansIdx = -1;
        for (int i=0; i<s.length(); i++) {
            char ci = s.charAt(i);
            if (countMap.get(ci) == 1) {
                ansIdx = i;
                break;
            }
        }
        return ansIdx;
    }
}