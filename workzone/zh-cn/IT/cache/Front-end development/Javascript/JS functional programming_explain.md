

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