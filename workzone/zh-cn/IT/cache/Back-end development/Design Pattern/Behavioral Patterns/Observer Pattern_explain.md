

Observer Pattern是一种经典的行为型设计模式，它定义了对象间的一种一对多（one-to-many）的依赖关系，以便当一个对象的状态发生改变时，其所有依赖对象都能够自动收到通知并更新状态。

该模式中包含三种角色：Subject（目标对像）、Observer（观察者对像）和ConcreteObserver（具体的观察者对像）。

Subject是被观察的对象，它维护一组Observer对象，当它的状态发生变化时，会遍历这些Observer对象，并调用它们的update()方法通知它们。

Observer是观察者对象，它定义了一个接口，用于接收关于Subject状态改变的通知，并根据需要更新自己的状态。

通过使用观察者模式，可以让 Subject 与 Observer 松耦合，能够实现在不同的场景中复用 Subject 和 Observer，并且添加或删除 Observer 更加容易，同时也减少了系统的耦合度。

举个例子，假设我们正在开发一个天气预报应用程序。在该程序中，我们需要将天气状态通知给不同的用户，这些用户可能是通过短信、邮件或者App来获取信息的。

在这种情况下，我们可以将天气数据模型作为 Subject，而不同的用户（比如短信用户、邮件用户、App用户）作为 Observer。当天气数据发生改变时，Subject会通知所有 Observer，并更新其状态。

举个简单的代码例子：

```java
// Subject
public interface WeatherData {
    public void attach(Observer o);
    public void detach(Observer o);
    public void notifyObservers();
}


// ConcreteSubject
public class WeatherDataImpl implements WeatherData {
    private List<Observer> observers;
    private float temperature;
    private float humidity;
    private float pressure;

    public WeatherDataImpl() {
        observers = new ArrayList<>();
    }

    @Override
    public void attach(Observer o) {
        observers.add(o);
    }

    @Override
    public void detach(Observer o) {
        observers.remove(o);
    }

    @Override
    public void notifyObservers() {
        for (Observer o : observers) {
            o.update(temperature, humidity, pressure);
        }
    }

    public void setMeasurements(float temperature, float humidity, float pressure) {
        this.temperature = temperature;
        this.humidity = humidity;
        this.pressure = pressure;
        measurementsChanged();
    }

    private void measurementsChanged() {
        notifyObservers();
    }
}


// Observer
public interface Observer {
    public void update(float temp, float humidity, float pressure);
}


// ConcreteObserver
public class AppUser implements Observer {
    private float temperature;
    private float humidity;
    private float pressure;

    @Override
    public void update(float temp, float humidity, float pressure) {
        this.temperature = temp;
        this.humidity = humidity;
        this.pressure = pressure;
        display();
    }

    private void display() {
        System.out.println("App User: Temperature - " + temperature +
                ", Humidity - " + humidity + ", Pressure - " + pressure);
    }
}

public class SMSUser implements Observer {
    private float temperature;
    private float humidity;
    private float pressure;

    @Override
    public void update(float temp, float humidity, float pressure) {
        this.temperature = temp;
        this.humidity = humidity;
        this.pressure = pressure;
        display();
    }

    private void display() {
        System.out.println("SMS User: Temperature - " + temperature +
                ", Humidity - " + humidity + ", Pressure - " + pressure);
    }
}

public class WeatherStation {
    public static void main(String args[]) {
        WeatherData weatherData = new WeatherDataImpl();
        Observer appUser = new AppUser();
        Observer smsUser = new SMSUser();

        weatherData.attach(appUser);
        weatherData.attach(smsUser);

        weatherData.setMeasurements(28, 70, 1020);

        weatherData.detach(smsUser);

        weatherData.setMeasurements(25, 60, 1005);

    }
}
```

在上面的代码中，WeatherData是 Subject 接口，WeatherDataImpl是 ConcreteSubject 具体实现，AppUser 和 SMSUser 是具体的 Observer 实现。可以看到，当天气数据更新时，AppUser 和 SMSUser 都能够接收到通知并更新最新的天气数据，同时还可以随时添加或移除不同的观察者对象。