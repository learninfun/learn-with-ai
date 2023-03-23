+++
title = "觀察者模式"
weight = "7"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Observer Pattern的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Observer Pattern？ 
2. 如何實現Observer Pattern？ 
3. 觀察者和被觀察者之間的關係是什麼？ 
4. 如何處理當被觀察者狀態改變時通知觀察者？ 
5. Observer Pattern 在什麼場景下會使用到？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Observer Pattern並舉例
{{< /ask_chatgpt >}}



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

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Observer Pattern的重點
{{< /ask_chatgpt >}}



Observer Pattern 的重點包括：

1. 定義主題（Subject）和觀察者（Observer）的關係，主題維護所有已註冊的觀察者，當主題狀態發生變化時，通知所有觀察者。

2. 主題和觀察者可分別定義一個介面，讓實現類別繼承並實現介面方法以定義主題和觀察者的行為。

3. 觀察者可註冊於多個主題，當某一主題狀態發生變化時，觀察者即可同時收到相應的通知。

4. Observer Pattern 可以實現鬆耦合關係，主題和觀察者只需維護各自的狀態，彼此無需知道對方的實現細節。

5. 主題和觀察者之間的通信可以使用多種方式實現，如事件通知、回呼方法等。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Observer Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實作一個訂閱系統，目標物件可以讓任意數量的Observer訂閱，當目標物件狀態改變時，通知所有Observer。 

2. 實作一個新聞發佈訂閱系統，資料提供者可以發佈新聞訊息給所有訂閱者，訂閱者也可以取消訂閱任意新聞類別。 

3. 實作一個股票價格監控系統，股票物件可以讓許多投資者訂閱，當股票價格變動時，通知訂閱者。 

4. 實作一個天氣預報系統，資料提供者可以在每天清晨發送當日天氣預報訊息給所有訂閱者，外出活動者可以根據天氣預報做出安排。 

5. 實作一個訂票系統，訂票顧客可以訂閱想要買的演唱會或活動，當有票開放購買時，通知所有訂閱者。 

答案：

1.

```java
import java.util.ArrayList;
import java.util.List;

public abstract class Subject {
    private List<Observer> observers = new ArrayList<>();

    public void attach(Observer observer) {
        observers.add(observer);
    }

    public void detach(Observer observer) {
        observers.remove(observer);
    }

    public void notifyObservers() {
        for (Observer observer : observers) {
            observer.update();
        }
    }
}

public class ConcreteSubject extends Subject {
    private int state;

    public int getState() {
        return state;
    }

    public void setState(int state) {
        this.state = state;
        notifyObservers();
    }
}

public abstract class Observer {
    public abstract void update();
}

public class ConcreteObserver extends Observer {
    private ConcreteSubject subject;

    public ConcreteObserver(ConcreteSubject subject) {
        this.subject = subject;
        subject.attach(this);
    }

    @Override
    public void update() {
        System.out.println("ConcreteObserver has updated its state to " + subject.getState());
    }
}

public class Main {
    public static void main(String[] args) {
        ConcreteSubject subject = new ConcreteSubject();
        ConcreteObserver observer1 = new ConcreteObserver(subject);
        ConcreteObserver observer2 = new ConcreteObserver(subject);

        subject.setState(5);  // expect observer1 and observer2 to print "ConcreteObserver has updated its state to 5"
    }
}
```

2.

```java
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public interface Observer {
    void update(String message, String category);
}

public interface Subject {
    void registerObserver(Observer observer, String category);
    void unregisterObserver(Observer observer, String category);
    void notifyObservers(String message, String category);
}

public class NewsPublisher implements Subject {
    private Map<String, List<Observer>> observersByCategory = new HashMap<>();

    @Override
    public void registerObserver(Observer observer, String category) {
        if (!observersByCategory.containsKey(category)) {
            observersByCategory.put(category, new ArrayList<>());
        }

        observersByCategory.get(category).add(observer);
    }

    @Override
    public void unregisterObserver(Observer observer, String category) {
        if (!observersByCategory.containsKey(category)) {
            return;
        }

        List<Observer> observers = observersByCategory.get(category);
        observers.remove(observer);
        if (observers.isEmpty()) {
            observersByCategory.remove(category);
        }
    }

    @Override
    public void notifyObservers(String message, String category) {
        if (!observersByCategory.containsKey(category)) {
            return;
        }

        List<Observer> observers = observersByCategory.get(category);
        for (Observer observer : observers) {
            observer.update(message, category);
        }
    }

    public void publishNews(String message, String category) {
        System.out.println("Publishing news: " + message);
        notifyObservers(message, category);
    }
}

public class NewsSubscriber implements Observer {
    private String name;

    public NewsSubscriber(String name) {
        this.name = name;
    }

    @Override
    public void update(String message, String category) {
        System.out.println(name + " received news update for category " + category + ": " + message);
    }
}

public class Main {
    public static void main(String[] args) {
        NewsPublisher publisher = new NewsPublisher();

        NewsSubscriber subscriber1 = new NewsSubscriber("Subscriber 1");
        publisher.registerObserver(subscriber1, "Politics");
        NewsSubscriber subscriber2 = new NewsSubscriber("Subscriber 2");
        publisher.registerObserver(subscriber2, "Sports");

        publisher.publishNews("Breaking news: Election results are in.", "Politics");  // expect subscriber1 to receive message
        publisher.publishNews("Sports update: Lakers win 112-105.", "Sports");  // expect subscriber2 to receive message

        publisher.unregisterObserver(subscriber1, "Politics");
        publisher.unregisterObserver(subscriber2, "Sports");
    }
}
```

3.

```java
import java.util.ArrayList;
import java.util.List;

public interface Subject {
    void registerObserver(Observer observer);
    void unregisterObserver(Observer observer);
    void notifyObservers();
}

public interface Observer {
    void update(Stock stock);
}

public class Stock {
    private String symbol;
    private double price;

    public Stock(String symbol, double price) {
        this.symbol = symbol;
        this.price = price;
    }

    public String getSymbol() {
        return symbol;
    }

    public void setSymbol(String symbol) {
        this.symbol = symbol;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }
}

public class StockMarket implements Subject {
    private List<Observer> observers = new ArrayList<>();
    private List<Stock> stocks = new ArrayList<>();

    @Override
    public void registerObserver(Observer observer) {
        observers.add(observer);
    }

    @Override
    public void unregisterObserver(Observer observer) {
        observers.remove(observer);
    }

    @Override
    public void notifyObservers() {
        for (Stock stock : stocks) {
            for (Observer observer : observers) {
                observer.update(stock);
            }
        }
    }

    public void addStock(Stock stock) {
        stocks.add(stock);
        notifyObservers();
    }

    public void updateStockPrice(String symbol, double price) {
        for (Stock stock : stocks) {
            if (stock.getSymbol().equals(symbol)) {
                stock.setPrice(price);
                notifyObservers();
                break;
            }
        }
    }
}

public class Investor implements Observer {
    private String name;

    public Investor(String name) {
        this.name = name;
    }

    @Override
    public void update(Stock stock) {
        System.out.println(name + " received update for stock " + stock.getSymbol() + ": " + "price is now " + stock.getPrice());
    }
}

public class Main {
    public static void main(String[] args) {
        StockMarket stockMarket = new StockMarket();

        Stock appleStock = new Stock("AAPL", 135.50);
        stockMarket.addStock(appleStock);

        Investor investor1 = new Investor("Investor 1");
        stockMarket.registerObserver(investor1);

        stockMarket.updateStockPrice("AAPL", 136.00);  // expect investor1 to receive update

        stockMarket.unregisterObserver(investor1);
    }
}
```

4.

```java
import java.util.ArrayList;
import java.util.List;

public interface Subject {
    void registerObserver(Observer observer);
    void unregisterObserver(Observer observer);
    void notifyObservers();
}

public interface Observer {
    void update(String forecast);
}

public class WeatherForecast implements Subject {
    private List<Observer> observers = new ArrayList<>();
    private String forecast;

    @Override
    public void registerObserver(Observer observer) {
        observers.add(observer);
    }

    @Override
    public void unregisterObserver(Observer observer) {
        observers.remove(observer);
    }

    @Override
    public void notifyObservers() {
        for (Observer observer : observers) {
            observer.update(forecast);
        }
    }

    public void setForecast(String forecast) {
        this.forecast = forecast;
        notifyObservers();
    }
}

public class OutdoorActivity implements Observer {
    private String activityName;

    public OutdoorActivity(String activityName) {
        this.activityName = activityName;
    }

    @Override
    public void update(String forecast) {
        System.out.println(activityName + " received weather forecast: " + forecast);
        if (forecast.equals("Sunny")) {
            System.out.println(activityName + " is going on " + activityName + "!");
        } else {
            System.out.println(activityName + " is staying home.");
        }
    }
}

public class Main {
    public static void main(String[] args) {
        WeatherForecast weatherForecast = new WeatherForecast();

        OutdoorActivity hiking = new OutdoorActivity("Hiking");
        weatherForecast.registerObserver(hiking);
        OutdoorActivity beach = new OutdoorActivity("Beach");
        weatherForecast.registerObserver(beach);

        weatherForecast.setForecast("Sunny");  // expect hiking and beach to receive forecast and go on activity
        weatherForecast.setForecast("Rainy");  // expect hiking and beach to stay home

        weatherForecast.unregisterObserver(hiking);
        weatherForecast.unregisterObserver(beach);
    }
}
```

5.

```java
import java.util.ArrayList;
import java.util.List;

public interface Subject {
    void registerObserver(Observer observer);
    void unregisterObserver(Observer observer);
    void notifyObservers();
}

public interface Observer {
    void update(String event);
}

public class TicketSystem implements Subject {
    private List<Observer> observers = new ArrayList<>();
    private List<String> availableEvents = new ArrayList<>();

    @Override
    public void registerObserver(Observer observer) {
        observers.add(observer);
    }

    @Override
    public void unregisterObserver(Observer observer) {
        observers.remove(observer);
    }

    @Override
    public void notifyObservers() {
        for (Observer observer : observers) {
            observer.update("New tickets available!");
        }
    }

    public void addEvent(String event) {
        availableEvents.add(event);
    }

    public void openTicketSales(String event) {
        availableEvents.remove(event);
        notifyObservers();
    }
}

public class Customer implements Observer {
    private String name;

    public Customer(String name) {
        this.name = name;
    }

    @Override
    public void update(String event) {
        System.out.println(name + " received ticket update: " + event);
    }
}

public class Main {
    public static void main(String[] args) {
        TicketSystem ticketSystem = new TicketSystem();
        ticketSystem.addEvent("Concert");
        ticketSystem.addEvent("Musical");

        Customer customer1 = new Customer("Customer 1");
        ticketSystem.registerObserver(customer1);
        Customer customer2 = new Customer("Customer 2");
        ticketSystem.registerObserver(customer2);

        ticketSystem.openTicketSales("Musical");  // expect customer1 and customer2 to receive update

        ticketSystem.unregisterObserver(customer1);
        ticketSystem.unregisterObserver(customer2);
    }
}
```   

