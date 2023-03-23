+++
title = "创建型模式"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Creational Patterns的问题
{{< /ask_chatgpt >}}



1. 甚麼是Creational Patterns？它們主要用於什麼目的？
2. 哪些是Creational Patterns的主要類別？請述其基本原理及使用情境。
3. 為什麼我們需要使用Creational Patterns？可以舉出一個實際的案例來說明嗎？
4. 如何在Creational Patterns中選擇適合的模式？有哪些因素需要考慮？
5. 何為Singleton模式？它和其他Creational Patterns的區別是甚麼？該如何適當地應用Singleton模式？   

## 说明知识
{{< ask_chatgpt >}}
说明Creational Patterns并举例
{{< /ask_chatgpt >}}



Creational Patterns是指軟件設計中一種與對像創建有關的模式，這些模式可以幫助我們創建不同類型的對象，並且更靈活地應對不同的需求。Creational Patterns主要涉及到以下三種模式： 

1. Singleton Pattern（單例模式）：用來保證一個類只會有一個實例，並且提供一個全局唯一的訪問點來訪問該實例。 

例如，一個應用程序可能需要一個全局設置或是資料庫連接，Singleton Pattern可以用來保證這些對像只會有一個實例，這樣可以減少管理和資源浪費。 

2. Factory Method Pattern（工廠方法模式）：定義一個用於創建對象的介面，讓子類來決定實例化哪個類。 

例如，當創建一個對像時，可能需要遵循一定的流程（例如設置參數、初始化等），使用工廠方法可以將這些流程封裝起來，並且交由子類實現，從而實現更加靈活的對象創建。 

3. Builder Pattern（建造者模式）：用來組合一個複雜的對象，同時隱藏其創建過程。 

例如，當創建一個複雜的對象時，可能涉及到許多子部分，並且需要按照一定的步驟來組裝，使用Builder Pattern可以將這些步驟封裝起來，讓用戶只需要指定需要的部分即可構建所需對象。 

以上就是三種Creational Patterns的基本介紹及其使用場景。   

## 汇总重点
{{< ask_chatgpt >}}
条列Creational Patterns的重点
{{< /ask_chatgpt >}}



Creational Patterns 是建立物件的软体设计模式，其重点包括：

1. 抽像化建立物件过程：Creational Patterns 通常使用工厂方法、抽像工厂、建造者、原型和单例等机制，以抽像化建立物件过程，让程式设计可以更灵活地处理物件的建立。

2. 简化物件建立：Creational Patterns 可以降低程式设计师建立物件的复杂度，让开发者专注于业务逻辑的实现。

3. 易于维护和扩展：Creational Patterns 可以帮助程式设计师维护和扩展系统，使程式更易于维护和修改。

4. 提高程式的可测性：Creational Patterns 可以改善程式的可测性，帮助开发者进行单元测试，以确保程式逻辑的正确性。

5. 适用范围广：Creational Patterns 适用于各种应用场景，包括桌面应用程式、Web 应用程式、移动应用程式等等，对于复杂的系统而言尤其有用。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Creational Patterns的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



问题1：在使用工厂模式时，当需要创建复杂对像时，应该使用哪种变体？

问题2：以下哪个是一个建造者模式的关键元素？

A.抽像生成器

B.具体生成器

C.产品

D.主管

问题3：以下哪个创建了单例模式的图形？

A.三角形

B.正方形

C.菱形

D.圆形

问题4：在使用抽像工厂模式时，如何确定哪个具体工厂应该被使用？

问题5：下面哪个Creational Pattern可以确保在一个应用程序中只有一个实例被创建？

A.工厂模式

B.建造者模式

C.原型模式

D.单例模式

答案：

1.抽像工厂模式

2.D.主管

3.D.圆形

4.根据应用程序的需要

5.D.单例模式   

