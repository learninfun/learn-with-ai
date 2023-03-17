## 習題預習
{{< ask_chatgpt >}}
給我5題Stack的問題
{{< /ask_chatgpt >}}



1. 什麼是Stack？
2. 如何實現一個Stack？
3. 如何在Stack中添加和移除元素？
4. 如何實現Stack中的迭代器？
5. 如何避免Stack中的內存泄漏？   

## 說明知識
{{< ask_chatgpt >}}
說明Stack並舉例
{{< /ask_chatgpt >}}



Stack是一種資料結構，通常用於儲存按照順序排列的數據，並且只允許在其頂部添加和刪除數據，並且是以先進後出（LIFO）的方式操作的。

舉個例子，比如栈中通常用於存放函數的返回地址以及參數，先進後出的特點可以讓我們更好的處理函數的遞迴調用。

假如我們現在要實現一個簡單的計算器程式，可以接收輸入的數字和運算符號，並且可以計算出結果。在這裡，我們使用Stack來實現計算器的計算過程。

當用戶輸入一個數字時，我們可以把它放入Stack中，等到用戶輸入一個運算符號時，我們可以從Stack中取出最近的兩個數字，然後進行計算，然後再把計算結果放回到Stack中。這樣，當用戶輸入了足夠的數字和運算符號後，最後Stack中就只剩下一個值，就是計算結果了。   

## 彙總重點
{{< ask_chatgpt >}}
條列Stack的重點
{{< /ask_chatgpt >}}



1. Stack 是一種資料結構，符合先進後出的原則。
2. 只有在堆疊的頂部（也就是最後一個新增的元素）才能夠進行插入和刪除操作。
3. 堆疊的初始狀態是空的，當堆疊中沒有元素時，我們稱之為空堆疊。
4. push() 函式用於將元素推入堆疊的頂部，pop() 函式用於從堆疊的頂部移除元素。
5. Stack 可以透過陣列或是鏈表來實現。
6. Stack 常用於函式調用時的遞迴操作，還有編譯器編譯時的符號表管理等方面。
7. 在資料結構中，Stack 是一個基礎的結構，許多其他資料結構都可以透過 Stack 來實現，例如 Queue、表達式求值等。
8. Stack 的時間複雜度為 O(1)。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Stack的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 包含min方法的Stack
實現一個包含min方法的Stack，除了常規的push、pop方法外，還要有返回stack內最小元素的方法min。要求：所有操作均為O(1)時間複雜度。

答案：https://leetcode.com/problems/min-stack/

2. 使用Stack實現Queue
使用Stack實現一個Queue。

答案：https://leetcode.com/problems/implement-queue-using-stacks/

3. 簡化路徑
給定一個表示絕對路徑的Unix風格的字符串，請簡化它。例如，路徑"/home//foo/"應簡化為"/home/foo"。

答案：https://leetcode.com/problems/simplify-path/

4. 小括號匹配
給定一個只包含"(", ")", "{", "}", "[", "]"的字符串，請判斷是否是合法的小括號串。要求：快速、高效且正確。

答案：https://leetcode.com/problems/valid-parentheses/

5. 最大矩形
給定一個只包含0和1的矩陣，找出其中最大的由1組成的矩形面積。要求：時間複雜度為O(mn)。

答案：https://leetcode.com/problems/maximal-rectangle/   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Stack的網路資料
{{< /ask_chatgpt >}}



1. What Is Stack in Software Development?
https://www.upgrad.com/blog/what-is-stack-in-software-development/

2. Understanding Stack Data Structure with Examples
https://www.javatpoint.com/stack-data-structure

3. The Stack in C++
https://www.geeksforgeeks.org/stack-in-cpp-stl/

4. Stack Overflow: understanding the website where developers learn to code
https://www.techradar.com/news/stack-overflow-understanding-the-website-where-developers-learn-to-code

5. A Beginner's Introduction to Stack Data Structures and Its Implementation in Python
https://www.analyticsvidhya.com/blog/2020/03/what-is-stack-data-structure-python-implementation/   

