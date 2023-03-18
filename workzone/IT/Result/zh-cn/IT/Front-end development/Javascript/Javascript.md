+++
title = "JavaScript"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Javascript的问题
{{< /ask_chatgpt >}}



1. 如何將一個字串反轉？

```
function reverseString(str) {
   return str.split("").reverse().join("");
}

console.log(reverseString("Hello")); // "olleH"
```

2. 如何取得一個Array的最大值？

```
function getMax(arr) {
  return Math.max.apply(null, arr);
}

console.log(getMax([3, 7, 2, 5, 6])); // 7
```

3. 如何檢查一個變數是否為數字？

```
function isNumeric(num) {
  return !isNaN(parseFloat(num)) && isFinite(num);
}

console.log(isNumeric('123')); // true
console.log(isNumeric('abc')); // false
```

4. 如何在HTML中插入一個圖片？

```
<img src="path/to/image.jpg" alt="Image Description">
```

5. 如何在Javascript中檢查一個變數是否定義？

```
if (typeof myVar === 'undefined') {
  console.log("myVar is not defined");
} else {
  console.log("myVar is defined");
}
```   

## 说明知识
{{< ask_chatgpt >}}
说明Javascript并举例
{{< /ask_chatgpt >}}



JavaScript是一種前端腳本語言，是網頁互動的核心之一。它可以用來編寫動態效果、交互式功能和驗證用戶輸入的表單。JavaScript可以與HTML和CSS融合在一起，並且可以在網頁上運行，而不需要額外的插件或程式。

以下是JavaScript的一些示例：

1. 更改網頁元素的內容：可以使用JavaScript來更改網頁上的內容，例如更改標題、段落或按鈕的文本內容。

2. 彈出式視窗：可以使用JavaScript來顯示彈出視窗，例如警示框、確認框或提示框，來讓用戶進行選擇或輸入。

3. 網頁動畫效果：可以使用JavaScript來創建動畫效果，例如簡單的滑動、淡入淡出或完整的視差滾動效果。

4. 表單驗證：可以使用JavaScript來驗證用戶輸入的表單，例如確認密碼是否匹配、檢查電子郵件地址格式或限制用戶輸入的內容。

5. 輪播效果：可以使用JavaScript來創建輪播效果，例如自動輪播圖像、手動輪播頁面或無限輪播效果。

6. Ajax請求：可以使用JavaScript來發送Ajax請求，與網頁上的伺服器進行交互，進行數據檢索或提交表單的時候。

總之，JavaScript是一個功能豐富且簡單易用的網頁腳本語言，可以實現各種互動效果，可以使網站更加生動和吸引人。   

## 汇总重点
{{< ask_chatgpt >}}
条列Javascript的重点
{{< /ask_chatgpt >}}



1. JavaScript 是一种脚本语言，可以在网页上运行，与 HTML 和 CSS 配合，实现动态交互的效果。

2. JavaScript 的数据类型包括数字、字符串、布尔值、数组、对象等，可以通过变量、常量、运算符等进行操作。

3. JavaScript 支持基础的逻辑控制语句，如 if…else、while、for 等，还可以使用函数和事件进行代码的封装和重用。

4. JavaScript 可以与 HTML 元素进行交互，可以通过 DOM 操作对网页元素进行增、删、改、查的操作。

5. JavaScript 的异步编程模型可以通过 Promise、async/await、setTimeout 等方式实现。

6. JavaScript 还可以使用框架和库来快速实现特定的功能，如 React、Vue、jQuery 等。

7. JavaScript 的错误处理可以使用 try…catch 语句来进行捕获和处理。

8. JavaScript 可以与服务端语言进行交互，如通过 AJAX 技术来获取和发送数据。

9. JavaScript 有很多常用的内置对象和方法，如 String、Array、Math、Date 等，可以大大简化编程任务。

10. JavaScript 也有许多新的特性和语法，如箭头函数、解构赋值、模板字符串等，可以提升代码的可读性和开发效率。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Javascript的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 寫一個函式，判斷一個數字是否為質數。

```
function isPrime(num) {
  if (num <= 1) return false;
  for (let i = 2; i <= Math.sqrt(num); i++) {
    if (num % i === 0) return false;
  }
  return true;
}

console.log(isPrime(11)); // true
console.log(isPrime(12)); // false
```


2. 寫一個函式，將一個整數轉換成羅馬數字。

```
function intToRoman(num) {
  const mapping = {
    M: 1000,
    CM: 900,
    D: 500,
    CD: 400,
    C: 100,
    XC: 90,
    L: 50,
    XL: 40,
    X: 10,
    IX: 9,
    V: 5,
    IV: 4,
    I: 1
  };

  let result = '';
  for (let key in mapping) {
    while (num >= mapping[key]) {
      result += key;
      num -= mapping[key];
    }
  }

  return result;
}

console.log(intToRoman(1994)); // "MCMXCIV"
```


3. 寫一個函式，將一個字符串中的單詞反轉。

```
function reverseWords(str) {
  return str
    .split(' ')
    .map(word => word.split('').reverse().join(''))
    .join(' ');
}

console.log(reverseWords("Let's code in JavaScript!")); // "s'teL edoc ni tpircSavaJ!"
```


4. 寫一個函式，移除一個數組中的重複元素。

```
function removeDuplicates(nums) {
  let i = 0;
  for (let j = 1; j < nums.length; j++) {
    if (nums[i] !== nums[j]) {
      i++;
      nums[i] = nums[j];
    }
  }
  
  return nums.slice(0, i + 1);
}

console.log(removeDuplicates([1,1,2,2,3,4,4,5,5,5])); // [1, 2, 3, 4, 5]
```


5. 寫一個函式，取得一個數組中第二大的數字。

```
function secondLargest(nums) {
  let max = -Infinity;
  let secondMax = -Infinity;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] > max) {
      secondMax = max;
      max = nums[i];
    } else if (nums[i] > secondMax && nums[i] !== max) {
      secondMax = nums[i];
    }
  }
  
  return secondMax;
}

console.log(secondLargest([3, 1, 5, 9, 2, 7])); // 7
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Javascript的网络数据
{{< /ask_chatgpt >}}



1. "JavaScript.com" - https://www.javascript.com/
這是由Stack Overflow推出的官方JavaScript學習資源網站，提供新手介紹、實用語法以及最新的JavaScript開發趨勢。

2. "MDN Web Docs - JavaScript" - https://developer.mozilla.org/en-US/docs/Web/JavaScript
這是Mozilla基金會提供的JavaScript開發技術文件，包括詳細的語法介紹、範例代碼、應用場景以及最新的API與規範。

3. "JavaScript.info" - https://javascript.info/
這是一個由俄羅斯Web開發者Ilya Kantor開發的JavaScript學習資源，提供Step by Step的學習路線、詳細的JavaScript內容解說以及免費的電子書下載。

4. "Eloquent JavaScript" - https://eloquentjavascript.net/
這是一本由Marijn Haverbeke撰寫的JavaScript學習書，內容包含JavaScript的基礎概念、流程控制、函式、物件、陣列等等，同時也提供互動式教學實驗室供學習者練習。

5. "JavaScript Weekly" - https://javascriptweekly.com/
這是一個免費的每週JavaScript技術資訊彙編，提供最新的前端開發技術、工具、框架與實用資源，也包括最新的JavaScript網路資訊與事件。   

