

1. Prototypes 是 JavaScript 中的一個重要的概念。
2. 原型鏈（Prototype Chain）是從一個對象找到它的屬性和方法的過程。
3. 繼承是一種面向對象程式設計中的重要概念，而在 JavaScript 中，原型繼承是現代 JavaScript 程式設計中的一個關鍵概念。
4. JavaScript 中的所有對象都有一個原型（prototype）屬性。
5. 使用 Object.create() 方法創建對象。
6. 調用對象的方法時，首先在本地尋找對應的方法，如果找不到就到原型鏈上繼續尋找。
7. 原型可以繼承屬性和方法，所以可以通過對原型進行修改，來影響所有繼承它的對象。
8. 原型可以用來擴展現有的對象，並且在創建新對象時，可以從現有的對象繼承原型。
9. Prototype 鏈的終點為 null，即 Object.prototype.__proto__ === null，表示沒有更高的原型繼承，Object.prototype 是所有內建對象的原型對象。
10. 改變原型對象可能會導致代碼中的不良後果，因此應該謹慎使用。