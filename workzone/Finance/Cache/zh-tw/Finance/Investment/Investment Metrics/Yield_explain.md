

Yield是一個在JavaScript中非常有用的關鍵字，它可以讓函數變成一個可暫停的生成器，可以在函數執行過程中返回多次，而不是只返回一次。

當使用yield來返回數據時，函數返回的是一個Iterator對象，該對象包含了生成器的狀態。每次調用Iterator對象的next()方法，生成器函數就會從之前暫停的位置繼續執行，直到下一次遇到yield語句為止。

以下是一個簡單的例子：

```
function* exampleGenerator() {
  yield 'Hello';
  yield 'World';
  return 'Done';
}

const iterator = exampleGenerator();
console.log(iterator.next()); // { value: 'Hello', done: false }
console.log(iterator.next()); // { value: 'World', done: false }
console.log(iterator.next()); // { value: 'Done', done: true }
```

在此例中，exampleGenerator是一個生成器函數，通過yield關鍵字，每次返回一個值。每當迭代器的next()方法被調用一次，生成器函數都會繼續從上次暫停的位置恢復執行，並且返回下一個yield的值，直到遇到return語句，此時done屬性為true，迭代器停止迭代。

通過yield，可以使用較簡單的代碼編寫出複雜的迭代邏輯，讓代碼更具可讀性和可維護性。