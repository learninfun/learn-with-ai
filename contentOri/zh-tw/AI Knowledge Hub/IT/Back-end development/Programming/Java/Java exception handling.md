## 習題預習
{{< ask_chatgpt >}}
給我5題Java exception handling的問題
{{< /ask_chatgpt >}}



1. 什麼是Java exception handling？如何利用try-catch block處理異常？

2. 請舉例說明什麼是unchecked exception和checked exception？如何處理它們？

3. 請列舉幾種Java異常處理中的關鍵字及其作用？

4. 如何設計自己的Java異常類別？請使用代碼示例說明。

5. 請比較throw和throws在Java異常處理中的區別，並舉例說明如何使用它們。   

## 說明知識
{{< ask_chatgpt >}}
說明Java exception handling並舉例
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

## 彙總重點
{{< ask_chatgpt >}}
條列Java exception handling的重點
{{< /ask_chatgpt >}}



1. 知道哪些异常（Exceptions）需要處理（Handled Exceptions）；
2. 使用try-catch區塊來處理異常；
3. 瞭解異常的繼承關係（Exception Hierarchies）；
4. 了解發生異常會如何影響程式流程；
5. 使用finally區塊來保護資源不會被釋放掉；
6. 遵循Java API慣例，根據異常型別和程式邏輯，組織try-catch區塊。
7. 瞭解何時可以使用throws關鍵字來處理異常。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Java exception handling的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 問題：請寫一個方法，接收一個整數作為參數，若該參數小於0，則拋出一個IllegalArgumentException，否則返回該參數的平方。
答案：

public int getSquare(int num) {
    if (num < 0) { //檢查參數是否小於0
        throw new IllegalArgumentException("The input number should be greater than or equal to 0.");
    }
    return num * num;
}

2. 問題：請寫一個方法，接收一個字串作為參數，若該字串為空或null，則拋出一個NullPointerException，否則返回該字串的長度。
答案：

public int getLength(String str) {
    if (str == null || str.isEmpty()) { //檢查字串是否為空或null
        throw new NullPointerException("The input string cannot be null or empty.");
    }
    return str.length();
}

3. 問題：請寫一個方法，接收一個整數陣列和一個索引值作為參數，若索引值小於0或大於等於陣列長度，則拋出一個IndexOutOfBoundsException，否則返回該索引所對應的陣列元素。
答案：

public int getElement(int[] arr, int index) {
    if (index < 0 || index >= arr.length) { //檢查索引值是否越界
        throw new IndexOutOfBoundsException("The index is out of bounds.");
    }
    return arr[index];
}

4. 問題：請寫一個方法，接收一個字串和一個正整數作為參數，若該字串長度小於等於該正整數，則拋出一個RuntimeException，否則返回該字串的前n個字元。
答案：

public String getSubstring(String str, int n) {
    if (str.length() <= n) { //檢查字串長度是否足夠
        throw new RuntimeException("The length of the input string is not greater than " + n + ".");
    }
    return str.substring(0, n);
}

5. 問題：請寫一個方法，接收一個整數作為參數，若該整數為奇數，則拋出一個Checked Exception，否則返回該整數的平方。
答案：

public int getSquareEven(int num) throws Exception {
    if (num % 2 != 0) { //檢查整數是否為奇數
        throw new Exception("The input number should be even.");
    }
    return num * num;
}   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Java exception handling的網路資料
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

