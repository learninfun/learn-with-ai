## 習題預習
<<ASK_CHATGPT>>  
給我5題Javascript的問題  
  
<<ANS_CHATGPT>>  


1. 如何將一個字符串轉換為數字？
2. 什麼是閉包？它有什麼作用？
3. 如何訪問對象中的屬性？
4. 如何在Javascript中建立一個類？
5. 什麼是事件？如何綁定事件？  

## 說明知識
<<ASK_CHATGPT>>  
說明Javascript並舉例  
  
<<ANS_CHATGPT>>  


Javascript 是一種網頁腳本語言，它可以在網頁中動態地加載、改變和操作內容，從而實現交互效果和網頁動態效果。

舉例如下：

1. 點擊按鈕觸發事件：當用戶在網頁中點擊按鈕時，Javascript 可以檢測到事件，並觸發相應的功能，例如顯示提示消息、向後端發送請求等。

2. 表單驗證：在網頁中，用戶提交表單時可以使用 Javascript 來檢查表單是否已經填寫完整，如果不完整就提示用戶還有哪些欄位需要填寫。

3. 輪播圖效果：網頁中的輪播圖效果也是通過 Javascript 達到的，通過控制圖片切換和顯示，使得網頁具有更好的視覺效果。

4. 動態顯示數據：Javascript 還可以與後端 API 進行交互，從而實現動態顯示數據。例如在一個電商網站上，當用戶登錄後，Javascript 可以通過向後端發送請求，獲取該用戶的購物車列表，並在網頁中實時顯示出來。  

## 彙總重點
<<ASK_CHATGPT>>  
條列Javascript的重點  
  
<<ANS_CHATGPT>>  


1. JavaScript是一种脚本语言，可以用于客户端和服务器端编程。
2. JavaScript可以与HTML和CSS一起使用来添加交互性和动态效果。
3. JavaScript支持面向对象编程和函数式编程。
4. 变量的作用域在JavaScript中非常重要，可以通过var，let和const来声明变量。
5. JavaScript具有丰富的内置对象，包括数组，字符串，日期等，还可以创建自己的对象。
6. JavaScript支持各种条件语句和循环语句，例如if语句，switch语句和while语句。
7. JavaScript可以执行异步操作，例如使用回调函数，Promise和async/await。
8. JavaScript可以通过使用DOM和jQuery来操作HTML文档和页面元素。
9. JavaScript可以通过使用Ajax和fetch来与服务器进行通信和获取数据。
10. JavaScript还可以通过Node.js和其他后端JavaScript框架来构建Web应用程序和API。  

## 知識測驗
<<ASK_CHATGPT>>  
給我5題Javascript的中等難度問題，並在後面列出答案  
  
<<ANS_CHATGPT>>  


1. 實作一個函式，可以找到一個字串裡出現最多次的字母，如果有多個字母出現次數相同，則回傳其中一個。

答案：

```javascript
function findMostFreqChar(str) {
  let maxCount = 0;
  let mostFreqChar = '';
  let charCount = {};
  for(let i = 0; i < str.length; i++) {
    let char = str[i];
    if(charCount[char]) {
      charCount[char]++;
    } else {
      charCount[char] = 1;
    }
    if(charCount[char] > maxCount) {
      maxCount = charCount[char];
      mostFreqChar = char;
    }
  }
  return mostFreqChar;
}
```

2. 實作一個函式，可以將一個數字的所有位數相加，如果結果還是不只一個位數，則繼續把各位數相加，重複這個步驟直到最後只有一個位數。

答案：

```javascript
function sumDigits(num) {
  while(num >= 10) {
    let sum = 0;
    while(num > 0) {
      sum += num % 10;
      num = Math.floor(num / 10);
    }
    num = sum;
  }
  return num;
}
```

3. 實作一個函式，可以將一個字串裡所有的英文字母往後移動幾個位置，例如將字母表往後移3個位置後，A變成D，B變成E，X變成A。（注意Z往後移動後變成A）

答案：

```javascript
function moveLetters(str, n) {
  let result = '';
  for(let i = 0; i < str.length; i++) {
    let charCode = str.charCodeAt(i);
    let shiftedCharCode = charCode + n;
    if(charCode >= 65 && charCode <= 90) { // A-Z
      shiftedCharCode = (shiftedCharCode % 90) + (shiftedCharCode > 90 ? 64 : 0);
    } else if(charCode >= 97 && charCode <= 122) { // a-z
      shiftedCharCode = (shiftedCharCode % 122) + (shiftedCharCode > 122 ? 96 : 0);
    }
    result += String.fromCharCode(shiftedCharCode);
  }
  return result;
}
```

4. 實作一個函式，可以找到一個陣列裡的最長遞增子序列的長度，例如[1, 4, 2, 5, 3]的最長遞增子序列是[1, 2, 3, 5]，長度為4。

答案：

```javascript
function findLengthOfLongestIncreasingSubsequence(nums) {
  let dp = new Array(nums.length).fill(1);
  for(let i = 1; i < nums.length; i++) {
    for(let j = 0; j < i; j++) {
      if(nums[j] < nums[i]) {
        dp[i] = Math.max(dp[i], dp[j] + 1);
      }
    }
  }
  return Math.max(...dp);
}
```

5. 實作一個函式，可以找到一個字串裡最長的迴文子串，例如在字串"babad"中，最長的迴文子串是"bab"或"aba"。

答案：

```javascript
function findLongestPalindrome(s) {
  let start = 0;
  let maxLength = 1;
  for(let i = 0; i < s.length; i++) {
    for(let j = i + 1; j < s.length; j++) {
      let len = j - i + 1;
      if(len > maxLength && isPalindrome(s, i, j)) {
        maxLength = len;
        start = i;
      }
    }
  }
  return s.substring(start, start + maxLength);
}

function isPalindrome(s, left, right) {
  while(left < right) {
    if(s[left] !== s[right]) {
      return false;
    }
    left++;
    right--;
  }
  return true;
}
```  

## 網路資料
<<ASK_CHATGPT>>  
給我5篇Javascript的網路資料  
  
<<ANS_CHATGPT>>  


1. JavaScript | MDN
https://developer.mozilla.org/en-US/docs/Web/JavaScript

This is the official resource for JavaScript documentation, maintained by the Mozilla Foundation. You can find everything from beginner information to advanced topics, and the website also includes examples and code snippets.

2. JavaScript.com
https://www.javascript.com/

This website is a great resource for developers of all levels. It includes tutorials, articles, and code snippets that will help you learn the language and become a proficient JavaScript developer.

3. JavaScript | W3Schools
https://www.w3schools.com/js/

W3Schools is a popular web development resource, and their JavaScript tutorials and examples are very helpful. They cover everything from basic syntax to advanced topics like object-oriented programming and AJAX.

4. Eloquent JavaScript
https://eloquentjavascript.net/

This free online book is a comprehensive guide to JavaScript. It's geared toward intermediate to advanced developers, but beginners can also benefit from the explanations and examples.

5. Codecademy JavaScript Course
https://www.codecademy.com/learn/introduction-to-javascript

Codecademy is a great place to begin learning JavaScript. Their JavaScript course includes interactive lessons and quizzes, and they also offer a paid subscription for more in-depth learning.  

