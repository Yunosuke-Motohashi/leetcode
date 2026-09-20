import sys

sys.setrecursionlimit(10**6)  # 再帰の深さを増やしておく


class Solution:

    def containsCycle(self, grid: list[list[str]]) -> bool:
        self.n, self.m = len(grid), len(grid[0])
        self.visited = [[False for _ in range(self.m)] for _ in range(self.n)]

        def dfs(i, j, pre_i, pre_j):
            now = grid[i][j]
            self.visited[i][j] = True
            for di, dj in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                next_i, next_j = i + di, j + dj
                if (
                    0 <= next_i < self.n
                    and 0 <= next_j < self.m
                    and now == grid[next_i][next_j]
                ):
                    if (next_i, next_j) == (pre_i, pre_j):  # 直前はskip
                        continue
                    if self.visited[next_i][next_j]:
                        # もう来ていた = 閉路発見
                        return True
                    if dfs(next_i, next_j, i, j):
                        # 再帰的にDFSして閉路が見つかったらtrue返す
                        return True
            return False

        for i in range(self.n):
            for j in range(self.m):
                if not self.visited[i][j]:
                    if dfs(i, j, -1, -1):
                        return True

        return False
