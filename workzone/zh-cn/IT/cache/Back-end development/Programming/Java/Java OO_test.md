

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