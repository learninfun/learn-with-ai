

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