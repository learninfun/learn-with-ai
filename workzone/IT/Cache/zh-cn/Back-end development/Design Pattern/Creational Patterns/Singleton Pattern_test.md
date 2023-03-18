

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