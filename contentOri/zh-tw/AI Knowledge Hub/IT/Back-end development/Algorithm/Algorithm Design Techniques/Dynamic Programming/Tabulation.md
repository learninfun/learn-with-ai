+++
title = "表格法"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Tabulation的中文問題
{{< /ask_chatgpt >}}



1. 根據不同年齡層的財務目標，將支出分類，以便更好地了解該年齡層的花費習慣。
2. 分析哪些餐廳在不同的城市或地區經常被顧客點餐，以協助制定營銷策略。
3. 紀錄每位員工的出席率，早退率和遲到率，以協助管理人員確定員工出勤表現並提供相應的獎勵或懲罰。
4. 記錄一個公司的優惠券使用量，以解決促銷營銷策略的投入和回報。
5. 通过记录每位學生的出勤率、作业完成情况和考试成绩来衡量教育系统的效果，以判断学生的整体表现和制定针对性的教育改革计划。   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Tabulation並舉例
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
以中文條列Tabulation的重點
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
以中文給我5題Tabulation的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1.问题：斐波那契数列的第 n 项是什么？
答案：1,1,2,3,5,8,13，...

2.问题：给定一个整数数组和一个目标值，找到数组中和为目标值的两个数字的索引。
答案：[0,1]

3.问题：给定一个非空字符串 s 和一个字典 wordDict ，判断 s 是否可以被空格拆分成一个或多个在字典中出现的单词。
答案：True

4.问题：给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
答案：6 （从第一位开始，连续的子数组可为 [−2,1,−3,4,-1,2,1,-5,4]，最大连续子数组为 [4,-1,2,1]，其和为 6）

5.问题：给定两个单词 word1 和 word2，找到使得 word1 转换成 word2 所需的最少操作数。你可以对一个单词进行如下三种操作：插入一个字符，删除一个字符，替换一个字符。
答案：3   

