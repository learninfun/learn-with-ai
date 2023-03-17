

ESLint是一個開源的靜態代碼檢查工具，用於找出JavaScript中的問題並將其標示為錯誤、警告或建議修復。它可以幫助開發者遵循一致的代碼風格，並在開發過程中捕捉潛在的錯誤。

ESLint可以在命令行中運行，也可以作為集成到編輯器中的插件使用。它支持許多不同的JavaScript編程風格和框架，並有龐大的社區支持。

以下是一個使用ESLint的示例：

```javascript
function calculateSum(a, b) {
  return a + b;
}

calculateSum(1, 2);
```

在這個例子中，我們使用了一個非常簡單的函數來計算兩個數字的總和。我們可以使用ESLint來檢查它是否有任何問題。在命令行中運行以下命令：

```
eslint calculateSum.js
```

結果將顯示任何問題。如果我們在此函數中添加類型錯誤或錯誤的JavaScript語法，ESLint將標記它們為錯誤或警告，以幫助開發者更快地找到問題。