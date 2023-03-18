+++
title = "Java输入输出"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Java IO的问题
{{< /ask_chatgpt >}}



1. 什麼是Java中的基本輸入輸出？ 
2. 如何使用Java IO將數據從文件讀取到內存中？ 
3. 如何使用Java IO將數據從內存中寫入文件？ 
4. 什麼是Java中的緩衝輸入輸出？ 
5. 如何在Java中使用字符流進行輸入輸出？   

## 说明知识
{{< ask_chatgpt >}}
说明Java IO并举例
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

## 汇总重点
{{< ask_chatgpt >}}
条列Java IO的重点
{{< /ask_chatgpt >}}



1. Input/Output Stream: Java提供了两种Stream来处理IO，一种是Input Stream用于从外部资源读取数据，另一种是Output Stream用于将数据写入外部资源。

2. Reader/Writer: 除了Stream外，Java还提供了Reader和Writer来处理文字资料的IO操作。Reader用于从外部资源读取文字数据，而Writer用于将文字数据写入外部资源。

3. File IO: Java提供了File类来处理档案IO，可以使用File类来创建、写入、读取和删除档案。

4. 字符集(Encoding): 在处理文字资料IO时，需要注意字符集的问题，Java提供了多种字符集，可以使用相应的字符集来处理不同语系的文字资料。

5. 序列化(Serialization): Java提供了序列化技术，可以将物件转换成byte数组，以便在不同的环境中进行传输和存储。

6. 缓冲区(Buffering): 为了提高IO操作的效率，Java提供了缓冲区(Buffer)，通过将IO操作的数据先缓存在缓冲区中，在一定条件下再将数据批量写入外部资源，可以提高IO操作的效率。

7. NIO(Non-blocking IO): Java提供了NIO技术，可以实现非阻塞的IO操作，增强系统的扩展性和吞吐量。NIO技术包括通道、缓冲区和选择器等组件。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Java IO的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 请撰写一个程式，从档案中读取不定数量的整数，并计算它们的总和。

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

2. 请撰写一个程式，将档案中的字符串逐行读取并逐行写入新的档案中。

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

3. 请撰写一个程式，将二进制档案中的数据按照特定格式进行读取和输出。

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

4. 请撰写一个程式，从档案中读取包含多个单字的字符串，并将每个单字的首字母转换为大写。

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

5. 请撰写一个程式，将多个档案中的数据合并到一个新档案中。

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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Java IO的网络数据
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

