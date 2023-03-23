+++
title = "Java物件导向"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Java OO的问题
{{< /ask_chatgpt >}}



1. 什麼是Java中的封裝(encapsulation)？請舉例說明。

2. 什麼是繼承(inheritance)？請舉例說明。

3. 什麼是多態(polymorphism)？請舉例說明。

4. 如何在Java中實現抽象類別(abstract class)？它與接口(interface)的區別是什麼？

5. 什麼是Java中的實例(instance)和類別(class)？它們之間有什麼區別？   

## 说明知识
{{< ask_chatgpt >}}
说明Java OO并举例
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

## 汇总重点
{{< ask_chatgpt >}}
条列Java OO的重点
{{< /ask_chatgpt >}}



以下是Java Object-Oriented (OO) 的重点：

1. 抽像化 (Abstraction) - 用来隐藏复杂性质，将其简化，以便于管理。

2. 封装 (Encapsulation) - 资料、方法的保护，防止私有资料不被乱改造，保护资源分配。

3. 继承 (Inheritance) - 继承能够从一个类别获得那些类别的资料和方法，使其成为一种新的类别。

4. 多型 (Polymorphism) - 多型概念代表能应用于不同形态之物件，也称为方法多载、方法覆载、物件多载等。

5. 类别 (Class) - Java 总是基于物件导向概念，具体表示就是定义一个 Class 并完成对 Class 的实例化。

6. 物件 (Object) - 由类别 Instance 到的物件，在程式执行时，被物件所组成。

7. 接口 (Interface) - 一种抽像的类型，Java 社群使用此类型来协调类。

8. 重载 (Overloading) - 在一个类别里，不同的方法可以有相同的函式名称，但参数类型不同或数目不同。

9. 重写 (Overriding) - 子类别重写父类别中具有相同名称和参数的方法。

10. 方法 (Method) - 方法是一种程式区块，可以接收参数，并能够产生一个或多个结果。在类别中，方法用来执行一些具体的功能。

11. 递移 (Recursion) - 递归是一种特殊的函数或方法，一个函数或方法可以在函数或方法内部呼叫自己。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Java OO的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 设计一个学生(Student)类别，包含以下属性和方法：

属性：
- 姓名(name)
- 年龄(age)
- 学号(studentID)
- 成绩(scores)

方法：
- 建构子(Constructor)
- 设置姓名(setName)
- 取得姓名(getName)
- 设置年龄(setAge)
- 取得年龄(getAge)
- 设置学号(setStudentID)
- 取得学号(getStudentID)
- 新增成绩(addScore)
- 取得成绩(getScore)
- 取得平均成绩(getAvgScore)

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

2. 设计一个矩形(Rectangle)类别，包含以下属性和方法：

属性：
- 宽度(width)
- 高度(height)

方法：
- 建构子(Constructor)
- 设置宽度(setWidth)
- 取得宽度(getWidth)
- 设置高度(setHeight)
- 取得高度(getHeight)
- 计算面积(getArea)

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

3. 设计一个彩票(Lottery)类别，包含以下属性和方法：

属性：
- 中奖号码(winningNumbers)
- 我的号码(myNumbers)

方法：
- 建构子(Constructor)
- 设置中奖号码(setWinningNumbers)
- 取得中奖号码(getWinningNumbers)
- 设置我的号码(setMyNumbers)
- 取得我的号码(getMyNumbers)
- 比较我的号码与中奖号码(match)

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

4. 设计一个银行帐户(BankAccount)类别，包含以下属性和方法：

属性：
- 帐户名称(accountName)
- 帐户余额(balance)

方法：
- 建构子(Constructor)
- 设置帐户名称(setAccountName)
- 取得帐户名称(getAccountName)
- 存款(deposit)
- 取款(withdraw)
- 取得帐户余额(getBalance)

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
      throw new IllegalArgumentException("余额不足");
    }
    this.balance -= amount;
  }

  public double getBalance() {
    return this.balance;
  }
}
```

5. 设计一个图形(Geometry)抽像类别，包含以下抽像方法：

方法：
- 计算面积(getArea)
- 计算周长(getPerimeter)

和以下实例方法：

方法：
- 取得图形名称(getName)

设计三个继承自Geometry类别的图形：三角形(Triangle)、矩形(Rectangle)和圆形(Circle)。

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

