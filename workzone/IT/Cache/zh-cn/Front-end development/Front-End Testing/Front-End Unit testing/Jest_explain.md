

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