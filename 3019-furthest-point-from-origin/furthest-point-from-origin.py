class Solution:
    def furthestDistanceFromOrigin(self, moves: str) -> int:
        cnt_r, cnt_l = 0, 0
        for s in moves:
            if s == "R":
                cnt_r += 1
            elif s =="L":
                cnt_l += 1
            else:
                continue
        ans = len(moves) - 2 * min(cnt_r, cnt_l)
        return ans