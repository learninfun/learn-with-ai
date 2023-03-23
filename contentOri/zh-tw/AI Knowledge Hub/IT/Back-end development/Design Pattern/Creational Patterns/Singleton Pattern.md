+++
title = "單例模式"
weight = "5"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Singleton Pattern的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Singleton Pattern？
2. 為什麼使用Singleton Pattern？
3. Singleton Pattern如何實現？
4. 單例模式有哪些常見的使用情境？
5. Singleton Pattern有哪些優點和缺點？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Singleton Pattern並舉例
{{< /ask_chatgpt >}}



Singleton Pattern是一種軟體設計模式，其目的在於確保一個類別只有一個實例，並提供一個全域唯一的訪問點。

舉例來說，一個應用程式可能需要一個設置管理器（例如全局配置設置），為了確保不會意外地創建多個實例，可以使用Singleton Pattern來確保該設置管理器只有一個實例存在。

以下為一個簡單的Singleton Pattern的範例：

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

在上面的範例中，我們創建了一個Singleton類別`MySingleton`，它具有一個私有的實例變量`__instance`。

在`__init__`函數中，我們檢查實例變量是否為None。如果它不是None，則意味著一個實例已經存在，並引發異常。否則，將`__instance`設置為當前的實例。

在`get_instance()`方法中，我們檢查實例變量是否為None。如果為None，則創建一個新的`MySingleton`實例並設置`__instance`。如果不為None，我們直接回傳`__instance`即可。

這樣一來，無論如何都只會存在一個`MySingleton`實例。

可以透過以下方式獲取`MySingleton`實例：

```python
my_singleton = MySingleton.get_instance()
```

如果嘗試創建更多的實例，即使直接創建實例也會引發異常，因為`__instance`變數已被設置為相應的實例。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Singleton Pattern的重點
{{< /ask_chatgpt >}}



1. Singleton模式是一種創建型設計模式，它確保一個類只有一個實例，並提供一個全局訪問點。

2. Singleton模式需要私有化構造函數，以防止該類被實例化多次。

3. Singleton模式需要提供一個靜態方法來獲取該類的唯一實例，通常是getInstance()方法。

4. Singleton的實例通常是延遲加載的，在需要的時候才創建。

5. Singleton模式提供了一種更好的控制全局狀態的方式，但也具有一些缺點，例如難以測試、可能會導致不必要的耦合等。

6. Singleton模式有多種實現方式，包括靜態字段、枚舉、內部類等。選擇哪種方式取決於具體需求以及編程語言的特性。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Singleton Pattern的中等難度問題，並在後面列出答案
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

