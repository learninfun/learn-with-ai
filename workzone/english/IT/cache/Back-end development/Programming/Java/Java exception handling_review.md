1. What is an exception in Java? 
- An exception is an event that disrupts the normal flow of a program's execution.

2. What is the difference between checked and unchecked exceptions?
- Checked exceptions are checked at compile-time and require the programmer to either handle them or declare them in the method signatures. Unchecked exceptions, on the other hand, don't need to be declared or handled explicitly.

3. What is the purpose of the try-catch-finally block in Java exception handling?
- The try block contains the code that might throw an exception, and the catch block contains the code to handle the exception. The finally block is optional and contains the code that will be executed regardless of whether an exception is thrown or not.

4. Can you throw multiple exceptions in a single method in Java?
- Yes, it is possible to throw multiple exceptions in a single method. You can use multiple catch blocks to handle each of the exceptions thrown.

5. When should you use checked exceptions instead of unchecked exceptions in Java?
- Checked exceptions should be used when the caller of a method needs to handle the exception. For example, when reading a file, the caller might need to handle exceptions like FileNotFoundException. Unchecked exceptions should be used when the caller cannot or should not handle the exception, such as in cases of programming errors or insufficient resources.