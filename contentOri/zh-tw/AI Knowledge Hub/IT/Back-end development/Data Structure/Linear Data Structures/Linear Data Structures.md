+++
title = "線性數據結構"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Linear Data Structures的問題
{{< /ask_chatgpt >}}



1. 請問何謂Linear Data Structures？其特色是什麼？
2. 請問Stack和Queue分別是哪種Linear Data Structures？它們的差異在哪裡？
3. 請問Linked List、Array和Vector都是哪種Linear Data Structures？它們的優缺點分別是什麼？
4. 請問Hash Table和Heap分別是哪種Linear Data Structures？它們的特色和使用情境分別是什麼？
5. 描述一下Binary Search Tree的特性及其常見操作。   

## 說明知識
{{< ask_chatgpt >}}
說明Linear Data Structures並舉例
{{< /ask_chatgpt >}}



線性資料結構是指其資料元素按照一定的順序排列，並且每個資料元素都只有一個前驅元素（第一個資料元素除外）和一個後繼元素（最後一個資料元素除外），即資料元素之間存在一對一的前後關係。常見的線性資料結構包括陣列，鏈表，佇列，堆棧等。

以下是幾種常見的線性資料結構：

1. 陣列：陣列是一個在內存中分配連續記憶體的資料結構，它通過索引來訪問和操作元素。陣列的特點是能夠快速訪問元素，但在插入和刪除操作時需要移動陣列中其他元素的位置。例如，int nums [5] = {1, 2, 3, 4, 5}。

2. 鏈表：鏈表是由節點組成的資料結構，每個節點包含資料和指向下一個節點的指針。鏈表的特點是在插入和刪除記錄時不需要移動其他元素位置，但是訪問元素時需要遍歷整個鏈表。例如，單鏈表、雙向鏈表、循環鏈表等。

3. 佇列：佇列是具有先進先出（FIFO）特點的資料結構，類似於排隊。在佇列的一端添加元素，在另一端刪除元素。例如，等待列。

4. 堆棧：堆棧是具有後進先出（LIFO）特點的資料結構，類似於一疊盤子，最後放上去的最先被取下來。在堆棧中添加元素的操作稱為推入(push)，刪除元素的操作稱為彈出(pop)。例如，瀏覽器的後退按鈕。   

## 彙總重點
{{< ask_chatgpt >}}
條列Linear Data Structures的重點
{{< /ask_chatgpt >}}



1. 線性結構：線性數據結構是數據元素按線性次序排列的結構，具有先後順序，且每個元素之間均具有一對一的直接前驅和直接後繼關係。

2. 數組Array：數組是一種最常見的線性數據結構，它可以連續存儲多個相同類型的數據，並可以通過下標快速訪問和修改數據。

3. 鏈表Linked List：鏈表是一種線性數據結構，它由一個個節點按順序連接而成，每個節點包含數據域和指向下一個節點的指針。

4. 堆棧Stack：堆棧是一種先進後出（LIFO）的線性數據結構，它支持push和pop操作，用於處理臨時性數據。

5. 隊列Queue：隊列是一種先進先出（FIFO）的線性數據結構，它支持enqueue和dequeue操作，用於實現消息佇列等場景。

6. 雙向鏈表Doubly Linked List： 雙向鏈表在鏈表的基礎上增加了prev指針，實現快速遍歷以及反向遍歷，但同時也帶來了額外的空間和時間複雜度開銷。

7. 循環鏈表Circular Linked List： 循環鏈表是一種特殊的鏈表，尾節點的next指向頭節點，形成一個循環。

8. 陣列堆積Array Heap： 堆是一種特殊的數據結構，它是一個完全二叉樹，且每個節點都大於或等於（或小於或等於）其子節點，常用於實現優先級佇列等場景。

9. 優先級佇列Priority Queue： 優先級佇列是一種特殊的隊列，每個元素都有一個優先級，優先級高的元素先出隊列。

10. 哈希表Hash Table： 哈希表是一種基於散列表實現的映射結構，它能夠實現高效的數據查詢和修改操作。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Linear Data Structures的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 將陣列中的元素循環向右移動k位。例如，當陣列為[1,2,3,4,5,6]且k=2時，陣列變為[5,6,1,2,3,4]。
答案：https://leetcode.com/problems/rotate-array/
2. 實現一個簡單的LIFO棧，支持推入、彈出和取頂操作。
答案：https://stackoverflow.com/questions/19141648/how-to-implement-a-stack-in-python
3. 找到陣列中缺少的最小正整數。例如，當陣列為[1,2,0]時，缺少的最小正整數為3。
答案：https://leetcode.com/problems/first-missing-positive/
4. 計算陣列中每個元素左側和右側的最大值之差的絕對值的和。例如，當陣列為[2,1,3,4,1]時，計算結果為8。左側最大值為[0,2,2,3,3]，右側最大值為[4,4,4,4,5]。
答案：https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
5. 實現一個字典數據結構，支持插入、查詢和刪除操作。
答案：https://github.com/grantjenks/python-diskcached   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Linear Data Structures的網路資料
{{< /ask_chatgpt >}}



1. "Introduction to Linear Data Structures" by GeeksforGeeks - https://www.geeksforgeeks.org/introduction-to-linear-data-structures/
This article gives an introduction to linear data structures including arrays, linked lists, stacks, and queues. It explains the basic concepts and operations of each data structure with examples.

2. "Linear Data Structures in Python" by Towards Data Science - https://towardsdatascience.com/linear-data-structures-in-python-3fe78f70aac8
This article describes how to implement linear data structures in Python, including arrays, linked lists, stacks, and queues. It provides sample code and explanations for each data structure.

3. "A Guide to Linear Data Structures in Java" by Baeldung - https://www.baeldung.com/java-linear-data-structures
This guide discusses linear data structures in Java, including arrays, linked lists, stacks, and queues. It covers the basic properties of each data structure, how to implement them in Java, and how to use them in real-world applications.

4. "Data Structures: Linear Data Structures" by Khan Academy - https://www.khanacademy.org/computing/computer-science/algorithms/arrays/a/linear-data-structures
This article introduces linear data structures and their properties, focusing on arrays and linked lists. It includes video explanations and interactive exercises.

5. "Linear Data Structures" by JavaTpoint - https://www.javatpoint.com/linear-data-structure
This article provides a comprehensive overview of linear data structures in Java, including arrays, linked lists, stacks, and queues. It covers the features, operations, and use cases of each data structure, along with Java code examples.   

