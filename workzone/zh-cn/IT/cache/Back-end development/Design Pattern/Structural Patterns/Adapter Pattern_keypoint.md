

1.目标介面(Target Interface)
- 定义客户端希望使用的介面。
- 通常是一个抽像类别或介面。

2.适配器(Adapter)
- 实现目标介面。
- 将不相容的类别(即要转换的类别)转换成目标介面。
- 可以是类别型态或物件型态。

3.不相容的类别(Adaptee)
- 不能直接使用的类别。
- 通常是一个现有的类别。
- 介面和目标介面不相同。

4.客户端(Client)
- 想要使用目标介面的类别。

5.转换过程
- 客户端呼叫适配器的方法。
- 适配器接收到客户端的呼叫，转换成对不相容的类别的方法呼叫。
- 不相容的类别回传结果给适配器。
- 适配器将不相容的类别的结果转换成目标介面的结果，回传给客户端。

6.类别适配器和物件适配器
- 类别适配器: 继承适配器和不相容的类别。
- 物件适配器: 组合适配器和不相容的类别。