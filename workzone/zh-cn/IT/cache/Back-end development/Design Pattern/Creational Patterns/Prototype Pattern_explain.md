

Prototype Pattern是一种创建型设计模式，目的是基于现有的物件或原型，通过复制来产生新的物件实例，这样就可以避免重复创建物件导致的资源浪费和效能下降。

在JavaScript中的Prototype Pattern通常是通过对象的复制或原型继承来实现的。每个对象都有一个原型对象，当需要新的对象时，只需使用现有对象的原型进行复制便可创建出一个新对象。

下面是一个Prototype Pattern的例子：

```javascript
// 定义一个食物原型对像
var foodPrototype = {
  setName: function(name) {
    this.name = name;
  },
  getName: function() {
    return this.name;
  }
};

// 创建新的食物物件
var food1 = Object.create(foodPrototype);
food1.setName('苹果');
console.log(food1.getName()); // 输出：苹果

// 创建另一个新的食物物件
var food2 = Object.create(foodPrototype);
food2.setName('橘子');
console.log(food2.getName()); // 输出：橘子
```

在上述例子中，我们定义了一个食物原型对像foodPrototype，该对像包含两个方法：setName和getName，分别用于设置和获取食物名称。然后我们使用Object.create方法基于这个原型对像创建出两个新的食物物件food1和food2，并对它们分别设置了不同的名称。这样我们就达到了避免重复创建食物物件的目的，同时也简化了代码。