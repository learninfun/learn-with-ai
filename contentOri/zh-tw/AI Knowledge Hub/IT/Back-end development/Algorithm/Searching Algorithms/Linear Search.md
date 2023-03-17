## 習題預習
{{< ask_chatgpt >}}
給我5題Linear Search的問題
{{< /ask_chatgpt >}}



1. 在一個有序的整數陣列中，找尋目標數字。
2. 從一個字串陣列中，找到第一個出現在指定位置的字串。
3. 在一個雙精度浮點數陣列中，找到所有小於指定數值的數字。
4. 從一個物件陣列中，找到第一個符合條件的物件。
5. 在一個布林值陣列中，找到第一個為真的元素。   

## 說明知識
{{< ask_chatgpt >}}
說明Linear Search並舉例
{{< /ask_chatgpt >}}



Linear Search是一种基本的搜索算法，也叫顺序查找。它从序列的开头开始逐个查找，直到找到目标元素或者查找完整个序列为止。

例如，假设我们要查找一个数字3是否在一个给定的数组中，可以按以下过程实现线性搜索：

1.从数组的第一个元素开始遍历，即下标为0。

2.比较当前元素是否等于目标元素3。

3.如果相等，即找到目标元素，返回当前元素的下标。

4.如果不相等，继续遍历下一个元素。

5.如果全部元素遍历完后仍未找到目标元素，返回-1，表示未找到。

举例来说，假设给定一个数组arr={2,5,8,3,9,4}，要查找数字3是否在其中，那么按照以上过程，我们可以依次比较2、5、8、3，找到了目标元素3，返回它的下标3。

实际上，线性搜索算法可以应用于任何有序或无序的列表或数组中。但是，相对于其他高级的搜索算法，它的效率较低，因为它需要遍历整个序列才能找到目标元素。如果列表中有大量元素，这将会消耗大量的时间。   

## 彙總重點
{{< ask_chatgpt >}}
條列Linear Search的重點
{{< /ask_chatgpt >}}



1. Linear Search是一種簡單的搜尋演算法，也被稱為線性搜尋。

2. 線性搜尋演算法是一種線性時間複雜度的演算法，其時間複雜度為O(n)。

3. 線性搜尋演算法的基本原理是從頭到尾依次搜索待搜尋的元素，直到找到目標元素為止。

4. 線性搜尋演算法的優點是簡單易懂，容易實現，並且不需要事先對數據進行排序。

5. 線性搜尋演算法的缺點是當數據量很大時，處理速度會非常慢，因為需要遍歷整個數組或列表。

6. 在實際應用中，線性搜尋演算法通常用於小型數據集的查找，對於大型數據集和應用場景，應該使用更高效的搜尋算法。

7. 線性搜尋演算法的實現方式包括遞歸和迭代兩種方法，一般選擇迭代方法較為常見。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Linear Search的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 有一個字元陣列，要求在其中找到第一個不是英文字母的字元。如果找不到，回傳-1。

Input: {'a', 'b', '+', 'd', 'E', 'f'}
Output: 2

2. 有一個整數陣列，每個數字都出現兩次，只有一個數字出現一次。找到這個數字。

Input: {1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1}
Output: 6

3. 有一個正整數陣列，找到其中最小的缺失數字。也就是說，若陣列中有1、3、4，則回傳2。

Input: {1, 3, 5, 6, 7, 9}
Output: 2

4. 有一個整數陣列，找到其中第一個出現超過一次的數字。

Input: {1, 2, 3, 4, 5, 2, 3, 6, 7, 7}
Output: 2

5. 有一個字串陣列，找到其中最長的字串。

Input: {"cat", "dog", "elephant", "bird", "frog"}
Output: "elephant"

答案：

1. 

int findNonAlpha(char[] arr) {
    for(int i=0; i<arr.length; i++) {
        if(!Character.isLetter(arr[i])) {
            return i;
        }
    }
    return -1;
}

2. 

int findUnique(int[] arr) {
    int unique = arr[0];
    for(int i=1; i<arr.length; i++) {
        unique ^= arr[i];
    }
    return unique;
}

3.

int findMissing(int[] arr) {
    int n = arr.length;
    boolean[] present = new boolean[n+1];
    for(int i=0; i<n; i++) {
        if(arr[i] > 0 && arr[i] <= n) {
            present[arr[i]] = true;
        }
    }
    for(int i=1; i<=n; i++) {
        if(!present[i]) {
            return i;
        }
    }
    return n+1;
}

4.

int findDuplicate(int[] arr) {
    for(int i=0; i<arr.length; i++) {
        if(arr[Math.abs(arr[i])] >= 0) {
            arr[Math.abs(arr[i])] = -arr[Math.abs(arr[i])];
        } else {
            return Math.abs(arr[i]);
        }
    }
    return -1;
}

5.

String longestString(String[] arr) {
    int maxLength = 0;
    String longest = "";
    for(int i=0; i<arr.length; i++) {
        if(arr[i].length() > maxLength) {
            maxLength = arr[i].length();
            longest = arr[i];
        }
    }
    return longest;
}   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Linear Search的網路資料
{{< /ask_chatgpt >}}



1. GeeksforGeeks: Linear Search
https://www.geeksforgeeks.org/linear-search/

GeeksforGeeks is a popular programming website known for providing tutorials and articles related to computer science and programming. This article provides an introduction to Linear Search, explaining the concept and the algorithm with code examples in Java, C++, and Python.

2. Programiz: Linear Search Algorithm
https://www.programiz.com/dsa/linear-search

Programiz is another popular programming website that provides tutorials and articles on various programming concepts. This article explains Linear Search algorithm in detail, explaining its complexity, implementation, and time complexity through code examples in C++ and Python.

3. Studytonight: Linear Search Algorithm
https://www.studytonight.com/data-structures/linear-search

Studytonight is a website that provides free tutorials on various programming concepts, including OOPs, Data Structure, and Algorithms, and more. This article explains Linear Search Algorithm in detail, giving an example of how it works and how to implement it using C programming language.

4. Tutorials Point: Linear Search Algorithm
https://www.tutorialspoint.com/data_structures_algorithms/linear_search_algorithm.htm

Tutorials Point is another online platform that provides educational content related to computer science and programming. This article explains the concept of Linear Search Algorithm, how it works, and how to implement it in C programming language through code examples.

5. JavaTpoint: Linear Search
https://www.javatpoint.com/linear-search

JavaTpoint is a Java-based programming website that provides educational content on Java programming language and other related concepts. This article explains the concept of Linear Search algorithm and shows how to implement it using Java programming language. The article also explains its time and space complexity and how to optimize the code for better speed.   

