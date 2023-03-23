+++
title = "堆栈"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Stack的问题
{{< /ask_chatgpt >}}



1. 什麼是Stack？
2. 如何實現一個Stack？
3. 如何在Stack中添加和移除元素？
4. 如何實現Stack中的迭代器？
5. 如何避免Stack中的內存泄漏？   

## 说明知识
{{< ask_chatgpt >}}
说明Stack并举例
{{< /ask_chatgpt >}}



Stack是一種資料結構，通常用於儲存按照順序排列的數據，並且只允許在其頂部添加和刪除數據，並且是以先進後出（LIFO）的方式操作的。

舉個例子，比如栈中通常用於存放函數的返回地址以及參數，先進後出的特點可以讓我們更好的處理函數的遞迴調用。

假如我們現在要實現一個簡單的計算器程式，可以接收輸入的數字和運算符號，並且可以計算出結果。在這裡，我們使用Stack來實現計算器的計算過程。

當用戶輸入一個數字時，我們可以把它放入Stack中，等到用戶輸入一個運算符號時，我們可以從Stack中取出最近的兩個數字，然後進行計算，然後再把計算結果放回到Stack中。這樣，當用戶輸入了足夠的數字和運算符號後，最後Stack中就只剩下一個值，就是計算結果了。   

## 汇总重点
{{< ask_chatgpt >}}
条列Stack的重点
{{< /ask_chatgpt >}}



1. Stack 是一種資料結構，符合先進後出的原則。
2. 只有在堆疊的頂部（也就是最後一個新增的元素）才能夠進行插入和刪除操作。
3. 堆疊的初始狀態是空的，當堆疊中沒有元素時，我們稱之為空堆疊。
4. push() 函式用於將元素推入堆疊的頂部，pop() 函式用於從堆疊的頂部移除元素。
5. Stack 可以透過陣列或是鏈表來實現。
6. Stack 常用於函式調用時的遞迴操作，還有編譯器編譯時的符號表管理等方面。
7. 在資料結構中，Stack 是一個基礎的結構，許多其他資料結構都可以透過 Stack 來實現，例如 Queue、表達式求值等。
8. Stack 的時間複雜度為 O(1)。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Stack的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 包含min方法的Stack
实现一个包含min方法的Stack，除了常规的push、pop方法外，还要有返回stack内最小元素的方法min。要求：所有操作均为O(1)时间复杂度。

答案：https://leetcode.com/problems/min-stack/

2. 使用Stack实现Queue
使用Stack实现一个Queue。

答案：https://leetcode.com/problems/implement-queue-using-stacks/

3. 简化路径
给定一个表示绝对路径的Unix风格的字符串，请简化它。例如，路径"/home//foo/"应简化为"/home/foo"。

答案：https://leetcode.com/problems/simplify-path/

4. 小括号匹配
给定一个只包含"(", ")", "{", "}", "[", "]"的字符串，请判断是否是合法的小括号串。要求：快速、高效且正确。

答案：https://leetcode.com/problems/valid-parentheses/

5. 最大矩形
给定一个只包含0和1的矩阵，找出其中最大的由1组成的矩形面积。要求：时间复杂度为O(mn)。

答案：https://leetcode.com/problems/maximal-rectangle/   

