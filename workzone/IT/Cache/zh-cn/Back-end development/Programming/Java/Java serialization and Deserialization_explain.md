

Java serialization 意味着将对象转换为字节序列，以便可以存储在文件中或通过网络传输，并在需要时重新创建对象。在进行序列化时，对象的状态信息被保存到文件或流中，例如Byte Stream或Disk, 这样便可以在需要时通过反序列化将其转换回原始状态。Java序列化允许将任何对象保存到存储器中，以便稍后反序列化时使用。Java提供了Serializable接口，所有类都通过实现此接口来启用序列化和反序列化操作。

Java Deserialization 是将字节流转换回对象的过程。在执行反序列化时，文件中的原始字节被读取，并将其转换成对应的Java对象。反序列化过程是反向处理对象序列化的过程，先从粘贴板中获取序列化的内容，然后被反序列化成Java对象。

举个例子，假设有一个名为Person的类，其中包含name, age和address属性。如下例所示：

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

现在可以将此类实例化并将其序列化为文件或字节中。假设有一个Person对象，将其写入文件ser_file.ser中，如下所示：

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

现在，我们可以从ser_file.ser文件中读取对象并将其反序列化回Person对象，如下所示：

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

此时，程序将输出以下内容：

```
Name: John, Age: 30, Address: New York
```