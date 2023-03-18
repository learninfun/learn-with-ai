

1. 有一个阵列 arr，里面存放了一些正整数，请写一个函式，判断是否所有元素都是偶数。

范例：

```
const arr = [2, 4, 6, 8];
const arr2 = [1, 3, 4, 7];

console.log(isAllEven(arr)); // true
console.log(isAllEven(arr2)); // false
```

答案：

```
function isAllEven(arr) {
  return arr.every(num => num % 2 === 0);
}
```

2. 请写一个函式，接收一个字串作为参数，并回传反转后的字串。

范例：

```
console.log(reverseString('hello')); // 'olleh'
console.log(reverseString('world')); // 'dlrow'
```

答案：

```
function reverseString(str) {
  return str.split('').reverse().join('');
}
```

3. 请写一个函式，接收一个阵列和一个数字，判断阵列中是否存在两个数字之和等于该数字。

范例：

```
const arr = [1, 2, 3, 4, 5];

console.log(isSumInArray(arr, 7)); // true
console.log(isSumInArray(arr, 10)); // false
```

答案：

```
function isSumInArray(arr, sum) {
  for (let i = 0; i < arr.length; i++) {
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[i] + arr[j] === sum) {
        return true;
      }
    }
  }
  return false;
}
```

4. 请写一个函式，接收一个字串作为参数，并回传该字串中每个字出现次数的统计结果。

范例：

```
console.log(countCharacters('hello')); // {h: 1, e: 1, l: 2, o: 1}
console.log(countCharacters('world')); // {w: 1, o: 1, r: 1, l: 1, d: 1}
```

答案：

```
function countCharacters(str) {
  const result = {};
  for (let i = 0; i < str.length; i++) {
    const char = str[i];
    if (!result[char]) {
      result[char] = 1;
    } else {
      result[char]++;
    }
  }
  return result;
}
```

5. 请写一个函式，接收一个物件阵列和一个属性名称作为参数，回传该属性的平均值。

范例：

```
const arr = [
  {name: 'Alice', age: 20},
  {name: 'Bob', age: 30},
  {name: 'Charlie', age: 40}
];

console.log(getAverageValue(arr, 'age')); // 30
```

答案：

```
function getAverageValue(arr, propertyName) {
  const sum = arr.reduce((acc, cur) => {
    return acc + cur[propertyName];
  }, 0);
  return sum / arr.length;
}
```