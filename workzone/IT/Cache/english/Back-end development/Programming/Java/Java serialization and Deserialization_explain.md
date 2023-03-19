

Java serialization is a process of converting an object's state into a byte stream, which can be stored into a file or transferred over a network. Java deserialization is the reverse process of converting the byte stream back into an object's state.

For example, suppose we have a class called "Employee" with fields like name, salary, and ID. We can serialize an instance of this class into a byte stream and save it to a file. Later we can deserialize it to recreate the Employee object.

To achieve serialization in Java, we need to implement the Serializable interface, which is a marker interface indicating that the class can be serialized. We also need to use an ObjectOutputStream to write the object to a file or network and ObjectInputStream to read the object from the file or network.

Here is an example code snippet showing how serialization and deserialization are done in Java:

```java
import java.io.*;

public class Employee implements Serializable {
   public String name;
   public String address;
   public transient int ssn;  // transient fields are not serialized
   
   public void setAddress(String address) {
      this.address = address;
   }

   public void setSSN(int ssn) {
      this.ssn = ssn;
   }

   public static void serialize(Employee emp, String fileName) throws IOException {
      FileOutputStream fileOut = new FileOutputStream(fileName);
      ObjectOutputStream outStream = new ObjectOutputStream(fileOut);
      outStream.writeObject(emp);  // write the Employee object to the file
      outStream.close();
      fileOut.close();
   }

   public static Employee deserialize(String fileName) throws IOException, ClassNotFoundException {
      FileInputStream fileIn = new FileInputStream(fileName);
      ObjectInputStream inStream = new ObjectInputStream(fileIn);
      Employee emp = (Employee) inStream.readObject();  // read the Employee object from the file
      inStream.close();
      fileIn.close();
      return emp;
   }
}

// Main program

public class Main {
   public static void main(String [] args) {
      Employee emp = new Employee();
      emp.name = "John Doe";
      emp.setAddress("123 Main St, San Francisco");
      emp.setSSN(123456789);

      try {
         Employee.serialize(emp, "employee.ser");
         Employee empDeserialized = Employee.deserialize("employee.ser");
      } catch (IOException e) {
         e.printStackTrace();
      } catch (ClassNotFoundException e) {
         e.printStackTrace();
      }
   }
}
```

In the example above, we created an Employee object, serialized it to a file called "employee.ser", and then deserialized it back to create a new Employee object called empDeserialized.