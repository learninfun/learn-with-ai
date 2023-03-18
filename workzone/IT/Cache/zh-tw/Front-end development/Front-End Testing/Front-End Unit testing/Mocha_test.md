

1. 有一個陣列 arr，裡面存放了一些正整數，請寫一個函式，判斷是否所有元素都是偶數。

範例：

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

2. 請寫一個函式，接收一個字串作為參數，並回傳反轉後的字串。

範例：

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

3. 請寫一個函式，接收一個陣列和一個數字，判斷陣列中是否存在兩個數字之和等於該數字。

範例：

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

4. 請寫一個函式，接收一個字串作為參數，並回傳該字串中每個字出現次數的統計結果。

範例：

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

5. 請寫一個函式，接收一個物件陣列和一個屬性名稱作為參數，回傳該屬性的平均值。

範例：

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