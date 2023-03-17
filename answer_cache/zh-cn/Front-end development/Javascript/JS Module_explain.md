

JS Module是一种将复杂的代码分解成独立的功能单元，方便维护和协作的技术。在JS中，Module通常指代一个独立的js文件，这个文件中将某个特定的功能进行了封装，并提供了对外的接口。其他js文件可以通过这些暴露出来的接口来使用这个模块。

举例来说，有一个名为"utils.js"的模块，里面定义了一些函数用于辅助其他模块的开发。

```javascript
// utils.js
export function add(a, b) {
  return a + b;
}

export function substract(a, b) {
  return a - b;
}

export function multiply(a, b) {
  return a * b;
}
```

其他模块可以使用"utils.js"中提供的函数，只需要在文件中引入即可。

```javascript
// main.js
import { add, substract } from './utils';

console.log(add(1, 2)); // 3
console.log(substract(5, 3)); // 2
```

JS模组化能够提高代码的可读性和可维护性，有助于团队协作和代码重复利用，是现代Web开发中不可缺少的一部分技术。