

Babel是一個JavaScript編譯器，用於轉換新版本的JavaScript代碼為較舊的版本，以確保它們能夠在各種瀏覽器和環境中運行。

例如，Babel可以將以下新的ES6代碼：

```
const greeting = (name) => {
  return `Hello, ${name}!`;
}

```

轉換為舊版的ES5代碼：

```
"use strict";

var greeting = function greeting(name) {
  return "Hello, " + name + "!";
};
```

這樣做可以確保這段代碼在舊版瀏覽器和其它環境中運行良好而不會出現錯誤。Babel還可以支援其他語言功能，如TypeScript和Flow。 通過Babel創建的JavaScript代碼可以更容易地在不同環境中運行，並且更簡潔、可讀性高。