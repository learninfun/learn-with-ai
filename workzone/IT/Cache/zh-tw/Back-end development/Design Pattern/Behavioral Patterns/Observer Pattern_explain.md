

Observer Pattern是一種經典的行為型設計模式，它定義了對象間的一種一對多（one-to-many）的依賴關係，以便當一個對象的狀態發生改變時，其所有依賴對象都能夠自動收到通知並更新狀態。

該模式中包含三種角色：Subject（目標對像）、Observer（觀察者對像）和ConcreteObserver（具體的觀察者對像）。

Subject是被觀察的對象，它維護一組Observer對象，當它的狀態發生變化時，會遍歷這些Observer對象，並調用它們的update()方法通知它們。

Observer是觀察者對象，它定義了一個接口，用於接收關於Subject狀態改變的通知，並根據需要更新自己的狀態。

通過使用觀察者模式，可以讓 Subject 與 Observer 松耦合，能夠實現在不同的場景中復用 Subject 和 Observer，並且添加或刪除 Observer 更加容易，同時也減少了系統的耦合度。

舉個例子，假設我們正在開發一個天氣預報應用程序。在該程序中，我們需要將天氣狀態通知給不同的用戶，這些用戶可能是通過短信、郵件或者App來獲取信息的。

在這種情況下，我們可以將天氣數據模型作為 Subject，而不同的用戶（比如短信用戶、郵件用戶、App用戶）作為 Observer。當天氣數據發生改變時，Subject會通知所有 Observer，並更新其狀態。

舉個簡單的代碼例子：

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

在上面的代碼中，WeatherData是 Subject 接口，WeatherDataImpl是 ConcreteSubject 具體實現，AppUser 和 SMSUser 是具體的 Observer 實現。可以看到，當天氣數據更新時，AppUser 和 SMSUser 都能夠接收到通知並更新最新的天氣數據，同時還可以隨時添加或移除不同的觀察者對象。