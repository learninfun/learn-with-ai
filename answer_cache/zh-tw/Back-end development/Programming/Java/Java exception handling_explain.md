

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