

Java IO (Input/Output) is a pre-defined package in Java used for performing various input/output operations. It provides a set of classes to read and write data from/to various sources such as files, network connections, and memory buffers.

Example:
One of the most common example of Java I/O is reading and writing data to a file. This could be implemented as follows:

// Writing Data to a File:
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;

public class WriteToFile {
  public static void main(String[] args) {
    try {
      String content = "This is the data we're writing to a file.";
      File file = new File("example.txt");
      FileWriter writer = new FileWriter(file);
      writer.write(content);
      writer.close();
    } catch (IOException e) {
      System.out.println("An error occurred while writing to the file.");
      e.printStackTrace();
    }
  }
}

// Reading Data from a File:
import java.io.File;
import java.util.Scanner;

public class ReadFile {
  public static void main(String[] args) {
    try {
      File file = new File("example.txt");
      Scanner scanner = new Scanner(file);
      while (scanner.hasNextLine()) {
        String line = scanner.nextLine();
        System.out.println(line);
      }
      scanner.close();
    } catch (IOException e) {
      System.out.println("An error occurred while reading from the file.");
      e.printStackTrace();
    }
  }
}

In this example, we create a file object and use FileWriter to write data to a file. We then use Scanner to read the data from the file.