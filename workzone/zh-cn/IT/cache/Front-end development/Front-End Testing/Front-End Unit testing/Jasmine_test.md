

1. 给定一个整数 n，输出所有小于 n 的正整数中，能被 3 或 5 整除的数字的总和。
   范例输入：10
   范例输出：8 (3 + 5)

2. 给定一个整数阵列 nums 和一个目标整数 target，找出 nums 中两个数字的和等于 target，并返回它们的索引。
   假设每个输入仅有一个解，且同一个元素不能使用两次。
   范例输入：nums = [2, 7, 11, 15], target = 9
   范例输出：[0, 1]

3. 实现一个费伯纳契数列的函数。费伯纳契数列是一系列的整数，其中每个数字都是前两个数字的和。
   范例输入：10
   范例输出：[0, 1, 1, 2, 3, 5, 8, 13, 21, 34]

4. 实现一个函数，将一个给定的字符串转换为小写。
   范例输入：Jasmine
   范例输出：jasmine

5. 给定一个整数 x，实现一个函数，计算 x 的平方根。如果 x 是负数，则返回 null。
   范例输入：16
   范例输出：4.0

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