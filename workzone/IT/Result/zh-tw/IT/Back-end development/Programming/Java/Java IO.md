+++
title = "Java輸入輸出"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Java IO的問題
{{< /ask_chatgpt >}}



1. 什麼是Java中的基本輸入輸出？ 
2. 如何使用Java IO將數據從文件讀取到內存中？ 
3. 如何使用Java IO將數據從內存中寫入文件？ 
4. 什麼是Java中的緩衝輸入輸出？ 
5. 如何在Java中使用字符流進行輸入輸出？   

## 說明知識
{{< ask_chatgpt >}}
說明Java IO並舉例
{{< /ask_chatgpt >}}



Java IO（Input/Output）是指Java語言使用的標準輸入輸出流程。Java IO主要用於讀取和寫入文件以及與其他設備交互，例如網絡設備。Java IO的主要目的是讓Java應用程序能夠讀寫數據，包括字符（文本）和字節（二進制）數據。

Java IO的主要類包括：

- InputStream和OutputStream（字節流）
- Reader和Writer（字符流）

Java IO還有其他類，例如File、RandomAccessFile、BufferedReader和BufferedWriter等，這些類可以方便地讀取和寫入文件。

下面是一個簡單的Java IO示例，該示例演示了如何從文件中讀取數據並將其輸出到控制台：

```
import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class ReadFromFile {

    public static void main(String[] args) {

        try (BufferedReader br = new BufferedReader(new FileReader("input.txt"))) {

            String line;
            while ((line = br.readLine()) != null) {
                System.out.println(line);
            }

        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```

在這個例子中，我們使用BufferedReader來讀取一個文件的內容。我們首先創建一個FileReader對象來讀取文件（使用try-with-resources語法，因此不需要手動關閉文件）。我們然後使用BufferedReader來緩存輸入，以提高效率。最後，我們使用while迴圈逐行讀取文件，並將讀取到的每行輸出到控制台。

這僅僅是Java IO的一個簡單示例，Java IO還有很多強大的功能和技術可以探索。   

## 彙總重點
{{< ask_chatgpt >}}
條列Java IO的重點
{{< /ask_chatgpt >}}



1. Input/Output Stream: Java提供了兩種Stream來處理IO，一種是Input Stream用於從外部資源讀取數據，另一種是Output Stream用於將數據寫入外部資源。

2. Reader/Writer: 除了Stream外，Java還提供了Reader和Writer來處理文字資料的IO操作。Reader用於從外部資源讀取文字數據，而Writer用於將文字數據寫入外部資源。

3. File IO: Java提供了File類來處理檔案IO，可以使用File類來創建、寫入、讀取和刪除檔案。

4. 字符集(Encoding): 在處理文字資料IO時，需要注意字符集的問題，Java提供了多種字符集，可以使用相應的字符集來處理不同語系的文字資料。

5. 序列化(Serialization): Java提供了序列化技術，可以將物件轉換成byte數組，以便在不同的環境中進行傳輸和存儲。

6. 緩衝區(Buffering): 為了提高IO操作的效率，Java提供了緩衝區(Buffer)，通過將IO操作的數據先緩存在緩衝區中，在一定條件下再將數據批量寫入外部資源，可以提高IO操作的效率。

7. NIO(Non-blocking IO): Java提供了NIO技術，可以實現非阻塞的IO操作，增強系統的擴展性和吞吐量。NIO技術包括通道、緩衝區和選擇器等組件。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Java IO的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 請撰寫一個程式，從檔案中讀取不定數量的整數，並計算它們的總和。

```
import java.io.*;
import java.util.Scanner;

public class SumIntegersFromFile {
    public static void main(String[] args) {
        File inputFile = new File("input.txt");
        try {
            Scanner scanner = new Scanner(inputFile);
            int sum = 0;
            while (scanner.hasNextInt()) {
                sum += scanner.nextInt();
            }
            System.out.println("The sum of integers is: " + sum);
            scanner.close();
        } catch (IOException e) {
            System.out.println("An error occurred: " + e);
        }
    }
}
```

2. 請撰寫一個程式，將檔案中的字符串逐行讀取並逐行寫入新的檔案中。

```
import java.io.*;

public class CopyFileLineByLine {
    public static void main(String[] args) {
        File inputFile = new File("input.txt");
        File outputFile = new File("output.txt");
        try {
            BufferedReader reader = new BufferedReader(new FileReader(inputFile));
            BufferedWriter writer = new BufferedWriter(new FileWriter(outputFile));
            String line;
            while ((line = reader.readLine()) != null) {
                writer.write(line);
                writer.newLine();
            }
            reader.close();
            writer.close();
        } catch (IOException e) {
            System.out.println("An error occurred: " + e);
        }
    }
}
```

3. 請撰寫一個程式，將二進制檔案中的數據按照特定格式進行讀取和輸出。

```
import java.io.*;

public class BinaryFileFormat {
    public static void main(String[] args) {
        File inputFile = new File("input.bin");
        File outputFile = new File("output.bin");
        try {
            DataInputStream inputStream = new DataInputStream(new FileInputStream(inputFile));
            DataOutputStream outputStream = new DataOutputStream(new FileOutputStream(outputFile));
            int num1 = inputStream.readInt();
            double num2 = inputStream.readDouble();
            boolean num3 = inputStream.readBoolean();
            outputStream.writeUTF("Number 1: " + num1 + ", Number 2: " + num2 + ", Number 3: " + num3);
            inputStream.close();
            outputStream.close();
        } catch (IOException e) {
            System.out.println("An error occurred: " + e);
        }
    }
}
```

4. 請撰寫一個程式，從檔案中讀取包含多個單字的字符串，並將每個單字的首字母轉換為大寫。

```
import java.io.*;
import java.util.Scanner;

public class CapitalizeFirstLetter {
    public static void main(String[] args) {
        File inputFile = new File("input.txt");
        File outputFile = new File("output.txt");
        try {
            Scanner scanner = new Scanner(inputFile);
            BufferedWriter writer = new BufferedWriter(new FileWriter(outputFile));
            while (scanner.hasNext()) {
                String word = scanner.next();
                writer.write(Character.toUpperCase(word.charAt(0)) + word.substring(1));
                writer.write(" ");
            }
            scanner.close();
            writer.close();
        } catch (IOException e) {
            System.out.println("An error occurred: " + e);
        }
    }
}
```

5. 請撰寫一個程式，將多個檔案中的數據合併到一個新檔案中。

```
import java.io.*;

public class MergeFiles {
    public static void main(String[] args) {
        File[] inputFiles = { new File("input1.txt"), new File("input2.txt"), new File("input3.txt")};
        File outputFile = new File("output.txt");
        try {
            BufferedWriter writer = new BufferedWriter(new FileWriter(outputFile));
            for (File inputFile : inputFiles) {
                BufferedReader reader = new BufferedReader(new FileReader(inputFile));
                String line;
                while ((line = reader.readLine()) != null) {
                    writer.write(line);
                    writer.newLine();
                }
                reader.close();
            }
            writer.close();
        } catch (IOException e) {
            System.out.println("An error occurred: " + e);
        }
    }
}
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Java IO的網路資料
{{< /ask_chatgpt >}}



以下是5篇有關Java IO的網路資料，供您參考：

1. Java IO Tutorial: https://www.tutorialspoint.com/java/java_io.htm

這個教學說明了Java IO的概念和使用方法，包括文件輸入/輸出、字符輸入/輸出等基礎內容。

2. Java IO – Input Output in Java with Examples: https://www.guru99.com/java-io.html

此文章介紹了Java IO的基本操作，如範例中的文件輸入/輸出、字節流、字符流等。

3. Java IO Tutorial – Java IO (Input/Output) Classes: https://www.baeldung.com/java-io

此篇文章詳細介紹了Java IO的各種類別，諸如文件輸入/輸出、字符輸入/輸出、字節流、轉換流等等。

4. Java NIO Tutorial: http://tutorials.jenkov.com/java-nio/index.html

NIO即New Input/Output，是Java IO的一個新特性，具有非阻塞和多路複用的特點，該文章詳細介紹如何使用Java NIO。

5. Java IO and NIO – A Quick Guide: https://www.careerbless.com/java/io/io_nio.php

這篇文章簡要介紹了Java IO的基本概念和應用，以及Java NIO的優點和缺點，是一篇簡易入門的文章。   

