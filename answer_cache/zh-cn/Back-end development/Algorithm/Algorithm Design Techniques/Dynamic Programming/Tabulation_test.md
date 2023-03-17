

1. Problem statement: 
Given an integer n, write a function to return the count of possible ways to climb n stairs. You can climb 1 or 2 stairs at a time.

Answer: 

def climb_stairs(n):
    if n == 1:
        return 1
    if n == 2:
        return 2
    table = [0]*(n+1)
    table[1] = 1
    table[2] = 2
    for i in range(3, n+1):
        table[i] = table[i-1] + table[i-2]
    return table[n]

Example input: climb_stairs(5)
Example output: 8

2. Problem statement: 
Given an array of integers, find the maximum sum subsequence where the subsequence contains no adjacent elements.

Answer: 

def max_sum_subseq(arr):
    n = len(arr)
    if n == 0:
        return 0
    if n == 1:
        return arr[0]
    table = [0]*n
    table[0] = arr[0]
    table[1] = max(arr[0], arr[1])
    for i in range(2, n):
        table[i] = max(table[i-2]+arr[i], table[i-1])
    return table[n-1]

Example input: max_sum_subseq([5, 1, 1, 5])
Example output: 10

3. Problem statement: 
Given a list of non-negative integers, find the maximum sum of a subsequence with the constraint that no two numbers in the sequence should be adjacent in the array.

Answer: 

def max_sum_nonadj(arr):
    n = len(arr)
    if n == 0:
        return 0
    if n == 1:
        return arr[0]
    table = [0]*n
    table[0] = arr[0]
    table[1] = max(arr[0], arr[1])
    for i in range(2, n):
        table[i] = max(table[i-2]+arr[i], table[i-1])
    return table[n-1]

Example input: max_sum_nonadj([3, 2, 5, 10, 7])
Example output: 15

4. Problem statement: 
Given an array arr of n integers, construct a Maximum Sum Subsequence of the given array where no two consecutive elements in the subsequence are adjacent in the given array.

Answer: 

def max_sum_nocnx(arr):
    n = len(arr)
    if n == 0:
        return []
    if n == 1:
        return [arr[0]]
    table = [0]*n
    table[0] = arr[0]
    table[1] = max(arr[0], arr[1])
    for i in range(2, n):
        table[i] = max(table[i-2]+arr[i], table[i-1])
    subseq = []
    i = n-1
    while i >= 0:
        if i == 0:
            subseq.append(arr[i])
            break
        elif i == 1:
            subseq.append(arr[i-1])
            break
        elif table[i]-table[i-2] == arr[i]:
            subseq.append(arr[i])
            i -= 2
        else:
            i -= 1
    subseq.reverse()
    return subseq

Example input: max_sum_nocnx([3, 2, 5, 10, 7])
Example output: [3, 5, 7]

5. Problem statement: 
Given a list of strings, return the largest word you can create by concatenating subsequent words together.

Answer: 

def largest_conc_word(words):
    n = len(words)
    if n == 0:
        return ''
    table = [[]]*(n+1)
    table[0] = ['']
    for i in range(1, n+1):
        maxlen = 0
        maxidx = 0
        for j in range(i):
            if len(words[j]) > maxlen and words[i-1].startswith(words[j]):
                maxlen = len(words[j])
                maxidx = j
        table[i] = table[maxidx] + [words[i-1]]
    return ''.join(max(table, key=len))

Example input: largest_conc_word(['cat', 'dog', 'catdog'])
Example output: 'catdog'