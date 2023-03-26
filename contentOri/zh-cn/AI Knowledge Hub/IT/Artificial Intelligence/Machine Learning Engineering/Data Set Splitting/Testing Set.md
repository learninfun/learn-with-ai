+++
title = "测试集"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Testing Set的中文问题
{{< /ask_chatgpt >}}

1. 這本書的作者是誰？
2. 這張照片是在哪裡拍的？
3. 這道菜的主要材料有哪些？
4. 這個城市有哪些著名的旅遊景點？
5. 你認識那個人嗎？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Testing Set并举例
{{< /ask_chatgpt >}}

Testing Set是指在機器學習中，用來驗證模型的一個獨立數據集，它獨立於訓練集和驗證集，一般用於評估模型的泛化能力。在進行模型訓練時，訓練集用來學習模型參數，驗證集用來調整超參數，而Testing Set則用來測試模型的真正功能。

例如，我們要訓練一個圖像識別模型，我們可以將數據集分為3個部分：訓練集、驗證集和測試集。訓練集用來訓練模型，驗證集用來調整模型參數，而最後的Testing Set則用來評估模型的準確度。我們可以從數據集中隨機選擇一定比例的數據作為Testing Set，並且保證Testing Set與訓練集和驗證集獨立。最後，我們可以利用Testing Set來測試模型的準確度，並且對模型進行優化和改進。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Testing Set的重点
{{< /ask_chatgpt >}}

1. Testing Set是机器学习中测试模型效能的重要工具。
2. Testing Set必须与Training Set是互相独立且具有相同概率分布的资料集。
3. Testing Set要能够涵盖足够广泛且多元的资料，以确保模型能够在各种情况下都有良好的表现。
4. Testing Set的样本数量应足够大，以测试模型的泛化能力。
5. Testing Set的准确度、精度、召回率等指标能够有效地评估模型的表现。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Testing Set的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}

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

