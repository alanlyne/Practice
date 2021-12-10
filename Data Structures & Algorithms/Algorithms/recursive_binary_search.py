def recursive_binary_search(list, target):
    if len(list) == 0:
        return False
    else:
        midpoint = (len(list))//2

        if list[midpoint] == target:
            return True
        else:
            if list[midpoint] < target:
                return recursive_binary_search(list[midpoint+1:], target) # midpoint+1: From midpoint+1 to the end of the list 
            else:
                return recursive_binary_search(list[:midpoint], target) # :midpoint From the start of the list to the index before midpoint

def verify(result):
    print("Target found :", result)

numbers = [1,2,3,4,5,6,7,8,9,10]

result = recursive_binary_search(numbers, 12)
verify(result)

result = recursive_binary_search(numbers, 3)
verify(result)