

1. 有一個陣列，裡面的元素都是字串。請寫一個 Linter 函數，確認每一個字串都是以 "https://" 或 "http://" 開始。

答案：

```javascript
function checkProtocol(arr) {
  const regex = /^https?:\/\//;

  return arr.every(str => regex.test(str));
}

console.log(checkProtocol(['https://example.com', 'http://www.google.com'])); // true
console.log(checkProtocol(['ftp://example.com', 'http://www.google.com'])); // false
```

2. 有一個物件，裡面有幾個屬性。請寫一個 Linter 函數，確認每個屬性的值都是字串或數字型態。

答案：

```javascript
function checkType(obj) {
  const values = Object.values(obj);
  const regex = /^([0-9]+|\d+\.\d+|[a-z]+)$/i;

  return values.every(val => typeof val === 'string' || typeof val === 'number' && regex.test(val));
}

console.log(checkType({ name: 'John', age: 35, city: 'New York' })); // true
console.log(checkType({ name: 'John', age: true, city: 'New York' })); // false
```

3. 有一個陣列，裡面的元素都是字串。請寫一個 Linter 函數，確認每一個字串的長度都在 5 到 10 個字元之間。

答案：

```javascript
function checkLength(arr) {
  return arr.every(str => str.length >= 5 && str.length <= 10);
}

console.log(checkLength(['apple', 'banana', 'peach'])); // true
console.log(checkLength(['apple', 'cherry', 'peach', 'strawberry'])); // false
```

4. 有一個陣列，裡面的元素都是物件。請寫一個 Linter 函數，確認每個物件都有 "name" 和 "age" 兩個屬性。

答案：

```javascript
function checkObjectProps(arr) {
  return arr.every(obj => obj.hasOwnProperty('name') && obj.hasOwnProperty('age'));
}

console.log(checkObjectProps([
  { name: 'John', age: 35 },
  { name: 'Mary', age: 28 },
])); // true

console.log(checkObjectProps([
  { name: 'John', city: 'Taipei' },
  { age: 28 },
])); // false
```

5. 有一個陣列，裡面的元素都是數字。請寫一個 Linter 函數，確認每個數字都是 2 的冪次方。

答案：

```javascript
function checkPowerOfTwo(arr) {
  const log2 = n => Math.log(n) / Math.log(2);
  return arr.every(num => Number.isInteger(log2(num)));
}

console.log(checkPowerOfTwo([2, 4, 8, 16])); // true
console.log(checkPowerOfTwo([1, 3, 5, 10])); // false
```