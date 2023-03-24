

Builder Pattern是一種軟體設計模式，用於在創建複雜的物件時簡化其構建過程。通過分解物件創建過程，它使得客戶端程式能夠更直觀地構建添加參數。

這個模式包含建造者，指導者和產品。建造者是負責創建複雜物件的介面，指導者負責使用建造者來創建物件，產品是建造完成的複雜物件。

以下是Builder Pattern的範例：

假設我們需要建立一個汽車物件，這個物件有各種屬性，如顏色，引擎類型，座位數量等。使用Builder Pattern，我們可以定義一個建造者介面`CarBuilder`，並實現不同的建造者類別，以構建不同的汽車屬性。

```
public interface CarBuilder {
    void setColor(String color);
    void setEngine(String engine);
    void setSeats(int seats);
    Car getResult();
}

public class SportsCarBuilder implements CarBuilder {
    private Car car;
    public SportsCarBuilder() {
        car = new Car();
    }
    public void setColor(String color) {
        car.setColor(color);
    }
    public void setEngine(String engine) {
        car.setEngine(engine);
    }
    public void setSeats(int seats) {
        car.setSeats(seats);
    }
    public Car getResult() {
        return car;
    }
}

public class SUVBuilder implements CarBuilder {
    private Car car;
    public SUVBuilder() {
        car = new Car();
    }
    public void setColor(String color) {
        car.setColor(color);
    }
    public void setEngine(String engine) {
        car.setEngine(engine);
    }
    public void setSeats(int seats) {
        car.setSeats(seats);
    }
    public Car getResult() {
        return car;
    }
}
```

建造者類別實現`CarBuilder`介面，該介面定義了用於設置汽車屬性的方法。每個建造者都有一個屬於自己的汽車物件實例，它最終將被返回為結果。 

最後，創建指導者類`Director`，它負責將汽車物件創建出來。

```
public class Director {
    private CarBuilder builder;
    public void setBuilder(CarBuilder builder) {
        this.builder = builder;
    }
    public Car getCar() {
        return builder.getResult();
    }
    public void constructCar(String color, String engine, int seats) {
        builder.setColor(color);
        builder.setEngine(engine);
        builder.setSeats(seats);
    }
}

public class Client {
    public static void main(String[] args) {
        Director director = new Director();
        CarBuilder sportsCarBuilder = new SportsCarBuilder();
        CarBuilder suvBuilder = new SUVBuilder();

        director.setBuilder(sportsCarBuilder);
        director.constructCar("Red", "V8", 2);
        Car sportsCar = director.getCar();

        director.setBuilder(suvBuilder);
        director.constructCar("Blue", "V6", 7);
        Car suv = director.getCar();
    }
}
```

客戶端使用指導者來創建汽車物件，使用構造方法傳遞汽車屬性，指導者則使用建造者創建完成的汽車物件。這樣，客戶端不需要知道汽車物件是如何創建的，而是通過使用建造者和指導者來隔離物件創建過程，使得程式碼更加清晰、簡潔。