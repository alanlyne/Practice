# Given 2 strings s1 & s2, chick if they're anagrams.
# Two strings are anagrams if they're made of the same characterss with the same frequencies

from collections import Counter


def are_anagrams(s1, s2):
    if len(s1) != len(s2):
        return False

    # s 0(n)
    freq1 = {}
    # s 0(n)
    freq2 = {}

    # n 0(1)
    for ch in s1: 
        if ch in freq1:
            freq1[ch] += 1
        else:
            freq1[ch] = 1
    
    # n 0(1)
    for ch in s2:
        if ch in freq2:
            freq2[ch] += 1
        else:
            freq2[ch] = 1

    # n 0(1)
    for key in freq1:
        if key not in freq2 or freq1[key] != freq2[key]:
            return False

    # T(n) = 0(n) + 0(n) + 0(n)
    # T(n) = 0(n)
    # S(n) = 0(n) + 0(n)
    # S(n) = 0(n)
    return True

def are_anagrams_import(s1, s2):
    if len(s1) != len(s2):
        return False
    return Counter(s1) == Counter(s2)

def are_anagrams_sort(s1, s2):
    if len(s1) != len(s2):
        return False
    #0(nlogn) + n + 0(nlogn) = T(n) = 0(nlogn)
    return sorted(s1) == sorted(s2)


s1 = "dog"
s2 = "god"

print(are_anagrams(s1, s2))

