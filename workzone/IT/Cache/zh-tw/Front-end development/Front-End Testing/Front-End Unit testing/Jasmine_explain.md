

Jasmine是一個用於JavaScript測試的行為驅動開發（BDD）框架。它的設計目的是使JavaScript測試變得簡單明了，並且能夠直觀地表達代碼的預期行為。

在Jasmine中，測試以“描述”（describe）和“斷言”（expect）的形式進行。描述是對代碼功能的解釋性的文字描述，可以嵌套描述和斷言，進一步剖析代碼的不同層面。 斷言是一個在測試中對代碼的預期結果的陳述。Jasmine提供了多種斷言方法，比如toBe，toEqual，toContain等等。

舉個例子，假設我們有一個JavaScript函數，例如：

```
function addNumbers(a, b) {
  return a + b;
}
```

要測試這個函數的行為，我們可以使用Jasmine。我們可以寫一個描述來解釋此函數的功能，然後寫一個斷言來驗證函數是否正確，如下所示：

```
describe('addNumbers function', function() {
  
  it('adds two numbers and returns the correct result', function() {
    var result = addNumbers(2, 3);
    expect(result).toEqual(5);
  });
  
});
```

在這個示例中，我們編寫了一個描述來說明這個測試是關於“addNumbers function”函數的。在描述的內部，我們定義了一個it（也稱為“測試用例”），在此測試用例中，我們使用了addNumbers函數來加入兩個數字，然後使用Jasmine的expect斷言來驗證這個函數的返回值是否為5。如果測試通過，Jasmine將會顯示一個綠色符號表示成功，否則將會顯示一個紅色符號表示失敗。