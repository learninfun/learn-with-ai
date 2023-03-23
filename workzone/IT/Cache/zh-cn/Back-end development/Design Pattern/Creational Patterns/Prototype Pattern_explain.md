

Prototype Pattern是一種創建型設計模式，目的是基於現有的物件或原型，通過複製來產生新的物件實例，這樣就可以避免重複創建物件導致的資源浪費和效能下降。

在JavaScript中的Prototype Pattern通常是通過對象的複製或原型繼承來實現的。每個對象都有一個原型對象，當需要新的對象時，只需使用現有對象的原型進行複製便可創建出一個新對象。

下面是一個Prototype Pattern的例子：

```javascript
// 定義一個食物原型對像
var foodPrototype = {
  setName: function(name) {
    this.name = name;
  },
  getName: function() {
    return this.name;
  }
};

// 創建新的食物物件
var food1 = Object.create(foodPrototype);
food1.setName('蘋果');
console.log(food1.getName()); // 輸出：蘋果

// 創建另一個新的食物物件
var food2 = Object.create(foodPrototype);
food2.setName('橘子');
console.log(food2.getName()); // 輸出：橘子
```

在上述例子中，我們定義了一個食物原型對像foodPrototype，該對像包含兩個方法：setName和getName，分別用於設置和獲取食物名稱。然後我們使用Object.create方法基於這個原型對像創建出兩個新的食物物件food1和food2，並對它們分別設置了不同的名稱。這樣我們就達到了避免重複創建食物物件的目的，同時也簡化了代碼。