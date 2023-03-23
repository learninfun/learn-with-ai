+++
title = "Jest"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Jest的中文問題
{{< /ask_chatgpt >}}



1. Jest是什麼，可以用來做什麼？
2. Jest如何設置和使用？
3. Jest有哪些測試方法和斷言可用？及其使用方式。
4. Jest能否與其他測試工具集成？例如Babel或TypeScript。
5. Jest如何處理異步測試？有哪些方法可供選擇？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Jest並舉例
{{< /ask_chatgpt >}}



Jest是一款由Facebook所開發出的 JavaScript測試框架，專門用於測試React應用程式、React Native應用程式以及任何其他的JavaScript專案。Jest具有易上手且功能豐富的特點，它不需要額外的配置文件即可開始使用，並且支持快照測試、Mock、斷言、覆蓋率等測試類型。

以下是一個簡單的Jest測試案例：

```js
// sum.js
function sum(a, b) {
  return a + b;
}

module.exports = sum;
```

```js
// sum.test.js
const sum = require('./sum');

test('adds 1 + 2 to equal 3', () => {
  expect(sum(1, 2)).toBe(3);
});
```

在這個案例中，我們首先定義了一個簡單的sum函數，接著建立了一個測試檔案sum.test.js，在這個測試檔案中，我們使用了Jest提供的test函數來描述這個測試案例。在這個測試案例中，我們期望sum(1, 2)的結果會等於3，這個期望值透過Jest提供的斷言函數expect和匹配器toBe來實現。

執行Jest測試需要在終端機中輸入以下指令：

```
npm test
```

執行測試之後，Jest將會輸出以下結果：

```
PASS  ./sum.test.js
 ✓ adds 1 + 2 to equal 3 (5ms)
```

Jest指示測試通過，並且提供了測試的描述以及耗時。這個簡單的測試案例展示了Jest在JavaScript專案中的使用方式。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Jest的重點
{{< /ask_chatgpt >}}



1. Jest是一个流行的JavaScript测试框架。
2. Jest支持自动化测试、持续集成和代码覆盖率。
3. Jest具有简洁且易于使用的API，使得开发人员可以轻松地创建和管理测试用例。
4. Jest提供了模块化和异步测试功能，方便测试异步代码。
5. Jest还支持快照测试，用于比较不同版本的代码输出的差异。
6. Jest内置了Mock和Spy功能，可以模拟外部依赖项和检测函数的调用情况。
7. Jest具有快速和可扩展的特点，可以满足各种测试环境的需求。
8. Jest是基于JavaScript的，可以与React、Vue、Angular等流行的JavaScript框架和库使用。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Jest的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1) 以下是一個計算器的 add 函數，請寫測試用例來測試該函數是否正確：

```js
function add(a, b) {
  return a + b;
}
```

答案：

```js
test("add function works correctly", () => {
  expect(add(2, 2)).toBe(4);
  expect(add(0, 4)).toBe(4);
  expect(add(-2, 3)).toBe(1);
  expect(add(1.2, 2.3)).toBeCloseTo(3.5);
});
```

2) 以下是一個輸入金額，計算總價的函數，請寫測試用例來測試該函數是否正確：

```js
function calculateTotalPrice(price, quantity) {
  const total = price * quantity;
  return `Total price is ${total}`;
}
```

答案：

```js
test("calculateTotalPrice function works correctly", () => {
  expect(calculateTotalPrice(10, 3)).toBe("Total price is 30");
  expect(calculateTotalPrice(5, 0)).toBe("Total price is 0");
  expect(calculateTotalPrice(2.5, 4)).toBe("Total price is 10");
  expect(calculateTotalPrice(8.99, 1)).toBe("Total price is 8.99");
});
```

3) 以下是一個判斷年份是否為閏年的函數，請寫測試用例來測試該函數是否正確：

```js
function isLeapYear(year) {
  if (year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0)) {
    return true;
  } else {
    return false;
  }
}
```

答案：

```js
test("isLeapYear function works correctly", () => {
  expect(isLeapYear(2000)).toBe(true);
  expect(isLeapYear(1900)).toBe(false);
  expect(isLeapYear(2020)).toBe(true);
  expect(isLeapYear(2022)).toBe(false);
});
```

4) 以下是一個輸入條件，判斷是否符合註冊要求的函數，請寫測試用例來測試該函數是否正確：

```js
function isRegisterValid(username, password, confirmPassword) {
  if (username.length >= 3 && password.length >= 6 && password === confirmPassword) {
    return true;
  } else {
    return false;
  }
}
```

答案：

```js
test("isRegisterValid function works correctly", () => {
  expect(isRegisterValid("john", "123456", "123456")).toBe(true);
  expect(isRegisterValid("joe123", "password", "password")).toBe(false);
  expect(isRegisterValid("user", "123456", "654321")).toBe(false);
  expect(isRegisterValid("admin", "password", "password")).toBe(false);
});
```

5) 以下是一個輸入年月日，計算星期幾的函數，請寫測試用例來測試該函數是否正確：

```js
function getDayOfWeek(year, month, day) {
  const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
  const date = new Date(`${month}/${day}/${year}`);
  const dayOfWeek = daysOfWeek[date.getDay()];
  return dayOfWeek;
}
```

答案：

```js
test("getDayOfWeek function works correctly", () => {
  expect(getDayOfWeek(2022, 11, 16)).toBe("Wednesday");
  expect(getDayOfWeek(1970, 1, 1)).toBe("Thursday");
  expect(getDayOfWeek(2021, 7, 4)).toBe("Sunday");
  expect(getDayOfWeek(2030, 12, 25)).toBe("Wednesday");
});
```   

