+++
title = "Java異常處理"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Java exception handling的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Java exception handling？如何利用try-catch block處理異常？

2. 請舉例說明什麼是unchecked exception和checked exception？如何處理它們？

3. 請列舉幾種Java異常處理中的關鍵字及其作用？

4. 如何設計自己的Java異常類別？請使用代碼示例說明。

5. 請比較throw和throws在Java異常處理中的區別，並舉例說明如何使用它們。   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Java exception handling並舉例
{{< /ask_chatgpt >}}



Java exception handling是一種解決代碼運行時遇到異常情況的機制。Java的exception handling機制允許程序在運行時處理錯誤，避免異常使程序崩潰並減少對代碼的影響。

Java異常處理語句的結構是：

```java
try {
      //代碼塊
} catch (exceptionType1 e1) {
      //異常處理語句
} catch (exceptionType2 e2) {
      //異常處理語句
} catch (exceptionType3 e3) {
      //異常處理語句
} finally {
      //可選代碼塊
}
```

其中，try塊包含可能會造成異常的代碼，catch塊用於處理異常類型，finally塊可選，包含在執行完try和catch塊之後始終執行的代碼。

以下是一個Java異常處理的示例，其中將嘗試讀取一個不存在的文件，捕獲FileNotFoundException並輸出錯誤信息：

```java
import java.io.*;

class ExceptionExample {

     public static void main(String[] args) {

          try {
               // 打開文件
               FileInputStream file = new FileInputStream("example.txt");
          } catch (FileNotFoundException e) {
               System.out.println("找不到文件");
               e.printStackTrace();
          } 

     }
}
```

在上述代碼中，當程序嘗試讀取一個不存在的文件時，拋出FileNotFoundException異常。try塊的打開文件代碼可能會拋出異常，所以我們將其置於try塊中。如果FileNotFoundException異常被拋出，則catch塊將被執行。在catch塊中，我們輸出一個錯誤信息並使用e.printStackTrace()方法打印異常的調用棧信息，以幫助我們更好地理解錯誤原因。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Java exception handling的重點
{{< /ask_chatgpt >}}



1. 知道哪些異常（Exceptions）需要處理（Handled Exceptions）；
2. 使用try-catch區塊來處理異常；
3. 瞭解異常的繼承關係（Exception Hierarchies）；
4. 瞭解發生異常會如何影響程式流程；
5. 使用finally區塊來保護資源不會被釋放掉；
6. 遵循Java API慣例，根據異常型別和程式邏輯，組織try-catch區塊。
7. 瞭解何時可以使用throws關鍵字來處理異常。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Java exception handling的中等難度問題，並在後面列出答案
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

