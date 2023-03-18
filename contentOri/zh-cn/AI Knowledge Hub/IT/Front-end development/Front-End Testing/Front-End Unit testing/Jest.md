+++
title = "Jest"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Jest的问题
{{< /ask_chatgpt >}}



1. Jest是什么，可以用来做什么？
2. Jest如何设置和使用？
3. Jest有哪些测试方法和断言可用？及其使用方式。
4. Jest能否与其他测试工具集成？例如Babel或TypeScript。
5. Jest如何处理异步测试？有哪些方法可供选择？   

## 说明知识
{{< ask_chatgpt >}}
说明Jest并举例
{{< /ask_chatgpt >}}



Jest是一款由Facebook所开发出的 JavaScript测试框架，专门用于测试React应用程式、React Native应用程式以及任何其他的JavaScript专案。Jest具有易上手且功能丰富的特点，它不需要额外的配置文件即可开始使用，并且支持快照测试、Mock、断言、覆盖率等测试类型。

以下是一个简单的Jest测试案例：

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

在这个案例中，我们首先定义了一个简单的sum函数，接着建立了一个测试档案sum.test.js，在这个测试档案中，我们使用了Jest提供的test函数来描述这个测试案例。在这个测试案例中，我们期望sum(1, 2)的结果会等于3，这个期望值透过Jest提供的断言函数expect和匹配器toBe来实现。

执行Jest测试需要在终端机中输入以下指令：

```
npm test
```

执行测试之后，Jest将会输出以下结果：

```
PASS  ./sum.test.js
 ✓ adds 1 + 2 to equal 3 (5ms)
```

Jest指示测试通过，并且提供了测试的描述以及耗时。这个简单的测试案例展示了Jest在JavaScript专案中的使用方式。   

## 汇总重点
{{< ask_chatgpt >}}
条列Jest的重点
{{< /ask_chatgpt >}}



1. Jest是一个流行的JavaScript测试框架。
2. Jest支持自动化测试、持续集成和代码覆盖率。
3. Jest具有简洁且易于使用的API，使得开发人员可以轻松地创建和管理测试用例。
4. Jest提供了模块化和异步测试功能，方便测试异步代码。
5. Jest还支持快照测试，用于比较不同版本的代码输出的差异。
6. Jest内置了Mock和Spy功能，可以模拟外部依赖项和检测函数的调用情况。
7. Jest具有快速和可扩展的特点，可以满足各种测试环境的需求。
8. Jest是基于JavaScript的，可以与React、Vue、Angular等流行的JavaScript框架和库使用。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Jest的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1) 以下是一个计算器的 add 函数，请写测试用例来测试该函数是否正确：

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

2) 以下是一个输入金额，计算总价的函数，请写测试用例来测试该函数是否正确：

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

3) 以下是一个判断年份是否为闰年的函数，请写测试用例来测试该函数是否正确：

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

4) 以下是一个输入条件，判断是否符合注册要求的函数，请写测试用例来测试该函数是否正确：

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

5) 以下是一个输入年月日，计算星期几的函数，请写测试用例来测试该函数是否正确：

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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Jest的网络数据
{{< /ask_chatgpt >}}



1. 官方网站：https://jestjs.io/

Jest是由Facebook开发的JavaScript测试框架。它支持自动化测试和手动测试，可用于单元测试、集成测试和端对端测试等多种测试类型。官方网站提供了详细的文档和教程，方便开发者快速上手使用。

2. Medium文章：https://medium.com/welldone-software/an-intro-to-jest-what-it-is-and-how-to-use-it-27b1f3dcaf5c

该Medium文章是一个对Jest的简介，主要介绍了Jest的目的、功能以及如何使用它。它还包含了示例代码和实际的测试用例，帮助读者更好地理解Jest的工作方式。

3. React中文文档：https://zh-hans.reactjs.org/docs/testing.html#jest

Jest是React社区中使用最广泛的测试框架之一。React中文文档专门介绍了如何在React项目中使用Jest进行单元测试和端对端测试。该文档还提供了许多有用的示例代码和指南，可让读者更好地理解如何使用Jest进行React测试。

4. GitHub库：https://github.com/facebook/jest

Jest是一个开源项目，它的代码存储在GitHub上。GitHub库提供了Jest核心代码的最新版本，开发者可以通过分支、拉取请求等方式对代码进行贡献。在GitHub库中还可以找到Jest的说明文档、问题跟踪、社区讨论等信息。

5. Udemy课程：https://www.udemy.com/course/react-testing-with-jest-and-enzyme/

该Udemy课程是一个更深入地学习Jest的资源。它提供了详细的课程内容，帮助学员学习如何使用Jest进行React测试。该课程包含了许多有用的示例代码和小测试项目，让学员可以实际练习Jest的使用方法。   

