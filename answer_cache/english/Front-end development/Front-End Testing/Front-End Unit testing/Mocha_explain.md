

Mocha is a JavaScript testing framework used for unit testing, and it runs on Node.js and the browser. It makes it easier to write test cases and generates clear, detailed test results.

An example of using Mocha would be:

```
const assert = require('assert');

describe('Math', function() {
  describe('#addition()', function() {
    it('should return 2 when adding 1 and 1', function() {
      assert.equal(1 + 1, 2);
    });
    
    it('should return 0 when adding -1 and 1', function() {
      assert.equal(-1 + 1, 0);
    });
  });
  
  describe('#subtraction()', function() {
    it('should return 0 when subtracting 1 from 1', function() {
      assert.equal(1 - 1, 0);
    });
    
    it('should return -1 when subtracting 1 from 0', function() {
      assert.equal(0 - 1, -1);
    });
  });
});
```

In this example, we are testing the Math object's addition() and subtraction() functions. Mocha provides a describe() function to group test cases together and an it() function to define individual test cases. The assert() function checks that the expected value matches the actual value. When the tests are run, Mocha will report the results of each test case, making it easy to identify errors in the code.