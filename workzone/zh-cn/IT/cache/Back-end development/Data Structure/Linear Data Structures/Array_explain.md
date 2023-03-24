

Array（陣列）是一種儲存多個相同類型值的資料結構。它可以在一個變數中存放多個值，並且每一個值都可以通過索引來訪問。索引通常是整數型別，它從 0 開始，依次增加。

以下是一些 Array 的例子：

1. 儲存一組數值

```
let numbers = [1, 2, 3, 4, 5];
console.log(numbers); // [1, 2, 3, 4, 5]
```

2. 儲存多個字符串

```
let fruits = ["apple", "banana", "orange", "grape"];
console.log(fruits); // ["apple", "banana", "orange", "grape"]
```

3. 儲存布林值

```
let bools = [true, false, true];
console.log(bools); // [true, false, true]
```

4. 儲存對象

```
let users = [
  { name: "Amy", age: 25 },
  { name: "Bob", age: 30 },
  { name: "Carl", age: 28 }
];
console.log(users); // [{name: "Amy", age: 25}, {name: "Bob", age: 30}, {name: "Carl", age: 28}]
```

需要注意的是，Array 有一些屬性和方法，可以對其中的值進行操作，比如 `length` 屬性可以返回陣列的大小，`push` 方法可以在陣列末尾添加一個元素，`splice` 方法可以在指定位置插入或刪除元素等。