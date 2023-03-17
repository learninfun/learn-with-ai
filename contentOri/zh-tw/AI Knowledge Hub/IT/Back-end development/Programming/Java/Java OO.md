## 習題預習
{{< ask_chatgpt >}}
給我5題Java OO的問題
{{< /ask_chatgpt >}}



1. 什麼是Java中的封裝(encapsulation)？請舉例說明。

2. 什麼是繼承(inheritance)？請舉例說明。

3. 什麼是多態(polymorphism)？請舉例說明。

4. 如何在Java中實現抽象類別(abstract class)？它與接口(interface)的區別是什麼？

5. 什麼是Java中的實例(instance)和類別(class)？它們之間有什麼區別？   

## 說明知識
{{< ask_chatgpt >}}
說明Java OO並舉例
{{< /ask_chatgpt >}}



Java是一種以物件為中心的程式設計語言，也就是說它是一種OOP（Object Oriented Programming）程式設計語言。

OOP是一種程序開發方法，專注於組織程式可以訪問的數據。在OOP中，程序被組織為一系列的對象（Object），每個對象都有其屬性和方法。

以下是一個Java OOP的簡單示例：

public class Car {
  private String make;
  private String model;
  private int year;

  public Car(String make, String model, int year) {
    this.make = make;
    this.model = model;
    this.year = year;
  }

  public String getMake() {
    return make;
  }

  public String getModel() {
    return model;
  }

  public int getYear() {
    return year;
  }

  public void setMake(String make) {
    this.make = make;
  }

  public void setModel(String model) {
    this.model = model;
  }

  public void setYear(int year) {
    this.year = year;
  }
}

在這個例子中，Car類就是一個對象，該對象具有3個屬性：make（製造商），model（型號）和year（年份）。該類還有一個public的構造函數（Constructor），它接受make，model和year作為參數，並設置對象的屬性。還有一些公共的getter和setter方法來訪問和設置對象的屬性。

使用這個Car類，可以創建具有不同make，model和year的多個對象。

Car car1 = new Car("Honda", "Civic", 2021);
Car car2 = new Car("Toyota", "Camry", 2019);

這裡，car1和car2都是Car對象，並且可以調用它們的方法來訪問和設置它們的屬性。

String makeOfCar1 = car1.getMake(); // Returns "Honda"
int yearOfCar2 = car2.getYear(); // Returns 2019

總之， Java OOP 就是使用物件（Object）作為程式的基本單位，專注於組織程式可以訪問的數據和方法，使得程式開發更加模組化和易於維護。   

## 彙總重點
{{< ask_chatgpt >}}
條列Java OO的重點
{{< /ask_chatgpt >}}



以下是Java Object-Oriented (OO) 的重點：

1. 抽象化 (Abstraction) - 用來隱藏複雜性質，將其簡化，以便於管理。

2. 封裝 (Encapsulation) - 資料、方法的保護，防止私有資料不被亂改造，保護資源分配。

3. 繼承 (Inheritance) - 繼承能夠從一個類別獲得那些類別的資料和方法，使其成為一種新的類別。

4. 多型 (Polymorphism) - 多型概念代表能應用於不同形態之物件，也稱為方法多載、方法覆載、物件多載等。

5. 類別 (Class) - Java 總是基於物件導向概念，具體表示就是定義一個 Class 並完成對 Class 的實例化。

6. 物件 (Object) - 由類別 Instance 到的物件，在程式執行時，被物件所組成。

7. 接口 (Interface) - 一種抽象的類型，Java 社群使用此類型來協調類。

8. 重載 (Overloading) - 在一個類別裡，不同的方法可以有相同的函式名稱，但參數類型不同或數目不同。

9. 重寫 (Overriding) - 子類別重寫父類別中具有相同名稱和參數的方法。

10. 方法 (Method) - 方法是一種程式區塊，可以接收參數，並能夠產生一個或多個結果。在類別中，方法用來執行一些具體的功能。

11. 遞移 (Recursion) - 遞歸是一種特殊的函數或方法，一個函數或方法可以在函數或方法內部呼叫自己。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Java OO的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 設計一個學生(Student)類別，包含以下屬性和方法：

屬性：
- 姓名(name)
- 年齡(age)
- 學號(studentID)
- 成績(scores)

方法：
- 建構子(Constructor)
- 設置姓名(setName)
- 取得姓名(getName)
- 設置年齡(setAge)
- 取得年齡(getAge)
- 設置學號(setStudentID)
- 取得學號(getStudentID)
- 新增成績(addScore)
- 取得成績(getScore)
- 取得平均成績(getAvgScore)

答案：

```
public class Student {
  private String name;
  private int age;
  private String studentID;
  private ArrayList<Integer> scores = new ArrayList<Integer>();

  public Student(String name, int age, String studentID) {
    this.name = name;
    this.age = age;
    this.studentID = studentID;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getName() {
    return this.name;
  }

  public void setAge(int age) {
    this.age = age;
  }

  public int getAge() {
    return this.age;
  }

  public void setStudentID(String studentID) {
    this.studentID = studentID;
  }

  public String getStudentID() {
    return this.studentID;
  }

  public void addScore(int score) {
    this.scores.add(score);
  }

  public int getScore(int index) {
    return this.scores.get(index);
  }

  public double getAvgScore() {
    double total = 0;
    for (int score : this.scores) {
      total += score;
    }
    return total / this.scores.size();
  }
}
```

2. 設計一個矩形(Rectangle)類別，包含以下屬性和方法：

屬性：
- 寬度(width)
- 高度(height)

方法：
- 建構子(Constructor)
- 設置寬度(setWidth)
- 取得寬度(getWidth)
- 設置高度(setHeight)
- 取得高度(getHeight)
- 計算面積(getArea)

答案：

```
public class Rectangle {
  private int width;
  private int height;

  public Rectangle(int width, int height) {
    this.width = width;
    this.height = height;
  }

  public void setWidth(int width) {
    this.width = width;
  }

  public int getWidth() {
    return this.width;
  }

  public void setHeight(int height) {
    this.height = height;
  }

  public int getHeight() {
    return this.height;
  }

  public int getArea() {
    return this.width * this.height;
  }
}
```

3. 設計一個彩票(Lottery)類別，包含以下屬性和方法：

屬性：
- 中獎號碼(winningNumbers)
- 我的號碼(myNumbers)

方法：
- 建構子(Constructor)
- 設置中獎號碼(setWinningNumbers)
- 取得中獎號碼(getWinningNumbers)
- 設置我的號碼(setMyNumbers)
- 取得我的號碼(getMyNumbers)
- 比較我的號碼與中獎號碼(match)

答案：

```
public class Lottery {
  private int[] winningNumbers;
  private int[] myNumbers;

  public Lottery(int[] winningNumbers) {
    this.winningNumbers = winningNumbers;
  }

  public void setWinningNumbers(int[] winningNumbers) {
    this.winningNumbers = winningNumbers;
  }

  public int[] getWinningNumbers() {
    return this.winningNumbers;
  }

  public void setMyNumbers(int[] myNumbers) {
    this.myNumbers = myNumbers;
  }

  public int[] getMyNumbers() {
    return this.myNumbers;
  }

  public int match() {
    int count = 0;
    for (int i = 0; i < this.myNumbers.length; i++) {
      for (int j = 0; j < this.winningNumbers.length; j++) {
        if (this.myNumbers[i] == this.winningNumbers[j]) {
          count++;
        }
      }
    }
    return count;
  }
}
```

4. 設計一個銀行帳戶(BankAccount)類別，包含以下屬性和方法：

屬性：
- 帳戶名稱(accountName)
- 帳戶餘額(balance)

方法：
- 建構子(Constructor)
- 設置帳戶名稱(setAccountName)
- 取得帳戶名稱(getAccountName)
- 存款(deposit)
- 取款(withdraw)
- 取得帳戶餘額(getBalance)

答案：

```
public class BankAccount {
  private String accountName;
  private double balance;

  public BankAccount(String accountName, double balance) {
    this.accountName = accountName;
    this.balance = balance;
  }

  public void setAccountName(String accountName) {
    this.accountName = accountName;
  }

  public String getAccountName() {
    return this.accountName;
  }

  public void deposit(double amount) {
    this.balance += amount;
  }

  public void withdraw(double amount) {
    if (amount > this.balance) {
      throw new IllegalArgumentException("餘額不足");
    }
    this.balance -= amount;
  }

  public double getBalance() {
    return this.balance;
  }
}
```

5. 設計一個圖形(Geometry)抽象類別，包含以下抽象方法：

方法：
- 計算面積(getArea)
- 計算周長(getPerimeter)

和以下實例方法：

方法：
- 取得圖形名稱(getName)

設計三個繼承自Geometry類別的圖形：三角形(Triangle)、矩形(Rectangle)和圓形(Circle)。

答案：

```
public abstract class Geometry {
  public abstract double getArea();
  public abstract double getPerimeter();
  public String getName() {
    return this.getClass().getSimpleName();
  }
}

public class Triangle extends Geometry {
  private double base;
  private double height;
  private double side1;
  private double side2;
  
  public Triangle(double base, double height, double side1, double side2) {
    this.base = base;
    this.height = height;
    this.side1 = side1;
    this.side2 = side2;
  }
  
  public double getArea() {
    return this.base * this.height / 2;
  }
  
  public double getPerimeter() {
    return this.base + this.side1 + this.side2;
  }
}

public class Rectangle extends Geometry {
  private double width;
  private double height;
  
  public Rectangle(double width, double height) {
    this.width = width;
    this.height = height;
  }
  
  public double getArea() {
    return this.width * this.height;
  }
  
  public double getPerimeter() {
    return 2 * (this.width + this.height);
  }
}

public class Circle extends Geometry {
  private double radius;
  
  public Circle(double radius) {
    this.radius = radius;
  }
  
  public double getArea() {
    return Math.PI * this.radius * this.radius;
  }
  
  public double getPerimeter() {
    return 2 * Math.PI * this.radius;
  }
}
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Java OO的網路資料
{{< /ask_chatgpt >}}



1. Java OOP Tutorial: Object-Oriented Programming For Beginners
https://www.journaldev.com/1663/java-oops-concepts-with-examples

This tutorial by JournalDev covers the fundamentals of Object-Oriented Programming (OOP) in Java, including encapsulation, inheritance, and polymorphism. It includes code snippets and examples to aid in understanding.

2. Java OOP Concepts Explained: Inheritance
https://stackify.com/oop-concepts-inheritance/

This article by Stackify delves deeper into inheritance, one of the key concepts in Java OOP. It discusses different types of inheritance and gives examples of each in use. 

3. Java OOP Concepts Explained: Polymorphism
https://stackify.com/oop-concepts-polymorphism/

This is another article by Stackify that covers polymorphism in Java OOP. It explains what polymorphism is and how it can be used to make code more efficient and organized. 

4. Java OOP Concepts Explained: Abstraction
https://stackify.com/oop-concepts-abstraction/

Abstraction is a complex but important concept in Java OOP. This article by Stackify breaks down what it is and how it works. It includes examples to help readers understand how to implement abstraction in code. 

5. Java Object-Oriented Programming: Understanding Classes
https://www.c-sharpcorner.com/article/java-object-oriented-programming-understanding-classes/

This tutorial by C# Corner explains how to create classes in Java, including details on constructors and access modifiers. Additionally, it explains how to utilize these classes within other Java projects.   

