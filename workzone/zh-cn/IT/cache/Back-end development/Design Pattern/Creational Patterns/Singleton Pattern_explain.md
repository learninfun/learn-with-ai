

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