

1. 使用 Omega Notation來表示下列函式最低的複雜度：
   ```
   function search(arr, x) {
     for (let i = 0; i < arr.length; i++) {
       if (arr[i] === x) {
         return i;
       }
     }
     return -1;
   }
   ```
   答案：Ω(1)

2. 使用 Omega Notation來表示下列函式最低的複雜度：
   ```
   function bubbleSort(arr) {
     for (let i = 0; i < arr.length; i++) {
       for (let j = 0; j < arr.length - i - 1; j++) {
         if (arr[j] > arr[j + 1]) {
           [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
         }
       }
     }
     return arr;
   }
   ```
   答案：Ω(n)

3. 使用 Omega Notation來表示下列函式最低的複雜度：
   ```
   function mergeSort(arr) {
     if (arr.length <= 1) {
       return arr;
     }
     const mid = Math.floor(arr.length / 2);
     const leftArr = mergeSort(arr.slice(0, mid));
     const rightArr = mergeSort(arr.slice(mid));
     return merge(leftArr, rightArr);
   }
   
   function merge(arr1, arr2) {
     const resultArr = [];
     let i = 0;
     let j = 0;
     while (i < arr1.length && j < arr2.length) {
       if (arr1[i] < arr2[j]) {
         resultArr.push(arr1[i]);
         i++;
       } else {
         resultArr.push(arr2[j]);
         j++;
       }
     }
     if (i < arr1.length) {
       resultArr.push(...arr1.slice(i));
     }
     if (j < arr2.length) {
       resultArr.push(...arr2.slice(j));
     }
     return resultArr;
   }
   ```
   答案：Ω(n log n)

4. 使用 Omega Notation來表示下列函式最低的複雜度：
   ```
   function fibonacci(n) {
     if (n <= 1) {
       return n;
     }
     return fibonacci(n - 1) + fibonacci(n - 2);
   }
   ```
   答案：Ω(φ^n) (φ為黃金比例，約為1.618)

5. 使用 Omega Notation來表示下列函式最低的複雜度：
   ```
   function binarySearch(arr, x) {
     let left = 0;
     let right = arr.length - 1;
     while (left <= right) {
       const mid = Math.floor((left + right) / 2);
       if (arr[mid] === x) {
         return mid;
       } else if (arr[mid] < x) {
         left = mid + 1;
       } else {
         right = mid - 1;
       }
     }
     return -1;
   }
   ```
   答案：Ω(log n)