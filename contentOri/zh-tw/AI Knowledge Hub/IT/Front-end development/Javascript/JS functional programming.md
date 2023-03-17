## 習題預習
{{< ask_chatgpt >}}
給我5題JS functional programming的問題
{{< /ask_chatgpt >}}



1. 如何實現一個 map 函數，用於對數組中的每個元素進行指定操作，並返回新的數組？

2. 如何使用 reduce 函數實現數組中所有數字的加總？

3. 如何使用 filter 函數實現選擇數組中所有大於10的數字？

4. 如何使用 compose 函數把多個函數串接起來，實現複合函數？

5. 如何使用柯裡化（currying）實現一個通用的 add 函數，能夠進行任意數字的相加？   

## 說明知識
{{< ask_chatgpt >}}
說明JS functional programming並舉例
{{< /ask_chatgpt >}}



JS functional programming 是一種程式設計的風格或範式，其主要特點是使用純函數（pure functions）的方式來進行開發。

純函數是指對於相同的輸入，函數總是返回相同的輸出，且不會對全域變數產生影響。換句話說，純函數只受其輸入參數所影響，並不依賴於外部環境的任何因素。

以下是一個純函數的例子：

```
function add(a, b) {
  return a + b;
}
```

這個函數只是取兩個數相加後返回結果，沒有任何副作用（side effect）。

使用 functional programming 的好處是可以讓程式碼更易於理解、測試和重用，因為每個函數都是獨立的、沒有副作用的。

除了純函數之外，JS functional programming 還涉及到很多其他的概念，如高階函數、閉包等等。這些概念可以協助開發人員更好地組織和抽像代碼。

以下是一個使用高階函數的例子：

```
function map(arr, fn) {
  const result = [];
  for(let i = 0; i < arr.length; i++) {
    result.push(fn(arr[i]));
  }
  return result;
}

const numbers = [1, 2, 3, 4, 5];
const squares = map(numbers, function(num) {
  return num * num;
});
console.log(squares); // [1, 4, 9, 16, 25]
```

這裡定義了一個 `map` 函數，它接受一個數組和一個函數作為輸入，並返回一個新的數組，其中每個元素都是原始數組元素應用函數之後的值。這個函數就是一個高階函數，它可以接受另一個函數作為輸入。

總之，JS functional programming 是一個強大的程式設計風格，它可以幫助開發人員改善代碼的品質和效率。通過使用純函數、高階函數等概念，開發人員可以更好地組織代碼、寫出更易於理解和測試的代碼。   

## 彙總重點
{{< ask_chatgpt >}}
條列JS functional programming的重點
{{< /ask_chatgpt >}}



1. 價值原則：函式是值，而不是命令式的指令。

2. 純函式：函式對相同的引數輸入返回相同的結果，並且沒有副作用。

3. 不可變性：數據在創建後保持不變。

4. 好的過程：函式應該是被設計成盡可能簡單明瞭、易於閱讀和測試。

5. 組合性：函式應該能夠被組合在一起以創建更複雜的函式。

6. 高階函式：函式可以接受函式作為引數，並且可以返回一個函式。

7. 遞歸：使用遞歸解決問題，可以使代碼更簡潔和易於閱讀。

8. 類型推斷：類型應該基於值和行為的推斷來進行，而不是顯式聲明。

9. 惰性求值：僅計算必須使用的值，可以使代碼更有效率。

10. 柯裡化與部分應用：柯裡化和部分應用技術可以簡化函式的實現，使其更容易理解和使用。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題JS functional programming的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實作一個函式，接收一個數字陣列並回傳最大值。

```
const getMax = arr => {
  return Math.max(...arr);
}

console.log(getMax([1, 2, 3, 4, 5]));  // 5
console.log(getMax([-1, -3, -5, -10]));  // -1
```

2. 實作一個函式，接收一個數字陣列並回傳所有元素的總和。

```
const getSum = arr => {
  return arr.reduce((acc, curr) => acc + curr, 0);
}

console.log(getSum([1, 2, 3, 4, 5]));  // 15
console.log(getSum([-1, -3, -5, -10]));  // -19
```

3. 實作一個函式，接收一個數字陣列並回傳去重後的陣列。

```
const getUnique = arr => {
  return [...new Set(arr)];
}

console.log(getUnique([1, 2, 3, 4, 5]));  // [1, 2, 3, 4, 5]
console.log(getUnique([1, 2, 2, 3, 3, 3, 4, 5, 5]));  // [1, 2, 3, 4, 5]
```

4. 實作一個函式，接收一個字串，回傳其所有單字都首字母大寫的字串。

```
const formatString = str => {
  return str.split(' ')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
            .join(' ');
}

console.log(formatString('hello world'));  // 'Hello World'
console.log(formatString('JAVASCRIPT IS AWESOME'));  // 'Javascript Is Awesome'
```

5. 實作一個函式，接收一個數字陣列並回傳陣列中大於10的元素所組成的陣列。

```
const getGreaterThanTen = arr => {
  return arr.filter(num => num > 10);
}

console.log(getGreaterThanTen([1, 5, 10, 15, 20]));  // [15, 20]
console.log(getGreaterThanTen([2, 4, 6, 8]));  // []
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇JS functional programming的網路資料
{{< /ask_chatgpt >}}



1. "Functional Programming in JavaScript" by Mozilla Developer Network: This comprehensive guide covers the basics of functional programming, how it works in JavaScript, and what benefits it offers to developers. It includes examples of functional programming concepts and shows how they can be implemented in JavaScript.

2. "7 Reasons to Use Functional Programming in JavaScript" by Akshay Shanbhag: This article outlines seven key reasons why functional programming is a valuable approach to writing JavaScript code. It discusses benefits such as code readability, easier testing, and improved handling of side effects.

3. "Functional Programming for Beginners: Understanding JavaScript's Reduce Method" by SitePoint: This tutorial explains how to use the reduce method in JavaScript, a key tool for functional programming. It covers the basics of the reduce method, how it works, and how it can be used to solve common programming problems.

4. "Functional Programming in JavaScript: What, Why, and How" by Kevin Kononenko: This video presentation provides a concise overview of functional programming in JavaScript, including its history, principles, and benefits. It also includes examples of functional programming concepts in action.

5. "Writing Clean JavaScript with Functional Programming" by Caroline Artz: This article explores the role of functional programming in writing clean, maintainable JavaScript code. It includes examples of how functional programming concepts can be used to improve code quality and reduce bugs.   

