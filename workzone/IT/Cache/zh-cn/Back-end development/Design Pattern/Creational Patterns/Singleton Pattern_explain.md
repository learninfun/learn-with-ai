

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