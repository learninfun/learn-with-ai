

1) 以下是一个计算器的 add 函数，请写测试用例来测试该函数是否正确：

```js
function add(a, b) {
  return a + b;
}
```

答案：

```js
test("add function works correctly", () => {
  expect(add(2, 2)).toBe(4);
  expect(add(0, 4)).toBe(4);
  expect(add(-2, 3)).toBe(1);
  expect(add(1.2, 2.3)).toBeCloseTo(3.5);
});
```

2) 以下是一个输入金额，计算总价的函数，请写测试用例来测试该函数是否正确：

```js
function calculateTotalPrice(price, quantity) {
  const total = price * quantity;
  return `Total price is ${total}`;
}
```

答案：

```js
test("calculateTotalPrice function works correctly", () => {
  expect(calculateTotalPrice(10, 3)).toBe("Total price is 30");
  expect(calculateTotalPrice(5, 0)).toBe("Total price is 0");
  expect(calculateTotalPrice(2.5, 4)).toBe("Total price is 10");
  expect(calculateTotalPrice(8.99, 1)).toBe("Total price is 8.99");
});
```

3) 以下是一个判断年份是否为闰年的函数，请写测试用例来测试该函数是否正确：

```js
function isLeapYear(year) {
  if (year % 4 === 0 && (year % 100 !== 0 || year % 400 === 0)) {
    return true;
  } else {
    return false;
  }
}
```

答案：

```js
test("isLeapYear function works correctly", () => {
  expect(isLeapYear(2000)).toBe(true);
  expect(isLeapYear(1900)).toBe(false);
  expect(isLeapYear(2020)).toBe(true);
  expect(isLeapYear(2022)).toBe(false);
});
```

4) 以下是一个输入条件，判断是否符合注册要求的函数，请写测试用例来测试该函数是否正确：

```js
function isRegisterValid(username, password, confirmPassword) {
  if (username.length >= 3 && password.length >= 6 && password === confirmPassword) {
    return true;
  } else {
    return false;
  }
}
```

答案：

```js
test("isRegisterValid function works correctly", () => {
  expect(isRegisterValid("john", "123456", "123456")).toBe(true);
  expect(isRegisterValid("joe123", "password", "password")).toBe(false);
  expect(isRegisterValid("user", "123456", "654321")).toBe(false);
  expect(isRegisterValid("admin", "password", "password")).toBe(false);
});
```

5) 以下是一个输入年月日，计算星期几的函数，请写测试用例来测试该函数是否正确：

```js
function getDayOfWeek(year, month, day) {
  const daysOfWeek = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"];
  const date = new Date(`${month}/${day}/${year}`);
  const dayOfWeek = daysOfWeek[date.getDay()];
  return dayOfWeek;
}
```

答案：

```js
test("getDayOfWeek function works correctly", () => {
  expect(getDayOfWeek(2022, 11, 16)).toBe("Wednesday");
  expect(getDayOfWeek(1970, 1, 1)).toBe("Thursday");
  expect(getDayOfWeek(2021, 7, 4)).toBe("Sunday");
  expect(getDayOfWeek(2030, 12, 25)).toBe("Wednesday");
});
```