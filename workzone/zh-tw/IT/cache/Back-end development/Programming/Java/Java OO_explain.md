

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