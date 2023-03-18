

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