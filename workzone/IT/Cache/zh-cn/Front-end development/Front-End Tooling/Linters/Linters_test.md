

1. 有一个阵列，里面的元素都是字串。请写一个 Linter 函数，确认每一个字串都是以 "https://" 或 "http://" 开始。

答案：

```javascript
function checkProtocol(arr) {
  const regex = /^https?:\/\//;

  return arr.every(str => regex.test(str));
}

console.log(checkProtocol(['https://example.com', 'http://www.google.com'])); // true
console.log(checkProtocol(['ftp://example.com', 'http://www.google.com'])); // false
```

2. 有一个物件，里面有几个属性。请写一个 Linter 函数，确认每个属性的值都是字串或数字型态。

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

3. 有一个阵列，里面的元素都是字串。请写一个 Linter 函数，确认每一个字串的长度都在 5 到 10 个字元之间。

答案：

```javascript
function checkLength(arr) {
  return arr.every(str => str.length >= 5 && str.length <= 10);
}

console.log(checkLength(['apple', 'banana', 'peach'])); // true
console.log(checkLength(['apple', 'cherry', 'peach', 'strawberry'])); // false
```

4. 有一个阵列，里面的元素都是物件。请写一个 Linter 函数，确认每个物件都有 "name" 和 "age" 两个属性。

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

5. 有一个阵列，里面的元素都是数字。请写一个 Linter 函数，确认每个数字都是 2 的幂次方。

答案：

```javascript
function checkPowerOfTwo(arr) {
  const log2 = n => Math.log(n) / Math.log(2);
  return arr.every(num => Number.isInteger(log2(num)));
}

console.log(checkPowerOfTwo([2, 4, 8, 16])); // true
console.log(checkPowerOfTwo([1, 3, 5, 10])); // false
```