

JS functional programming 是一種程式設計的風格或範式，其主要特點是使用純函數（pure functions）的方式來進行開發。

純函數是指對於相同的輸入，函數總是返回相同的輸出，且不會對全域變數產生影響。換句話說，純函數只受其輸入參數所影響，並不依賴於外部環境的任何因素。

以下是一個純函數的例子：

```
function add(a, b) {
  return a + b;
}
```

這個函數只是取兩個數相加後返回結果，沒有任何副作用（side effect）。

使用 functional programming 的好處是可以讓程式碼更易於理解、測試和重用，因為每個函數都是獨立的、沒有副作用的。

除了純函數之外，JS functional programming 還涉及到很多其他的概念，如高階函數、閉包等等。這些概念可以協助開發人員更好地組織和抽象代碼。

以下是一個使用高階函數的例子：

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

這裡定義了一個 `map` 函數，它接受一個數組和一個函數作為輸入，並返回一個新的數組，其中每個元素都是原始數組元素應用函數之後的值。這個函數就是一個高階函數，它可以接受另一個函數作為輸入。

總之，JS functional programming 是一個強大的程式設計風格，它可以幫助開發人員改善代碼的品質和效率。通過使用純函數、高階函數等概念，開發人員可以更好地組織代碼、寫出更易於理解和測試的代碼。