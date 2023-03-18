+++
title = "表格法"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Tabulation的問題
{{< /ask_chatgpt >}}



1. 根據不同年齡層的財務目標，將支出分類，以便更好地了解該年齡層的花費習慣。
2. 分析哪些餐廳在不同的城市或地區經常被顧客點餐，以協助制定營銷策略。
3. 紀錄每位員工的出席率，早退率和遲到率，以協助管理人員確定員工出勤表現並提供相應的獎勵或懲罰。
4. 記錄一個公司的優惠券使用量，以解決促銷營銷策略的投入和回報。
5. 通过记录每位學生的出勤率、作业完成情况和考试成绩来衡量教育系统的效果，以判断学生的整体表现和制定针对性的教育改革计划。   

## 說明知識
{{< ask_chatgpt >}}
說明Tabulation並舉例
{{< /ask_chatgpt >}}



Tabulation（表格法）是一種動態程式設計方法，用於解決子問題互相不相交的問題。通常用於解決最短路徑、最長公共子序列、編輯距離等問題。

在Tabulation方法中，動態規劃的解決方案從最小的子問題開始，通過填充表格來解決更大的問題。表格的每個單元格代表問題的某個特定狀態，每個單元格的計算取決於其他單元格中已經計算的值（即子問題）。

舉例來說，考慮最短路徑問題。一個圖形可以用一個鄰接矩陣來表示，其中矩陣的每個元素代表一對節點之間的邊。如果兩個節點之間沒有邊，那麼該元素為無窮大。

步驟如下：

1. 創建一個二維表格來記錄每個節點的最短路徑。
2. 初始狀態下，表格中每個元素都是無窮大，除了起始節點的元素，它的值為0。
3. 遞歸計算表格中每個元素的值，該值表示從起始節點到該節點的最短路徑長度。
4. 當結束時，結果將存儲在表格的右下角元素中。

Tabulation方法可以大大簡化動態規劃問題的求解過程。它還可以更好地利用計算機內存，因為不需要存儲所有子問題的解答。   

## 彙總重點
{{< ask_chatgpt >}}
條列Tabulation的重點
{{< /ask_chatgpt >}}



1. Tabulation 是一種資料處理方法，將資料整理成表格形式，方便分析和理解。

2. Tabulation 可以用於統計數據、調查結果、市場調查等領域。

3. Tabulation 要求資料清晰、一致和完整，這樣才能正確地進行統計和分析。

4. Tabulation 的重點是資料的分類、標準化、整理和呈現。

5. Tabulation 的核心工具是電子表格軟件，如 MS Excel、Google Sheets 等。

6. Tabulation 可以產生各種圖表，如柱狀圖、折線圖、餅圖等，以更好地展示資料。

7. Tabulation 的應用範圍廣泛，可用於學術研究、商業分析、社會調查等不同領域。

8. 需要注意的是，Tabulation 只是一種分析工具，可以幫助分析資料，但不能替代對資料本身的理解和分析。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Tabulation的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Tabulation的網路資料
{{< /ask_chatgpt >}}



1. GDP by Country
Country | 2017 GDP (in Trillions) | 2018 GDP (in Trillions) | % Change
United States | 19.39 | 20.49 | 5.67%
China | 12.24 | 13.61 | 11.16%
Japan | 4.87 | 4.97 | 2.12%
Germany | 3.68 | 3.94 | 7.07%
United Kingdom | 2.62 | 2.63 | 0.38%

2. Smartphone Market Share by Brand
Brand | Q3 2021 Market Share | Q3 2020 Market Share | % Change
Samsung | 19.2% | 21.9% | -2.7%
Apple | 15.7% | 11.9% | 3.8%
Xiaomi | 13.9% | 10.5% | 3.1%
Oppo | 10.2% | 8.2% | 2.0%
Vivo | 9.0% | 8.0% | 1.0%

3. Top 5 Most Visited Websites
Website | Monthly Traffic (estimated) | Country of Origin
Google | 94.4 billion | United States
YouTube | 34.6 billion | United States
Facebook | 25.5 billion | United States
Baidu | 11.3 billion | China
Wikipedia | 8.6 billion | United States

4. Average Hourly Wages by Occupation
Occupation | Average Hourly Wage | % Change since 2019
Registered Nurse | $36.22 | 4.7%
Software Developer | $54.03 | 2.1%
Security Guard | $14.32 | 2.2%
Accountant | $28.22 | 1.8%
Construction Worker | $21.39 | 3.5%

5. Global Carbon Emissions by Country
Country | 2019 Emissions (in million metric tons) | % of Global Total
China | 10,170 | 29.7%
United States | 5,131 | 14.9%
India | 2,654 | 7.7%
Russia | 1,711 | 5.0%
Japan | 1,244 | 3.6%   

