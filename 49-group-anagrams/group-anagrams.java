class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        Map<String, List<String>> strMap = new HashMap<>();
        for (String s : strs) {
            char[] chars = s.toCharArray();
            Arrays.sort(chars);
            String sortedS = new String(chars);
            if (!strMap.containsKey(sortedS)) {
                strMap.put(sortedS, new ArrayList<>());
            }
            strMap.get(sortedS).add(s);
        }
        return new ArrayList<>(strMap.values());
    }
}