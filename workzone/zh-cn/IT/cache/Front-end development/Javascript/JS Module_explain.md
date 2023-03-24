

JS Module是一種將複雜的代碼分解成獨立的功能單元，方便維護和協作的技術。在JS中，Module通常指代一個獨立的js文件，這個文件中將某個特定的功能進行了封裝，並提供了對外的接口。其他js文件可以通過這些暴露出來的接口來使用這個模塊。

舉例來說，有一個名為"utils.js"的模塊，裡面定義了一些函數用於輔助其他模塊的開發。

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

其他模塊可以使用"utils.js"中提供的函數，只需要在文件中引入即可。

```javascript
// main.js
import { add, substract } from './utils';

console.log(add(1, 2)); // 3
console.log(substract(5, 3)); // 2
```

JS模組化能夠提高代碼的可讀性和可維護性，有助於團隊協作和代碼重複利用，是現代Web開發中不可缺少的一部分技術。