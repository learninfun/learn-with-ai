+++
title = "JavaScript"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Javascript的中文问题
{{< /ask_chatgpt >}}



1. 如何将一个字串反转？

```
function reverseString(str) {
   return str.split("").reverse().join("");
}

console.log(reverseString("Hello")); // "olleH"
```

2. 如何取得一个Array的最大值？

```
function getMax(arr) {
  return Math.max.apply(null, arr);
}

console.log(getMax([3, 7, 2, 5, 6])); // 7
```

3. 如何检查一个变数是否为数字？

```
function isNumeric(num) {
  return !isNaN(parseFloat(num)) && isFinite(num);
}

console.log(isNumeric('123')); // true
console.log(isNumeric('abc')); // false
```

4. 如何在HTML中插入一个图片？

```
<img src="path/to/image.jpg" alt="Image Description">
```

5. 如何在Javascript中检查一个变数是否定义？

```
if (typeof myVar === 'undefined') {
  console.log("myVar is not defined");
} else {
  console.log("myVar is defined");
}
```   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Javascript并举例
{{< /ask_chatgpt >}}



JavaScript是一种前端脚本语言，是网页互动的核心之一。它可以用来编写动态效果、交互式功能和验证用户输入的表单。JavaScript可以与HTML和CSS融合在一起，并且可以在网页上运行，而不需要额外的插件或程式。

以下是JavaScript的一些示例：

1. 更改网页元素的内容：可以使用JavaScript来更改网页上的内容，例如更改标题、段落或按钮的文本内容。

2. 弹出式视窗：可以使用JavaScript来显示弹出视窗，例如警示框、确认框或提示框，来让用户进行选择或输入。

3. 网页动画效果：可以使用JavaScript来创建动画效果，例如简单的滑动、淡入淡出或完整的视差滚动效果。

4. 表单验证：可以使用JavaScript来验证用户输入的表单，例如确认密码是否匹配、检查电子邮件地址格式或限制用户输入的内容。

5. 轮播效果：可以使用JavaScript来创建轮播效果，例如自动轮播图像、手动轮播页面或无限轮播效果。

6. Ajax请求：可以使用JavaScript来发送Ajax请求，与网页上的伺服器进行交互，进行数据检索或提交表单的时候。

总之，JavaScript是一个功能丰富且简单易用的网页脚本语言，可以实现各种互动效果，可以使网站更加生动和吸引人。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Javascript的重点
{{< /ask_chatgpt >}}



1. JavaScript 是一种脚本语言，可以在网页上运行，与 HTML 和 CSS 配合，实现动态交互的效果。

2. JavaScript 的数据类型包括数字、字符串、布尔值、数组、对象等，可以通过变量、常量、运算符等进行操作。

3. JavaScript 支持基础的逻辑控制语句，如 if…else、while、for 等，还可以使用函数和事件进行代码的封装和重用。

4. JavaScript 可以与 HTML 元素进行交互，可以通过 DOM 操作对网页元素进行增、删、改、查的操作。

5. JavaScript 的异步编程模型可以通过 Promise、async/await、setTimeout 等方式实现。

6. JavaScript 还可以使用框架和库来快速实现特定的功能，如 React、Vue、jQuery 等。

7. JavaScript 的错误处理可以使用 try…catch 语句来进行捕获和处理。

8. JavaScript 可以与服务端语言进行交互，如通过 AJAX 技术来获取和发送数据。

9. JavaScript 有很多常用的内置对象和方法，如 String、Array、Math、Date 等，可以大大简化编程任务。

10. JavaScript 也有许多新的特性和语法，如箭头函数、解构赋值、模板字符串等，可以提升代码的可读性和开发效率。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Javascript的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 写一个函式，判断一个数字是否为质数。

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


2. 写一个函式，将一个整数转换成罗马数字。

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


3. 写一个函式，将一个字符串中的单词反转。

```
function reverseWords(str) {
  return str
    .split(' ')
    .map(word => word.split('').reverse().join(''))
    .join(' ');
}

console.log(reverseWords("Let's code in JavaScript!")); // "s'teL edoc ni tpircSavaJ!"
```


4. 写一个函式，移除一个数组中的重复元素。

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


5. 写一个函式，取得一个数组中第二大的数字。

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

