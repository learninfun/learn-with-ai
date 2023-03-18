

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