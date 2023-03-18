+++
title = "Jasmine"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Jasmine的問題
{{< /ask_chatgpt >}}



1. Jasmine喜歡做什麼樣的活動？
2. Jasmine的家人有幾個兄弟姊妹？
3. Jasmine最喜歡的食物是什麼？
4. Jasmine最崇拜的偶像是誰？
5. Jasmine平常喜歡用哪種顏色的筆寫字？   

## 說明知識
{{< ask_chatgpt >}}
說明Jasmine並舉例
{{< /ask_chatgpt >}}



Jasmine是一個用於JavaScript測試的行為驅動開發（BDD）框架。它的設計目的是使JavaScript測試變得簡單明了，並且能夠直觀地表達代碼的預期行為。

在Jasmine中，測試以“描述”（describe）和“斷言”（expect）的形式進行。描述是對代碼功能的解釋性的文字描述，可以嵌套描述和斷言，進一步剖析代碼的不同層面。 斷言是一個在測試中對代碼的預期結果的陳述。Jasmine提供了多種斷言方法，比如toBe，toEqual，toContain等等。

舉個例子，假設我們有一個JavaScript函數，例如：

```
function addNumbers(a, b) {
  return a + b;
}
```

要測試這個函數的行為，我們可以使用Jasmine。我們可以寫一個描述來解釋此函數的功能，然後寫一個斷言來驗證函數是否正確，如下所示：

```
describe('addNumbers function', function() {
  
  it('adds two numbers and returns the correct result', function() {
    var result = addNumbers(2, 3);
    expect(result).toEqual(5);
  });
  
});
```

在這個示例中，我們編寫了一個描述來說明這個測試是關於“addNumbers function”函數的。在描述的內部，我們定義了一個it（也稱為“測試用例”），在此測試用例中，我們使用了addNumbers函數來加入兩個數字，然後使用Jasmine的expect斷言來驗證這個函數的返回值是否為5。如果測試通過，Jasmine將會顯示一個綠色符號表示成功，否則將會顯示一個紅色符號表示失敗。   

## 彙總重點
{{< ask_chatgpt >}}
條列Jasmine的重點
{{< /ask_chatgpt >}}



抱歉，我需要更多的上下文才能回答這個問題，Jasmine是什麼？ -- AI助手   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Jasmine的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 給定一個整數 n，輸出所有小於 n 的正整數中，能被 3 或 5 整除的數字的總和。
   範例輸入：10
   範例輸出：8 (3 + 5)

2. 給定一個整數陣列 nums 和一個目標整數 target，找出 nums 中兩個數字的和等於 target，並返回它們的索引。
   假設每個輸入僅有一個解，且同一個元素不能使用兩次。
   範例輸入：nums = [2, 7, 11, 15], target = 9
   範例輸出：[0, 1]

3. 實現一個費伯納契數列的函數。費伯納契數列是一系列的整數，其中每個數字都是前兩個數字的和。
   範例輸入：10
   範例輸出：[0, 1, 1, 2, 3, 5, 8, 13, 21, 34]

4. 實現一個函數，將一個給定的字符串轉換為小寫。
   範例輸入：Jasmine
   範例輸出：jasmine

5. 給定一個整數 x，實現一個函數，計算 x 的平方根。如果 x 是負數，則返回 null。
   範例輸入：16
   範例輸出：4.0

答案：
1. 
function find_sum(n){
  let sum = 0;
  for (let i = 1; i < n; i++){
    if (i % 3 === 0 || i % 5 === 0){
      sum += i;
    }
  }
  return sum;
}

2. 
function twoSum(nums, target) {
  let map = new Map();
  for (let i = 0; i < nums.length; i++){
    let complement = target - nums[i];
    if (map.has(complement)){
      return [map.get(complement), i];
    }
    map.set(nums[i], i);
  }
}

3. 
function fibonacci(n) {
  let res = [];
  let a = 0;
  let b = 1;
  for (let i = 0; i < n; i++){
    res.push(a);
    let temp = a + b;
    a = b;
    b = temp;
  }
  return res;
}

4. 
function toLowerCase(str) {
  return str.toLowerCase();
}

5. 
function sqrt(x) {
  if (x < 0) return null;
  let left = 0;
  let right = x;
  while (left <= right) {
    let mid = Math.floor((left + right) / 2);
    if (mid * mid === x) return mid;
    if (mid * mid < x) left = mid + 1;
    if (mid * mid > x) right = mid - 1;
  }
  return right;
}   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Jasmine的網路資料
{{< /ask_chatgpt >}}



1. Jasmine - W3Docs
https://www.w3docs.com/tools/jasmine/jasmine-tutorial.html

Jasmine是一個JavaScript測試框架，它讓您可以創建自動化單元測試。這篇文章提供了Jasmine的入門指南，包括如何安裝和使用Jasmine，如何創建測試套件以及如何寫測試規格。

2. Jasmine JavaScript Testing - Tutorialspoint
https://www.tutorialspoint.com/jasminejs/index.htm

這是一個徹底的Jasmine教程，涵蓋從基礎到高級的主題。教程將教您如何創建Jasmine測試套件，利用匹配器寫測試規格，進行基礎和高級斷言，並在測試期間使用假數據。

3. Jasmine - GitHub
https://github.com/jasmine/jasmine

Jasmine是一個開源的JavaScript測試框架，由Pivotal Labs開發。在這裡，您可以找到Jasmine的源代碼，問題跟踪，社區貢獻方式等等。

4. Jasmine Documentation - Jasmine
https://jasmine.github.io/index.html

Jasmine的官方文檔提供了深入的Jasmine架構和功能的資訊。在這裡，您可以找到Jasmine的API文檔，包括匹配器，斷言，突變測試等等。

5. Unit Testing in JavaScript With Jasmine - SitePoint
https://www.sitepoint.com/unit-testing-javascript-jasmine/

這篇文章講解了JavaScript測試的基礎和Jasmine的使用。它探討了為什麼需要進行測試，以及如何使用Jasmine來創建和執行自動化測試套件。還提供了幾個Jasmine的示例，以及如何在開發過程中使用測試來增加代碼覆蓋率的幾個技巧。   

