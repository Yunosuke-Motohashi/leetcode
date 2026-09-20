class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        cnt_r, cnt_l, cnt__ = moves.count("R"), moves.count("L"), moves.count("_")
        ans = cnt__ + abs(cnt_r-cnt_l)
        return ans