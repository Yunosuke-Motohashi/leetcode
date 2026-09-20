# import string
# class Solution:
#     def reverseDegree(self, s: str) -> int:
#         # print(string.ascii_lowercase)
#         a_z_idx = {}
#         for i, v in enumerate(string.ascii_lowercase):
#             a_z_idx[v] = 26-i
#         # print(a_z_idx)
#         ans = 0
#         for i, v in enumerate(s):
#             # print(ans, i+1, a_z_idx[v])
#             ans += (i+1) * a_z_idx[v]

#         return ans

class Solution:
    def reverseDegree(self, s: str) -> int:
        # print(string.ascii_lowercase)
        # a_z_idx = {}
        # for i, v in enumerate(string.ascii_lowercase):
        #     a_z_idx[v] = 26-i
        # print(a_z_idx)
        ans = 0
        for i in range(len(s)):
            s_value = ord("z") - ord(s[i]) +1
            ans += (i+1) * s_value
            # print(ans, i+1, s_value)


        return ans
