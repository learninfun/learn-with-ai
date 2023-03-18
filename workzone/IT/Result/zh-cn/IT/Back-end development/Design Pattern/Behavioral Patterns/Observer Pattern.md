+++
title = "观察者模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Observer Pattern的问题
{{< /ask_chatgpt >}}



1. 什么是Observer Pattern？ 
2. 如何实现Observer Pattern？ 
3. 观察者和被观察者之间的关系是什么？ 
4. 如何处理当被观察者状态改变时通知观察者？ 
5. Observer Pattern 在什么场景下会使用到？   

## 说明知识
{{< ask_chatgpt >}}
说明Observer Pattern并举例
{{< /ask_chatgpt >}}



Observer Pattern是一种经典的行为型设计模式，它定义了对象间的一种一对多（one-to-many）的依赖关系，以便当一个对象的状态发生改变时，其所有依赖对象都能够自动收到通知并更新状态。

该模式中包含三种角色：Subject（目标对象）、Observer（观察者对象）和ConcreteObserver（具体的观察者对象）。

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

## 汇总重点
{{< ask_chatgpt >}}
条列Observer Pattern的重点
{{< /ask_chatgpt >}}



Observer Pattern 的重点包括：

1. 定义主题（Subject）和观察者（Observer）的关系，主题维护所有已注册的观察者，当主题状态发生变化时，通知所有观察者。

2. 主题和观察者可分别定义一个介面，让实现类别继承并实现介面方法以定义主题和观察者的行为。

3. 观察者可注册于多个主题，当某一主题状态发生变化时，观察者即可同时收到相应的通知。

4. Observer Pattern 可以实现松耦合关系，主题和观察者只需维护各自的状态，彼此无需知道对方的实现细节。

5. 主题和观察者之间的通信可以使用多种方式实现，如事件通知、回呼方法等。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Observer Pattern的中等难度问题，并在后面列出答案
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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Observer Pattern的网络数据
{{< /ask_chatgpt >}}



1. "Observer Pattern in Java: Example & Tutorial" - Medium
https://medium.com/@codingee/observer-pattern-in-java-example-tutorial-3c3cbb6cae71

This article offers a detailed explanation of the Observer Pattern in Java, with an example and tutorial. It covers the basics of the pattern, and how it can be implemented to achieve code modularity and flexibility.

2. "Observer Design Pattern" - Tutorials Point
https://www.tutorialspoint.com/design_pattern/observer_pattern.htm

Tutorials Point offers a concise overview of the Observer Pattern, including its definition, intent, participants, and implementation steps. It also includes sample code snippets to help readers understand the pattern better.

3. "Observer Pattern in JavaScript" - Codeburst
https://codeburst.io/observer-pattern-in-javascript-7d85cc8bdf37

Codeburst explains the Observer Pattern in the context of JavaScript, and how it can be used to create reactive user interfaces. It provides an easy-to-follow guide on how to implement event-driven programming using the pattern.

4. "Observer Pattern in C#" - C# Corner
https://www.c-sharpcorner.com/article/observer-pattern-in-c-sharp/

This article focuses on the implementation of the Observer Pattern in C#, with an example console application. It covers the benefits of using the pattern, and how it can be applied to different scenarios.

5. "Understanding the Observer Pattern in Swift" - Swift Tutorial
https://www.swifttutorialforbeginners.com/design-patterns/observer-pattern/

This Swift Tutorial offers a simplified explanation of the Observer Pattern, and how it can be implemented in Swift. It includes examples of how to use the pattern to update views in iOS applications.   

