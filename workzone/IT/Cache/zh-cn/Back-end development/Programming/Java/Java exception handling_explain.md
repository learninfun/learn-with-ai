

Java exception handling是一种解决代码运行时遇到异常情况的机制。Java的exception handling机制允许程序在运行时处理错误，避免异常使程序崩溃并减少对代码的影响。

Java异常处理语句的结构是：

```java
try {
      //代码块
} catch (exceptionType1 e1) {
      //异常处理语句
} catch (exceptionType2 e2) {
      //异常处理语句
} catch (exceptionType3 e3) {
      //异常处理语句
} finally {
      //可选代码块
}
```

其中，try块包含可能会造成异常的代码，catch块用于处理异常类型，finally块可选，包含在执行完try和catch块之后始终执行的代码。

以下是一个Java异常处理的示例，其中将尝试读取一个不存在的文件，捕获FileNotFoundException并输出错误信息：

```java
import java.io.*;

class ExceptionExample {

     public static void main(String[] args) {

          try {
               // 打开文件
               FileInputStream file = new FileInputStream("example.txt");
          } catch (FileNotFoundException e) {
               System.out.println("找不到文件");
               e.printStackTrace();
          } 

     }
}
```

在上述代码中，当程序尝试读取一个不存在的文件时，抛出FileNotFoundException异常。try块的打开文件代码可能会抛出异常，所以我们将其置于try块中。如果FileNotFoundException异常被抛出，则catch块将被执行。在catch块中，我们输出一个错误信息并使用e.printStackTrace()方法打印异常的调用栈信息，以帮助我们更好地理解错误原因。