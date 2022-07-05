# Given a sorted array of integers 'arr' & an integer 'target'
# Find the index of the first and last position of 'target' in 'arr'.
# If 'target' can't be found in 'arr', return [-1,-1]

# Linear Search
def first_last(target, arr):
    for i in range(len(arr)): # range(6) is 0 to 5
        if arr[i] == target:
            start = i
            while i+1 < len(arr) and arr[i+1] == target:
                i += 1
            return [start, i]
    return [-1, -1]
# T(n) = 0(n)
# S(n) = 0(1)


# Binary Search (Because list is sorted)
def find_start(target, arr):
    if arr[0] == target:
        return 0
    left, right = 0, len(arr)-1
    while left <= right:
        mid = (left+right)//2
        if arr[mid] == target and arr[mid-1] < target:
            return mid
        elif arr[mid] < target:
            left = mid+1
        else:
            right = mid-1
    return -1

def find_end(target, arr):
    if arr[0] == target:
        return 0
    left, right = 0, len(arr)-1
    while left <= right:
        mid = (left+right)//2
        if arr[mid] == target and arr[mid-1] > target:
            return mid
        elif arr[mid] > target:
            right = mid+1
        else:
            left = mid-1
    return -1

def first_last_binary(target, arr):
    if len(arr) == 0 or arr[0] > target or arr[-1] < target:
        return [-1, -1]
    start = find_start(arr, target)
    end = find_end(arr, target)
    return [start, end]
    # T(n) = 2 * 0(logn) = 0(logn)
    # S(n) = 0(1)
    

arr = [2,4,5,5,5,5,5,5,7,9,9]
target = 5
print(first_last(target, arr))