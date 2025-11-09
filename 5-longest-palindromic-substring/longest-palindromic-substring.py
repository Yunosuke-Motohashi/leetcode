
"""
回文の定義を考える
- 1文字の文字列は回文
- 2文字の文字列は同じ文字なら回文
- 3文字以上の文字列は、両端の文字が同じで、かつその内側の部分文字列が回文なら回文


ボトムアップ方式で記述
ボトムアップ = 小さい問題から大きい問題へ
Length = 2 から n までの長さの部分文字列を見ていく
先に長さが2の部分文字列を確認し、その後長さ3の部分文字列を確認する
Length = n-1の時、i:jの内側、つまりi+1:j-1の部分文字列が回文かどうか確認するためボトムアップである必要がある
"""

class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        dp = [[False] * n for _ in range(n)] # dp[i][j]: s[i:j+1]が回文かどうか
        for i in range(n):
            # 長さ1の部分文字列はすべて回文
            dp[i][i] = True
        max_length = 1
        start_idx = 0
        for length in range(2, n+1):
            # length: 部分文字列の長さ 2 から nまで
            for start in range(n-length+1):
                end = start + length -1
                # print(f"length: {length}, start: {start}, end: {end}", dp[start+1][end-1])
                # startからendまでの部分文字列s[start:end]が回文かどうかをチェック
                if s[start] == s[end]:
                    if length == 2:
                        # 長さが2の部分文字列で両端の文字が同じであれば回文
                        dp[start][end] = True
                    else:
                        # 長さが3以上の場合は、内部が回文かどうか確認
                        # 内部はstart+1 から end-1 まで
                        dp[start][end] = dp[start+1][end-1]
                # 回文であれば最大長さを更新
                if length > max_length and dp[start][end]:
                    max_length = length
                    start_idx = start
        # for d in dp: print(d)
        return s[start_idx:start_idx+max_length]
