

Java exception handling is a built-in feature in the Java programming language, which allows developers to handle runtime errors and exceptions that occur within their code. This feature helps ensure the stability and quality of the program and enables developers to provide a readable and effective error message to users.

In Java, exception handling involves the use of try, catch, and finally blocks, which work together to detect, catch, and handle exceptions. When an exception occurs inside a try block, the catch block catches that exception and executes a specific code block that handles the error, which prevents it from propagating further up the call chain. Finally block is optional and it executes right after the catch block, even before the try block exits.

Here is an example of Java exception handling code:

```java

public class ExceptionExample {
    public static void main(String[] args) {

        try {
            int[] data = new int[5];
            data[5] = 10; // This produces an ArrayIndexOutOfBoundsException
        } 
        catch (ArrayIndexOutOfBoundsException e) {
            System.out.println("Error: Array out of bounds.");
        } 
        finally {
            System.out.println("Finally block always executes.");
        }

        System.out.println("Program continues to run");
    }
}
```

In the above example, we have created an array of size 5, but we are trying to set the value of the 6th element, which causes an ArrayIndexOutOfBoundsException. The catch block detects the error and prints an error message. The finally block is optional and executes regardless of whether the exception occurred or not. 

The output of this program will be:

```
Error: Array out of bounds.
Finally block always executes.
Program continues to run
```

In the above output, you can see that the error message has been printed, and the program has continued to run, thanks to exception handling.