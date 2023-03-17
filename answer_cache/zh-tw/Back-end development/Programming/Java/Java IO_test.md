

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