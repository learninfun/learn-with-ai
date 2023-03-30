1. What is the difference between InputStream and OutputStream in Java?
- InputStream is used to read data from a source, while OutputStream is used to write data to a destination.

2. What is the purpose of FileReader and FileWriter in Java?
- FileReader is used to read character streams from a file, while FileWriter is used to write character streams to a file.

3. How do you read a file line by line in Java?
- You can use a BufferedReader to read the file line by line, using the readLine() method.

4. What is a BufferedOutputStream in Java, and why is it useful?
- BufferedOutputStream is a class that is used to improve the performance of output streams by buffering the data before it is written to the destination. This reduces the number of I/O operations required, which can improve performance when working with large amounts of data.

5. What is the purpose of the ObjectInputStream and ObjectOutputStream classes in Java?
- These classes are used to read and write objects from/to a stream. They are often used in conjunction with socket communication or file I/O, and allow objects to be serialized/deserialized in a platform-independent manner.