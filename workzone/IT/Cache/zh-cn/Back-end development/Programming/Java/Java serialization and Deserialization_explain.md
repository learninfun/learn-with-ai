

Java serialization 意味著將對像轉換為字節序列，以便可以存儲在文件中或通過網絡傳輸，並在需要時重新創建對象。在進行序列化時，對象的狀態信息被保存到文件或流中，例如Byte Stream或Disk, 這樣便可以在需要時通過反序列化將其轉換回原始狀態。Java序列化允許將任何對像保存到存儲器中，以便稍後反序列化時使用。Java提供了Serializable接口，所有類都通過實現此接口來啟用序列化和反序列化操作。

Java Deserialization 是將字節流轉換回對象的過程。在執行反序列化時，文件中的原始字節被讀取，並將其轉換成對應的Java對象。反序列化過程是反向處理對像序列化的過程，先從粘貼板中獲取序列化的內容，然後被反序列化成Java對象。

舉個例子，假設有一個名為Person的類，其中包含name, age和address屬性。如下例所示：

```
import java.io.*;

public class Person implements Serializable {
    private String name;
    private int age;
    private String address;

    public Person(String name, int age, String address) {
        this.name = name;
        this.age = age;
        this.address = address;
    }

    public String toString() {
        return "Name: " + name + ", Age: " + age + ", Address: " + address;
    }
}
```

現在可以將此類實例化並將其序列化為文件或字節中。假設有一個Person對象，將其寫入文件ser_file.ser中，如下所示：

```
import java.io.*;

public class SerializationDemo {
    public static void main(String[] args) {
        try {
            Person obj = new Person("John", 30, "New York");

            // Serialize the object
            FileOutputStream file = new FileOutputStream("ser_file.ser");
            ObjectOutputStream out = new ObjectOutputStream(file);
            out.writeObject(obj);
            out.close();
            file.close();

        } catch(IOException e) {
            e.printStackTrace();
        }
    }
}
```

現在，我們可以從ser_file.ser文件中讀取對象並將其反序列化回Person對象，如下所示：

```
import java.io.*;

public class DeserializationDemo {
    public static void main(String[] args) {
        try {
            // Deserialize the object
            FileInputStream file = new FileInputStream("ser_file.ser");
            ObjectInputStream in = new ObjectInputStream(file);
            Person obj = (Person)in.readObject();
            in.close();
            file.close();

            System.out.println(obj.toString());

        } catch(IOException e) {
            e.printStackTrace();
        } catch(ClassNotFoundException e) {
            e.printStackTrace();
        }
    }
}
```

此時，程序將輸出以下內容：

```
Name: John, Age: 30, Address: New York
```