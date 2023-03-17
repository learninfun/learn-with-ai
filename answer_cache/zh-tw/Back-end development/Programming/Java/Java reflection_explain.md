

Java reflection是一種技術，它允許Java程序在運行時獲取對象的屬性、方法和類型信息。它可以讓程序動態地調用對象的方法，以及查看和修改對象的屬性，而不需要在編譯時確定這些信息。

以下是Java reflection的一個例子：

假設有一個類別Person：

```java
public class Person {
    private String name;
    private int age;

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }

    public void sayHello() {
        System.out.println("Hello, my name is " + name + " and I am " + age + " years old.");
    }
}
```

我們可以通過下面的代碼在運行時創建一個Person對象，獲取對象的名字和年齡屬性，並動態地調用sayHello方法：

```java
// 創建一個Person對象
Person person = new Person("Tom", 18);

// 獲取對象的名字和年齡屬性
Class<?> personClass = person.getClass();
Field nameField = personClass.getDeclaredField("name");
Field ageField = personClass.getDeclaredField("age");
nameField.setAccessible(true);
ageField.setAccessible(true);
String name = (String) nameField.get(person);
int age = (int) ageField.get(person);

// 動態地調用sayHello方法
Method sayHelloMethod = personClass.getDeclaredMethod("sayHello");
sayHelloMethod.invoke(person);
```

在上面的代碼中，我們通過person.getClass()方法獲取Person類別對象，然後通過getDeclaredField獲取對象的屬性，使用setAccessible方法設置屬性可訪問性，使用get方法獲取屬性值。最後，我們使用getDeclaredMethod方法獲取對象的方法，使用invoke方法動態地調用sayHello方法。

使用Java reflection，我們可以在運行時動態地操作對象，並具有很高的靈活性。但是，由於Java reflection是一個動態的過程，它的效率比編譯時的代碼更低，因此應該避免在性能敏感的代碼中過度使用它。