1. What is the difference between a checked and an unchecked exception in Java?
Answer: A checked exception is a type of exception that the code must handle or declare in a method's signature. Examples include IOException and ClassNotFoundException. An unchecked exception, on the other hand, is an exception that the code doesn't require handling or declaring. Examples include NullPointerException and ArrayIndexOutOfBoundsException.

2. What is the purpose of the finally block in a try-catch-finally statement?
Answer: The finally block is a block of code that gets executed regardless of whether an exception was thrown or not. This is useful for releasing resources such as database connections or file handles.

3. What is the difference between an interface and an abstract class in Java?
Answer: An interface is a contract that specifies a set of methods that a class must implement, but it does not provide any implementation details. An abstract class, on the other hand, can provide implementation details for some methods, and it may have state as well. A class can only extend one abstract class, but it can implement multiple interfaces.

4. What is the difference between the equals() and == operators in Java?
Answer: The == operator tests if two variables reference the same object in memory. The equals() method, on the other hand, tests if two objects have the same logical values. In other words, two different object instances can be equal, but they won't be the same object in memory.

5. What is the purpose of the finalize() method in Java?
Answer: The finalize() method is a method that gets called by the garbage collector when an object is about to be deleted from memory. It can be used to perform any cleanup or resource releasing activities for the object. However, it's generally not recommended to rely on this method for finalization, and it's better to use try-with-resources or other constructs to manage resource cleanup.