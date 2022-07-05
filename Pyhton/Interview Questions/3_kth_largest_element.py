# Given an array of integers 'arr' and an integer 'k'
# Find the 'k'th largest element

import heapq


def kth_largest(arr, k):
    for i in range(k-1):
        arr.remove(max(arr))
    return max(arr)
# T(n, k) = (k-1)*2n+n
# T(n, k) = 2kn-n = 0(kn)
# S(n) = 0(1)


def kth_largest_sort(arr, k):
    n = len(arr)
    arr.sort()
    return arr[n-k]
# T(n,k) = 0(nlogn)+0(1) = 0(nlogn)


def kth_largest_heap(arr, k):
    arr = [-elem for elem in arr]   # n
    heapq.heapify(arr)              # n
    for i in range(k-1):            # k-1
        heapq.heappop(arr)          # logn
    return -heapq.heappop(arr)      # logn
# T(n,k) = 2n + (k-1) * logn + logn
# T(n,k) = 2n + klogn = 0(n + klogn)
# S(n,k) = 0(n)

arr = [4,2,9,7,5,6,7,1,3]
k = 4

