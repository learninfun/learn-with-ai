

1. 請問以下程式碼中的`cy.get()`會選取到幾個元素？
```
<div class="main-content">
  <div class="section">section1</div>
  <div class="section">section2</div>
  <div class="section">section3</div>
</div>
```

答案：3個元素

```cy.get('.section')```

2. 請問以下程式碼運行後，`input`元素的`值`會是什麼？
```
cy.fixture('mock-data.json').then((data) => {
  cy.get('#username').type(data.username);
  cy.get('#password').type(data.password);
});
```

答案：`input`元素的`值`會是`mock-data.json`中存儲的`username`和`password`。

3. 請問如何斷言元素是否可見？

答案：可以使用`cy.get().should('be.visible')`斷言元素是否可見。

4. 請問以下程式碼中的`cy.wait()`等待的時間是多少毫秒？
```
cy.get('.submit-btn').click();
cy.wait(2000);
```

答案：`cy.wait(2000)`等待的時間是`2000`毫秒，即2秒。

5. 請問如何獲取指定索引的元素？

答案：可以使用`cy.get().eq(index)`獲取指定索引的元素。例如`cy.get('.section').eq(0)`會獲得第一個`.section`元素。