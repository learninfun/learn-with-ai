1. 请写一个函数，判断一个字符串是否为回文（回文）。
2. 请写一个函数，将一个数组中的所有偶数移到前面，所有奇数移到后面。
3. 请写一个函数，将一个给定的整数转化为罗马数字。
4. 请写一个函数，判断一个整数是否为质数。
5. 请写一个函数，返回一个字符串中最长的不重复子串的长度。

答案：
1. 
```python
def is_palindrome(s):
    return s == s[::-1]
```
2. 
```python
def even_odd_sort(nums):
    left = 0
    right = len(nums) - 1
    while left < right:
        while left < right and nums[left] % 2 == 0:
            left += 1
        while left < right and nums[right] % 2 != 0:
            right -= 1
        if left < right:
            nums[left], nums[right] = nums[right], nums[left]
            left += 1
            right -= 1
    return nums
```
3. 
```python
def int_to_roman(num):
    romans = [
        (1000, 'M'), (900, 'CM'), (500, 'D'), (400, 'CD'),
        (100, 'C'), (90, 'XC'), (50, 'L'), (40, 'XL'),
        (10, 'X'), (9, 'IX'), (5, 'V'), (4, 'IV'),
        (1, 'I')
    ]
    roman_num = ''
    for value, letter in romans:
        while num >= value:
            roman_num += letter
            num -= value
    return roman_num
```
4. 
```python
def is_prime(num):
    if num < 2:
        return False
    for i in range(2, int(num**0.5)+1):
        if num % i == 0:
            return False
    return True
```
5. 
```python
def longest_substring(s):
    char_dict = {}
    start = max_len = 0
    for i, c in enumerate(s):
        if c in char_dict and char_dict[c] >= start:
            start = char_dict[c] + 1
        else:
            max_len = max(max_len, i - start + 1)
        char_dict[c] = i
    return max_len
```