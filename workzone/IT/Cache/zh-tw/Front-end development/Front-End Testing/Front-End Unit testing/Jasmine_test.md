

1. 給定一個整數 n，輸出所有小於 n 的正整數中，能被 3 或 5 整除的數字的總和。
   範例輸入：10
   範例輸出：8 (3 + 5)

2. 給定一個整數陣列 nums 和一個目標整數 target，找出 nums 中兩個數字的和等於 target，並返回它們的索引。
   假設每個輸入僅有一個解，且同一個元素不能使用兩次。
   範例輸入：nums = [2, 7, 11, 15], target = 9
   範例輸出：[0, 1]

3. 實現一個費伯納契數列的函數。費伯納契數列是一系列的整數，其中每個數字都是前兩個數字的和。
   範例輸入：10
   範例輸出：[0, 1, 1, 2, 3, 5, 8, 13, 21, 34]

4. 實現一個函數，將一個給定的字符串轉換為小寫。
   範例輸入：Jasmine
   範例輸出：jasmine

5. 給定一個整數 x，實現一個函數，計算 x 的平方根。如果 x 是負數，則返回 null。
   範例輸入：16
   範例輸出：4.0

答案：
1. 
function find_sum(n){
  let sum = 0;
  for (let i = 1; i < n; i++){
    if (i % 3 === 0 || i % 5 === 0){
      sum += i;
    }
  }
  return sum;
}

2. 
function twoSum(nums, target) {
  let map = new Map();
  for (let i = 0; i < nums.length; i++){
    let complement = target - nums[i];
    if (map.has(complement)){
      return [map.get(complement), i];
    }
    map.set(nums[i], i);
  }
}

3. 
function fibonacci(n) {
  let res = [];
  let a = 0;
  let b = 1;
  for (let i = 0; i < n; i++){
    res.push(a);
    let temp = a + b;
    a = b;
    b = temp;
  }
  return res;
}

4. 
function toLowerCase(str) {
  return str.toLowerCase();
}

5. 
function sqrt(x) {
  if (x < 0) return null;
  let left = 0;
  let right = x;
  while (left <= right) {
    let mid = Math.floor((left + right) / 2);
    if (mid * mid === x) return mid;
    if (mid * mid < x) left = mid + 1;
    if (mid * mid > x) right = mid - 1;
  }
  return right;
}