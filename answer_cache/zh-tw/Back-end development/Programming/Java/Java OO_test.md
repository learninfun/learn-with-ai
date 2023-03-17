

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

5. 設計一個圖形(Geometry)抽像類別，包含以下抽像方法：

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