'''
时间复杂度O（nlog(n)）,n为区间数量，除去排序的开销，只需要一次线性扫描，主要时间开销是排序的O(nlogn)

空间复杂度O（logn）,n为区间数量，计算的是存储答案之外，使用的额外空间。

'''

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: x[0])

        merged = []
        for interval in intervals:
            # 如果列表为空，或者当前区间与上一区间不重合，直接添加
            if not merged or merged[-1][1] < interval[0]:
                merged.append(interval)
            else:
                # 否则的话，我们就可以与上一区间进行合并
                merged[-1][1] = max(merged[-1][1], interval[1])

        return merged

