

1. 实作一个函式，接收一个数字阵列并回传最大值。

```
const getMax = arr => {
  return Math.max(...arr);
}

console.log(getMax([1, 2, 3, 4, 5]));  // 5
console.log(getMax([-1, -3, -5, -10]));  // -1
```

2. 实作一个函式，接收一个数字阵列并回传所有元素的总和。

```
const getSum = arr => {
  return arr.reduce((acc, curr) => acc + curr, 0);
}

console.log(getSum([1, 2, 3, 4, 5]));  // 15
console.log(getSum([-1, -3, -5, -10]));  // -19
```

3. 实作一个函式，接收一个数字阵列并回传去重后的阵列。

```
const getUnique = arr => {
  return [...new Set(arr)];
}

console.log(getUnique([1, 2, 3, 4, 5]));  // [1, 2, 3, 4, 5]
console.log(getUnique([1, 2, 2, 3, 3, 3, 4, 5, 5]));  // [1, 2, 3, 4, 5]
```

4. 实作一个函式，接收一个字串，回传其所有单字都首字母大写的字串。

```
const formatString = str => {
  return str.split(' ')
            .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
            .join(' ');
}

console.log(formatString('hello world'));  // 'Hello World'
console.log(formatString('JAVASCRIPT IS AWESOME'));  // 'Javascript Is Awesome'
```

5. 实作一个函式，接收一个数字阵列并回传阵列中大于10的元素所组成的阵列。

```
const getGreaterThanTen = arr => {
  return arr.filter(num => num > 10);
}

console.log(getGreaterThanTen([1, 5, 10, 15, 20]));  // [15, 20]
console.log(getGreaterThanTen([2, 4, 6, 8]));  // []
```