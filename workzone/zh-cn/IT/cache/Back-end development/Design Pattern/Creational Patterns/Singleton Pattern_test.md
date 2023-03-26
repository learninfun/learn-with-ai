

1. 请使用 Singleton Pattern 实现一个缓存 Cache，要求：

- Cache 能够储存多个不同类型的物件，且能够设定每个缓存物件的超时时间；
- Cache 能够根据某些条件撤销缓存的物件；
- Cache 能够在达到缓存上限时自动进行擦除。

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

2. 请使用 Singleton Pattern 实现一个全域配置管理器 ConfigurationManager，要求：

- ConfigurationManager 能够读取/写入配置文件，并支持不同格式配置文件的解析；
- ConfigurationManager 能够记录最近一次读取/写入配置文件的时间；
- ConfigurationManager 能够支持多执行绪并发存取，不会发生竞争状态。

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

3. 请使用 Singleton Pattern 实现一个图表绘制工具 ChartTool，要求：

- ChartTool 能够绘制不同类型的图表，如折线图、柱状图等；
- ChartTool 能够设定不同类型图表的显示样式，如颜色、字型等；
- ChartTool 能够支持曲线、标签、标题等多种元素的添加及修改。

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

4. 请使用 Singleton Pattern 实现一个日志系统 LogManager，要求：

- LogManager 能够支持不同类型的日志输出，如控制台、文件、网路等；
- LogManager 能够设定日志的等级，如 DEBUG、INFO、WARN、ERROR 等；
- LogManager 能够将日志输出到多个目标。

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

5. 请使用 Singleton Pattern 实现一个购物车 ShoppingCart，要求：

- ShoppingCart 能够添加、删除、修改购物车中的商品；
- ShoppingCart 能够显示购物车中所有商品的清单和总价格；
- ShoppingCart 能够将购物车中的商品保存到资料库中。

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