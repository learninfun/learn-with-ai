+++
title = "单例模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Singleton Pattern的问题
{{< /ask_chatgpt >}}



1. 什么是Singleton Pattern？
2. 为什么使用Singleton Pattern？
3. Singleton Pattern如何实现？
4. 单例模式有哪些常见的使用情境？
5. Singleton Pattern有哪些优点和缺点？   

## 说明知识
{{< ask_chatgpt >}}
说明Singleton Pattern并举例
{{< /ask_chatgpt >}}



Singleton Pattern是一种软体设计模式，其目的在于确保一个类别只有一个实例，并提供一个全域唯一的访问点。

举例来说，一个应用程式可能需要一个设置管理器（例如全局配置设置），为了确保不会意外地创建多个实例，可以使用Singleton Pattern来确保该设置管理器只有一个实例存在。

以下为一个简单的Singleton Pattern的范例：

```python
class MySingleton:
  __instance = None
  def __init__(self):
      if MySingleton.__instance != None:
          raise Exception("Singleton class, use get_instance() instead")
      else:
          MySingleton.__instance = self

  @staticmethod
  def get_instance():
      if MySingleton.__instance == None:
          MySingleton()
      return MySingleton.__instance
```

在上面的范例中，我们创建了一个Singleton类别`MySingleton`，它具有一个私有的实例变量`__instance`。

在`__init__`函数中，我们检查实例变量是否为None。如果它不是None，则意味着一个实例已经存在，并引发异常。否则，将`__instance`设置为当前的实例。

在`get_instance()`方法中，我们检查实例变量是否为None。如果为None，则创建一个新的`MySingleton`实例并设置`__instance`。如果不为None，我们直接回传`__instance`即可。

这样一来，无论如何都只会存在一个`MySingleton`实例。

可以透过以下方式获取`MySingleton`实例：

```python
my_singleton = MySingleton.get_instance()
```

如果尝试创建更多的实例，即使直接创建实例也会引发异常，因为`__instance`变数已被设置为相应的实例。   

## 汇总重点
{{< ask_chatgpt >}}
条列Singleton Pattern的重点
{{< /ask_chatgpt >}}



1. Singleton模式是一种创建型设计模式，它确保一个类只有一个实例，并提供一个全局访问点。

2. Singleton模式需要私有化构造函数，以防止该类被实例化多次。

3. Singleton模式需要提供一个静态方法来获取该类的唯一实例，通常是getInstance()方法。

4. Singleton的实例通常是延迟加载的，在需要的时候才创建。

5. Singleton模式提供了一种更好的控制全局状态的方式，但也具有一些缺点，例如难以测试、可能会导致不必要的耦合等。

6. Singleton模式有多种实现方式，包括静态字段、枚举、内部类等。选择哪种方式取决于具体需求以及编程语言的特性。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Singleton Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 請使用 Singleton Pattern 實現一個緩存 Cache，要求：

- Cache 能夠儲存多個不同類型的物件，且能夠設定每個緩存物件的超時時間；
- Cache 能夠根據某些條件撤銷緩存的物件；
- Cache 能夠在達到緩存上限時自動進行擦除。

答案：

```java
public class Cache {
  private static final long DEFAULT_TIMEOUT = 3600;
  private static final int MAX_CAPACITY = 1000;

  private static Cache instance;

  private Map<String, Object> cacheMap = new HashMap<>();
  private Map<String, Long> timeoutMap = new HashMap<>();
  private LinkedList<String> keys = new LinkedList<>();

  private Cache() {}

  public static Cache getInstance() {
    if (instance == null) {
      synchronized (Cache.class) {
        if (instance == null) {
          instance = new Cache();
        }
      }
    }
    return instance;
  }

  public void put(String key, Object value, long timeout) {
    synchronized (cacheMap) {
      if (keys.size() >= MAX_CAPACITY) {
        String firstKey = keys.getFirst();
        remove(firstKey);
      }
      cacheMap.put(key, value);
      timeoutMap.put(key, System.currentTimeMillis() + timeout);
      keys.add(key);
    }
  }

  public void put(String key, Object value) {
    put(key, value, DEFAULT_TIMEOUT);
  }

  public Object get(String key) {
    synchronized (cacheMap) {
      if (!cacheMap.containsKey(key)) {
        return null;
      }
      if (timeoutMap.get(key) < System.currentTimeMillis()) {
        remove(key);
        return null;
      }
      return cacheMap.get(key);
    }
  }

  public void remove(String key) {
    synchronized (cacheMap) {
      cacheMap.remove(key);
      timeoutMap.remove(key);
      keys.remove(key);
    }
  }

  public void cleanup() {
    synchronized (cacheMap) {
      Iterator<String> it = keys.iterator();
      while (it.hasNext()) {
        String key = it.next();
        if (timeoutMap.get(key) < System.currentTimeMillis()) {
          it.remove();
          cacheMap.remove(key);
          timeoutMap.remove(key);
        }
      }
    }
  }
}
```

2. 請使用 Singleton Pattern 實現一個全域配置管理器 ConfigurationManager，要求：

- ConfigurationManager 能夠讀取/寫入配置文件，並支持不同格式配置文件的解析；
- ConfigurationManager 能夠記錄最近一次讀取/寫入配置文件的時間；
- ConfigurationManager 能夠支持多執行緒並發存取，不會發生競爭狀態。

答案：

```java
public class ConfigurationManager {
  private static ConfigurationManager instance;

  private Map<String, String> config = new HashMap<>();
  private DateFormat dateFormat = new SimpleDateFormat("yyyy/MM/dd HH:mm:ss");

  private ConfigurationManager() {}

  public static ConfigurationManager getInstance() {
    if (instance == null) {
      synchronized (ConfigurationManager.class) {
        if (instance == null) {
          instance = new ConfigurationManager();
        }
      }
    }
    return instance;
  }

  public synchronized boolean load(String filename) {
    try {
      // read config file
      Properties properties = new Properties();
      try (InputStream input = new FileInputStream(filename)) {
        properties.load(input);
      }

      // parse config file
      for (String name : properties.stringPropertyNames()) {
        String value = properties.getProperty(name);
        config.put(name, value);
      }

      // record last modified time
      File configFile = new File(filename);
      config.put("last_modified_time", dateFormat.format(configFile.lastModified()));
      return true;
    } catch (IOException e) {
      System.err.println("Failed to load configuration file: " + filename);
      return false;
    }
  }

  public synchronized boolean save(String filename) {
    try {
      // write config file
      Properties properties = new Properties();
      for (Map.Entry<String, String> entry : config.entrySet()) {
        properties.setProperty(entry.getKey(), entry.getValue());
      }
      try (OutputStream output = new FileOutputStream(filename)) {
        properties.store(output, /* comments = */ null);
      }

      // record last modified time
      File configFile = new File(filename);
      config.put("last_modified_time", dateFormat.format(configFile.lastModified()));
      return true;
    } catch (IOException e) {
      System.err.println("Failed to save configuration file: " + filename);
      return false;
    }
  }

  public synchronized String get(String name) {
    return config.get(name);
  }

  public synchronized void set(String name, String value) {
    config.put(name, value);
  }

  public synchronized String getLastModifiedTime() {
    return config.get("last_modified_time");
  }
}
```

3. 請使用 Singleton Pattern 實現一個圖表繪製工具 ChartTool，要求：

- ChartTool 能夠繪製不同類型的圖表，如折線圖、柱狀圖等；
- ChartTool 能夠設定不同類型圖表的顯示樣式，如顏色、字型等；
- ChartTool 能夠支持曲線、標籤、標題等多種元素的添加及修改。

答案：

```java
public class ChartTool {
  private static ChartTool instance;

  private ChartTool() {}

  public static ChartTool getInstance() {
    if (instance == null) {
      synchronized (ChartTool.class) {
        if (instance == null) {
          instance = new ChartTool();
        }
      }
    }
    return instance;
  }

  public void drawLineChart(List<XYPoint> points, Color color, String title) {
    // draw line chart
  }

  public void drawBarChart(List<XYPoint> points, Color color, String title) {
    // draw bar chart
  }

  // add other drawing methods as needed

  public void setFont(Font font) {
    // set font
  }

  public void setColor(Color color) {
    // set color
  }

  public void setTitle(String title) {
    // set title
  }

  public void addCurve(List<XYPoint> points, Color color) {
    // add curve
  }

  public void addLabel(String text, XYPoint position) {
    // add label
  }
}
```

4. 請使用 Singleton Pattern 實現一個日誌系統 LogManager，要求：

- LogManager 能夠支持不同類型的日誌輸出，如控制臺、文件、網路等；
- LogManager 能夠設定日誌的等級，如 DEBUG、INFO、WARN、ERROR 等；
- LogManager 能夠將日誌輸出到多個目標。

答案：

```java
public class LogManager {
  private static LogManager instance;

  private List<LogOutput> logOutputs = new ArrayList<>();
  private LogLevel logLevel = LogLevel.INFO;

  private LogManager() {}

  public static LogManager getInstance() {
    if (instance == null) {
      synchronized (LogManager.class) {
        if (instance == null) {
          instance = new LogManager();
        }
      }
    }
    return instance;
  }

  public void addLogOutput(LogOutput logOutput) {
    logOutputs.add(logOutput);
  }

  public void setLogLevel(LogLevel logLevel) {
    this.logLevel = logLevel;
  }

  public void debug(String message) {
    log(LogLevel.DEBUG, message);
  }

  public void info(String message) {
    log(LogLevel.INFO, message);
  }

  public void warn(String message) {
    log(LogLevel.WARN, message);
  }

  public void error(String message) {
    log(LogLevel.ERROR, message);
  }

  private void log(LogLevel level, String message) {
    if (level.compareTo(logLevel) >= 0) {
      String formattedMessage = String.format("%s [%s] %s",
          new Date(), level.toString(), message);
      for (LogOutput logOutput : logOutputs) {
        logOutput.write(formattedMessage);
      }
    }
  }
}

public interface LogOutput {
  void write(String message);
}

public enum LogLevel {
  DEBUG, INFO, WARN, ERROR;
}
```

5. 請使用 Singleton Pattern 實現一個購物車 ShoppingCart，要求：

- ShoppingCart 能夠添加、刪除、修改購物車中的商品；
- ShoppingCart 能夠顯示購物車中所有商品的清單和總價格；
- ShoppingCart 能夠將購物車中的商品保存到資料庫中。

答案：

```java
public class ShoppingCart {
  private static ShoppingCart instance;

  private List<Item> items = new ArrayList<>();

  private ShoppingCart() {}

  public static ShoppingCart getInstance() {
    if (instance == null) {
      synchronized (ShoppingCart.class) {
        if (instance == null) {
          instance = new ShoppingCart();
        }
      }
    }
    return instance;
  }

  public void addItem(Item item) {
    items.add(item);
  }

  public void removeItem(Item item) {
    items.remove(item);
  }

  public void updateItemQuantity(Item item, int quantity) {
    item.setQuantity(quantity);
  }

  public List<Item> getItems() {
    return Collections.unmodifiableList(items);
  }

  public double getTotalPrice() {
    return items.stream()
        .mapToDouble(item -> item.getPrice() * item.getQuantity())
        .sum();
  }

  public void saveToDatabase() {
    // save items to database
  }
}

public class Item {
  private String id;
  private String name;
  private double price;
  private int quantity;

  public Item(String id, String name, double price, int quantity) {
    this.id = id;
    this.name = name;
    this.price = price;
    this.quantity = quantity;
  }

  // getters and setters
}
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Singleton Pattern的网络数据
{{< /ask_chatgpt >}}



1. Singleton Design Pattern in Java 
https://www.geeksforgeeks.org/singleton-design-pattern/
這是一篇介紹Singleton Pattern的基本概念和應用方式的文章，並以Java為例說明如何實現Singleton Pattern。包含比較常見的方法和變種方式。

2. Singleton Pattern in C# 
https://www.c-sharpcorner.com/article/singleton-design-pattern-in-c-sharp/
這篇文章同樣介紹了Singleton Pattern的基本概念和實現方式，但更專注於在C#中的應用。其中還包含一些範例代碼，可以幫助讀者清晰理解Singleton Pattern。

3. Singleton Pattern in Python 
https://realpython.com/primer-on-python-singleton-pattern/
這篇文章專注於介紹Python中的Singleton Pattern。除了說明基本概念和實現方式外，還講解了如何使用元程式來實現Singleton Pattern。同時也有使用範例代碼可以參考。

4. Singleton Pattern – The Classic Design Pattern 
https://www.javatpoint.com/singleton-design-pattern-in-java
這是一篇講解Singleton Pattern最基本的實現方式，即使用一個靜態變數和一個私有構造函數來實現Singleton。這種方式比較容易理解，也最常用。

5. Singleton Pattern in JavaScript 
https://www.dofactory.com/javascript/design-patterns/singleton
這篇文章介紹如何在JavaScript中實現Singleton Pattern。除了介紹Singleton Pattern基本概念外，還介紹了使用IIFE（立即調用函數表達式）的方式來實現Singleton。這種方式在JavaScript中比較常用。   

