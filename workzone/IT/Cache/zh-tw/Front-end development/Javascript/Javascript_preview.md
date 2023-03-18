

1. 如何將一個字串反轉？

```
function reverseString(str) {
   return str.split("").reverse().join("");
}

console.log(reverseString("Hello")); // "olleH"
```

2. 如何取得一個Array的最大值？

```
function getMax(arr) {
  return Math.max.apply(null, arr);
}

console.log(getMax([3, 7, 2, 5, 6])); // 7
```

3. 如何檢查一個變數是否為數字？

```
function isNumeric(num) {
  return !isNaN(parseFloat(num)) && isFinite(num);
}

console.log(isNumeric('123')); // true
console.log(isNumeric('abc')); // false
```

4. 如何在HTML中插入一個圖片？

```
<img src="path/to/image.jpg" alt="Image Description">
```

5. 如何在Javascript中檢查一個變數是否定義？

```
if (typeof myVar === 'undefined') {
  console.log("myVar is not defined");
} else {
  console.log("myVar is defined");
}
```