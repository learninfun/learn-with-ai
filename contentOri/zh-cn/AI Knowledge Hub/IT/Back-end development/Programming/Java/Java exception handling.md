+++
title = "Java异常处理"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Java exception handling的问题
{{< /ask_chatgpt >}}



1. 什麼是Java exception handling？如何利用try-catch block處理異常？

2. 請舉例說明什麼是unchecked exception和checked exception？如何處理它們？

3. 請列舉幾種Java異常處理中的關鍵字及其作用？

4. 如何設計自己的Java異常類別？請使用代碼示例說明。

5. 請比較throw和throws在Java異常處理中的區別，並舉例說明如何使用它們。   

## 说明知识
{{< ask_chatgpt >}}
说明Java exception handling并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列Java exception handling的重点
{{< /ask_chatgpt >}}



1. 知道哪些异常（Exceptions）需要处理（Handled Exceptions）；
2. 使用try-catch区块来处理异常；
3. 了解异常的继承关系（Exception Hierarchies）；
4. 了解发生异常会如何影响程式流程；
5. 使用finally区块来保护资源不会被释放掉；
6. 遵循Java API惯例，根据异常型别和程式逻辑，组织try-catch区块。
7. 了解何时可以使用throws关键字来处理异常。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Java exception handling的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 问题：请写一个方法，接收一个整数作为参数，若该参数小于0，则抛出一个IllegalArgumentException，否则返回该参数的平方。
答案：

public int getSquare(int num) {
    if (num < 0) { //检查参数是否小于0
        throw new IllegalArgumentException("The input number should be greater than or equal to 0.");
    }
    return num * num;
}

2. 问题：请写一个方法，接收一个字串作为参数，若该字串为空或null，则抛出一个NullPointerException，否则返回该字串的长度。
答案：

public int getLength(String str) {
    if (str == null || str.isEmpty()) { //检查字串是否为空或null
        throw new NullPointerException("The input string cannot be null or empty.");
    }
    return str.length();
}

3. 问题：请写一个方法，接收一个整数阵列和一个索引值作为参数，若索引值小于0或大于等于阵列长度，则抛出一个IndexOutOfBoundsException，否则返回该索引所对应的阵列元素。
答案：

public int getElement(int[] arr, int index) {
    if (index < 0 || index >= arr.length) { //检查索引值是否越界
        throw new IndexOutOfBoundsException("The index is out of bounds.");
    }
    return arr[index];
}

4. 问题：请写一个方法，接收一个字串和一个正整数作为参数，若该字串长度小于等于该正整数，则抛出一个RuntimeException，否则返回该字串的前n个字元。
答案：

public String getSubstring(String str, int n) {
    if (str.length() <= n) { //检查字串长度是否足够
        throw new RuntimeException("The length of the input string is not greater than " + n + ".");
    }
    return str.substring(0, n);
}

5. 问题：请写一个方法，接收一个整数作为参数，若该整数为奇数，则抛出一个Checked Exception，否则返回该整数的平方。
答案：

public int getSquareEven(int num) throws Exception {
    if (num % 2 != 0) { //检查整数是否为奇数
        throw new Exception("The input number should be even.");
    }
    return num * num;
}   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Java exception handling的网络数据
{{< /ask_chatgpt >}}



1. Oracle Java Tutorial: Exception Handling
https://docs.oracle.com/javase/tutorial/essential/exceptions/index.html

This is the official Java Tutorial from Oracle, providing a comprehensive overview of exception handling in Java. It covers topics such as the try-catch statement, multiple catch blocks, nested try statements, and finally blocks.

2. Java Exceptions - Baeldung
https://www.baeldung.com/java-exceptions-tutorial

This article from Baeldung provides a practical guide to Java exception handling, with examples of how to handle exceptions in different scenarios. It covers topics such as checked and unchecked exceptions, the try-with-resources statement, and custom exceptions.

3. Java Exception Handling - TutorialsPoint
https://www.tutorialspoint.com/java/java_exceptions.htm

TutorialsPoint provides a beginner's guide to Java exception handling, introducing the concept of exceptions and explaining how to handle them in Java. It covers topics such as the try-catch block, the throw statement, and the finally block.

4. Handling Exceptions in Java - GeeksforGeeks
https://www.geeksforgeeks.org/exceptions-in-java/

GeeksforGeeks provides an overview of Java exceptions, including the types of exceptions and how to handle them. It covers topics such as exception hierarchy, throwing custom exceptions, and exception propagation.

5. Exception Handling Best Practices in Java - DZone
https://dzone.com/articles/exception-handling-best-practices-in-java

This article from DZone provides tips and best practices for effective exception handling in Java. It covers topics such as handling checked and unchecked exceptions, throwing exceptions, and logging exceptions. It also includes examples of common exception handling patterns.   
