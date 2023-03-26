

Builder Pattern是一种软体设计模式，用于在创建复杂的物件时简化其构建过程。通过分解物件创建过程，它使得客户端程式能够更直观地构建添加参数。

这个模式包含建造者，指导者和产品。建造者是负责创建复杂物件的介面，指导者负责使用建造者来创建物件，产品是建造完成的复杂物件。

以下是Builder Pattern的范例：

假设我们需要建立一个汽车物件，这个物件有各种属性，如颜色，引擎类型，座位数量等。使用Builder Pattern，我们可以定义一个建造者介面`CarBuilder`，并实现不同的建造者类别，以构建不同的汽车属性。

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

建造者类别实现`CarBuilder`介面，该介面定义了用于设置汽车属性的方法。每个建造者都有一个属于自己的汽车物件实例，它最终将被返回为结果。 

最后，创建指导者类`Director`，它负责将汽车物件创建出来。

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

客户端使用指导者来创建汽车物件，使用构造方法传递汽车属性，指导者则使用建造者创建完成的汽车物件。这样，客户端不需要知道汽车物件是如何创建的，而是通过使用建造者和指导者来隔离物件创建过程，使得程式码更加清晰、简洁。