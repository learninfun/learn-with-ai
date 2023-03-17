## 習題預習
{{< ask_chatgpt >}}
給我5題Java reflection的問題
{{< /ask_chatgpt >}}



1. 什麼是Java反射？它如何提升Java的靈活性和可重用性？

2. 在Java中，如何使用反射技術來獲取類型的信息？並且如何利用反射来運行方法和訪問屬性？

3. 反射技術是如何在Java中實現的？它是如何同其他Java特性（如反序列化、代理、AOP）交互的？

4. 反射技術可能會導致性能下降，你有什麼其他選擇可以實現動態性？

5. 反射的應用有哪些？它在框架設計和開發中的作用是什麼？   

## 說明知識
{{< ask_chatgpt >}}
說明Java reflection並舉例
{{< /ask_chatgpt >}}



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

## 彙總重點
{{< ask_chatgpt >}}
條列Java reflection的重點
{{< /ask_chatgpt >}}



1. Java反射提供了动态操作对象的能力。使用反射API，可以在运行时获取类和对象的信息并操作它们，即使在编译时无法确定对象的类型也可以实现这个功能。

2. Java反射包含了三个核心类Class、Method和Field，分别用来获取类、方法和属性的信息。

3. Java反射可以用来创建新的对象、调用方法或修改属性值等，还可以获取类或对象的注解信息、修饰符、泛型信息等。

4. Java反射可以被用来实现一些高级功能，比如动态代理、依赖注入、ORM框架等。

5. Java反射应该谨慎使用，因为它会引入一些性能问题，如遍历对象树和解析注解等都比静态代码更慢。此外，反射还会增加代码的复杂性和不可预测性。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Java reflection的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 如何使用Java reflection取得某個class的所有方法名稱？

```
Class<?> clazz = MyClass.class;
Method[] methods = clazz.getDeclaredMethods();
for (Method method : methods) {
    System.out.println(method.getName());
}
```

2. 如何使用Java reflection取得某個class的所有變數名稱？

```
Class<?> clazz = MyClass.class;
Field[] fields = clazz.getDeclaredFields();
for (Field field : fields) {
    System.out.println(field.getName());
}
```

3. 如何使用Java reflection動態新增一個物件實例？

```
Class<?> clazz = MyClass.class;
Object newInstance = clazz.getDeclaredConstructor().newInstance();
```

4. 如何使用Java reflection設定某個物件實例的屬性值？

```
MyClass obj = new MyClass();
Class<?> clazz = MyClass.class;
Field field = clazz.getDeclaredField("fieldName");
field.setAccessible(true);
field.set(obj, "newValue");
```

5. 如何使用Java reflection呼叫某個物件實例的方法？

```
MyClass obj = new MyClass();
Class<?> clazz = MyClass.class;
Method method = clazz.getDeclaredMethod("methodName", String.class);
method.invoke(obj, "parameterValue");
```

答案僅供參考，實際實現方式可能有所不同。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Java reflection的網路資料
{{< /ask_chatgpt >}}



1. Java Reflection Tutorial: How to Use Reflection? 

https://www.baeldung.com/java-reflection 

This tutorial is a comprehensive guide to Java Reflection. The article covers what is Java Reflection, how to use Reflection to create and manipulate objects, how to access and modify fields and methods of an object at runtime, and how to use annotations with Reflection. The article also includes many examples and code snippets to help readers understand the concepts better. 

2. A Guide to Java Reflection 

https://www.javacodegeeks.com/guide-java-reflection/ 

This guide provides a deep dive into Java Reflection. It covers the essential concepts of Reflection, such as Classes and Objects, their fields and methods, and how to get and set their values at runtime. The article also covers the security implications of Reflection and how to use it carefully. There are many examples and practical use cases throughout the article. 

3. Java Reflection: A Practical tutorial 

https://www.oracle.com/technical-resources/articles/java/javareflection.html 

The Oracle Java Reflection tutorial is an excellent practical guide to using Reflection in Java applications. It covers the basics of Reflection, such as creating instances, invoking methods, and accessing fields. It also covers advanced topics such as Reflection on Generics, Reflection with Annotations, and Reflection on Enums. The article also provides practical examples and exercises that readers can follow to learn Java Reflection. 

4. Java Reflection: A Beginner's Guide 

https://www.tutorialspoint.com/java_reflection/ 

This beginner's guide is an excellent place to start learning Java Reflection. It covers the basics of Reflection and how to use Reflection to create and manipulate objects at runtime. The article also covers how to get information about a class and its methods and fields, how to access and modify them, and how to use Reflection to enhance your Java applications. The article includes many practical examples to help readers learn the concepts faster. 

5. Java Reflection – Tutorial with Examples 

https://www.journaldev.com/1789/java-reflection-example-tutorial 

This tutorial from journaldev.com is a step-by-step guide to Java Reflection with practical examples. The article covers the concepts of Reflection, such as Classes, Fields, and Methods, and how to use them in your Java applications to create and manipulate objects at runtime. The article also covers advanced topics, such as Reflection on Annotations, Reflection on Enums, and Reflection on Generics. The article also provides many examples and code snippets that readers can use to learn the concepts faster.   

