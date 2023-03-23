+++
title = "Jasmine"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Jasmine的问题
{{< /ask_chatgpt >}}



1. Jasmine喜歡做什麼樣的活動？
2. Jasmine的家人有幾個兄弟姊妹？
3. Jasmine最喜歡的食物是什麼？
4. Jasmine最崇拜的偶像是誰？
5. Jasmine平常喜歡用哪種顏色的筆寫字？   

## 说明知识
{{< ask_chatgpt >}}
说明Jasmine并举例
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

## 汇总重点
{{< ask_chatgpt >}}
条列Jasmine的重点
{{< /ask_chatgpt >}}



抱歉，我需要更多的上下文才能回答这个问题，Jasmine是什么？ -- AI助手   

## 知识测验
{{< ask_chatgpt >}}
给我5题Jasmine的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 给定一个整数 n，输出所有小于 n 的正整数中，能被 3 或 5 整除的数字的总和。
   范例输入：10
   范例输出：8 (3 + 5)

2. 给定一个整数阵列 nums 和一个目标整数 target，找出 nums 中两个数字的和等于 target，并返回它们的索引。
   假设每个输入仅有一个解，且同一个元素不能使用两次。
   范例输入：nums = [2, 7, 11, 15], target = 9
   范例输出：[0, 1]

3. 实现一个费伯纳契数列的函数。费伯纳契数列是一系列的整数，其中每个数字都是前两个数字的和。
   范例输入：10
   范例输出：[0, 1, 1, 2, 3, 5, 8, 13, 21, 34]

4. 实现一个函数，将一个给定的字符串转换为小写。
   范例输入：Jasmine
   范例输出：jasmine

5. 给定一个整数 x，实现一个函数，计算 x 的平方根。如果 x 是负数，则返回 null。
   范例输入：16
   范例输出：4.0

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

