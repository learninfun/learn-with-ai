

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