

A JS module is a self-contained block of code that performs a specific task or defines a specific functionality. It enables the encapsulation of code, which facilitates better organization, easier maintenance and reuse.

An example of a JS module would be a library that provides a set of utility functions, such as validation of email addresses, formatting of dates or calculation of math operations. Here is a basic example of a utility function as a module in JavaScript:

```javascript
// Define a module that returns the sum of two numbers
const sumModule = (function() {
  function add(a, b) {
    return a + b;
  }
  return { add };
})();

// Use the module by calling the function
console.log(sumModule.add(5, 3)); // Output: 8
```

In this example, the `sumModule` encapsulates the `add()` function, which returns the sum of two numbers. The module is defined using an IIFE (Immediately-invoked Function Expression), which immediately executes the enclosed function and returns an object containing the `add()` function. The `sumModule` object can be then used to call the `add()` function, which computes the sum of two numbers.