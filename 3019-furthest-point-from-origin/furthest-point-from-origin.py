class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        cnt_r, cnt_l = moves.count("R"), moves.count("L")
        ans = len(moves) - 2 * min(cnt_r, cnt_l)
        return ans