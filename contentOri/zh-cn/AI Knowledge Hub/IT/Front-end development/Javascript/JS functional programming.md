+++
title = "JavaScript 函数式编程"
weight = "7"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题JS functional programming的中文问题
{{< /ask_chatgpt >}}



1. 如何实现一个 map 函数，用于对数组中的每个元素进行指定操作，并返回新的数组？

2. 如何使用 reduce 函数实现数组中所有数字的加总？

3. 如何使用 filter 函数实现选择数组中所有大于10的数字？

4. 如何使用 compose 函数把多个函数串接起来，实现复合函数？

5. 如何使用柯里化（currying）实现一个通用的 add 函数，能够进行任意数字的相加？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明JS functional programming并举例
{{< /ask_chatgpt >}}



JS functional programming 是一种程式设计的风格或范式，其主要特点是使用纯函数（pure functions）的方式来进行开发。

纯函数是指对于相同的输入，函数总是返回相同的输出，且不会对全域变数产生影响。换句话说，纯函数只受其输入参数所影响，并不依赖于外部环境的任何因素。

以下是一个纯函数的例子：

```
function add(a, b) {
  return a + b;
}
```

这个函数只是取两个数相加后返回结果，没有任何副作用（side effect）。

使用 functional programming 的好处是可以让程式码更易于理解、测试和重用，因为每个函数都是独立的、没有副作用的。

除了纯函数之外，JS functional programming 还涉及到很多其他的概念，如高阶函数、闭包等等。这些概念可以协助开发人员更好地组织和抽像代码。

以下是一个使用高阶函数的例子：

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

这里定义了一个 `map` 函数，它接受一个数组和一个函数作为输入，并返回一个新的数组，其中每个元素都是原始数组元素应用函数之后的值。这个函数就是一个高阶函数，它可以接受另一个函数作为输入。

总之，JS functional programming 是一个强大的程式设计风格，它可以帮助开发人员改善代码的品质和效率。通过使用纯函数、高阶函数等概念，开发人员可以更好地组织代码、写出更易于理解和测试的代码。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列JS functional programming的重点
{{< /ask_chatgpt >}}



1. 价值原则：函式是值，而不是命令式的指令。

2. 纯函式：函式对相同的引数输入返回相同的结果，并且没有副作用。

3. 不可变性：数据在创建后保持不变。

4. 好的过程：函式应该是被设计成尽可能简单明了、易于阅读和测试。

5. 组合性：函式应该能够被组合在一起以创建更复杂的函式。

6. 高阶函式：函式可以接受函式作为引数，并且可以返回一个函式。

7. 递归：使用递归解决问题，可以使代码更简洁和易于阅读。

8. 类型推断：类型应该基于值和行为的推断来进行，而不是显式声明。

9. 惰性求值：仅计算必须使用的值，可以使代码更有效率。

10. 柯里化与部分应用：柯里化和部分应用技术可以简化函式的实现，使其更容易理解和使用。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题JS functional programming的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实作一个函式，接收一个数字阵列并回传最大值。

```
const getMax = arr => {
  return Math.max(...arr);
}

console.log(getMax([1, 2, 3, 4, 5]));  // 5
console.log(getMax([-1, -3, -5, -10]));  // -1
```

2. 实作一个函式，接收一个数字阵列并回传所有元素的总和。

```
const getSum = arr => {
  return arr.reduce((acc, curr) => acc + curr, 0);
}

console.log(getSum([1, 2, 3, 4, 5]));  // 15
console.log(getSum([-1, -3, -5, -10]));  // -19
```

3. 实作一个函式，接收一个数字阵列并回传去重后的阵列。

```
const getUnique = arr => {
  return [...new Set(arr)];
}

console.log(getUnique([1, 2, 3, 4, 5]));  // [1, 2, 3, 4, 5]
console.log(getUnique([1, 2, 2, 3, 3, 3, 4, 5, 5]));  // [1, 2, 3, 4, 5]
```

4. 实作一个函式，接收一个字串，回传其所有单字都首字母大写的字串。

```
const formatString = str => {
  return str.split(' ')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
            .join(' ');
}

console.log(formatString('hello world'));  // 'Hello World'
console.log(formatString('JAVASCRIPT IS AWESOME'));  // 'Javascript Is Awesome'
```

5. 实作一个函式，接收一个数字阵列并回传阵列中大于10的元素所组成的阵列。

```
const getGreaterThanTen = arr => {
  return arr.filter(num => num > 10);
}

console.log(getGreaterThanTen([1, 5, 10, 15, 20]));  // [15, 20]
console.log(getGreaterThanTen([2, 4, 6, 8]));  // []
```   

