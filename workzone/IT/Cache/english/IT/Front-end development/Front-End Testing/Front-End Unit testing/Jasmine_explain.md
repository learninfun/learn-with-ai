

Jasmine is a behavior-driven development framework for testing JavaScript code. It provides a readable, user-friendly syntax that allows developers to write automated tests for their applications.

The main idea behind Jasmine is to describe expected behavior in a clear and concise language. This makes it easier to write and maintain test cases, as well as to understand the intended functionalities of the code being tested.

Here's an example of a Jasmine test case:

```javascript
describe("Calculator", function() {
  var calc;

  beforeEach(function() {
    calc = new Calculator();
  });

  describe("addition", function() {
    it("should be able to add two numbers", function() {
      expect(calc.addition(2, 3)).toEqual(5);
    });

    it("should return NaN if one of the inputs is not a number", function() {
      expect(calc.addition("2", 3)).toBeNaN();
    });
  });

  describe("subtraction", function() {
    it("should be able to subtract two numbers", function() {
      expect(calc.subtraction(5, 3)).toEqual(2);
    });

    it("should return NaN if one of the inputs is not a number", function() {
      expect(calc.subtraction(5, "3")).toBeNaN();
    });
  });
});
```

In this test case, we're testing a Calculator object that has two methods, addition and subtraction. We're checking that these methods behave as expected by testing their output against known values. We're also testing that the methods return NaN (Not a Number) when non-numeric input is passed to them. 

Using Jasmine, we can run these tests automatically every time we make changes to our code, ensuring that it continues to behave as expected.