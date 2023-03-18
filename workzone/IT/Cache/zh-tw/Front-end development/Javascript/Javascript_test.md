

1. 寫一個函式，判斷一個數字是否為質數。

```
function isPrime(num) {
  if (num <= 1) return false;
  for (let i = 2; i <= Math.sqrt(num); i++) {
    if (num % i === 0) return false;
  }
  return true;
}

console.log(isPrime(11)); // true
console.log(isPrime(12)); // false
```


2. 寫一個函式，將一個整數轉換成羅馬數字。

```
function intToRoman(num) {
  const mapping = {
    M: 1000,
    CM: 900,
    D: 500,
    CD: 400,
    C: 100,
    XC: 90,
    L: 50,
    XL: 40,
    X: 10,
    IX: 9,
    V: 5,
    IV: 4,
    I: 1
  };

  let result = '';
  for (let key in mapping) {
    while (num >= mapping[key]) {
      result += key;
      num -= mapping[key];
    }
  }

  return result;
}

console.log(intToRoman(1994)); // "MCMXCIV"
```


3. 寫一個函式，將一個字符串中的單詞反轉。

```
function reverseWords(str) {
  return str
    .split(' ')
    .map(word => word.split('').reverse().join(''))
    .join(' ');
}

console.log(reverseWords("Let's code in JavaScript!")); // "s'teL edoc ni tpircSavaJ!"
```


4. 寫一個函式，移除一個數組中的重複元素。

```
function removeDuplicates(nums) {
  let i = 0;
  for (let j = 1; j < nums.length; j++) {
    if (nums[i] !== nums[j]) {
      i++;
      nums[i] = nums[j];
    }
  }
  
  return nums.slice(0, i + 1);
}

console.log(removeDuplicates([1,1,2,2,3,4,4,5,5,5])); // [1, 2, 3, 4, 5]
```


5. 寫一個函式，取得一個數組中第二大的數字。

```
function secondLargest(nums) {
  let max = -Infinity;
  let secondMax = -Infinity;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] > max) {
      secondMax = max;
      max = nums[i];
    } else if (nums[i] > secondMax && nums[i] !== max) {
      secondMax = nums[i];
    }
  }
  
  return secondMax;
}

console.log(secondLargest([3, 1, 5, 9, 2, 7])); // 7
```